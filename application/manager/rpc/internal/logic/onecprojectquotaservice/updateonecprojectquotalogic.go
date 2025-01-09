package onecprojectquotaservicelogic

import (
	"context"
	"errors"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/internal/code"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/internal/model"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/internal/svc"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/pb"
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateOnecProjectQuotaLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateOnecProjectQuotaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateOnecProjectQuotaLogic {
	return &UpdateOnecProjectQuotaLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateOnecProjectQuotaLogic) UpdateOnecProjectQuota(in *pb.UpdateOnecProjectQuotaReq) (*pb.UpdateOnecProjectQuotaResp, error) {
	// 设置参数默认值
	if in.CpuOvercommit == 0 {
		in.CpuOvercommit = 1
	}
	if in.MemoryOvercommit == 0 {
		in.MemoryOvercommit = 1
	}
	// 计算 项目的资源配额使用情况
	cpuLimit := float64(in.CpuQuota) * in.CpuOvercommit
	memoryLimit := in.MemoryQuota * in.MemoryOvercommit
	quota, err := l.svcCtx.ProjectQuotaModel.FindOne(l.ctx, in.Id)
	if err != nil {
		l.Logger.Errorf("查询项目资源失败: %v", err)
		return nil, code.GetProjectQuotaErr
	}

	project, err := l.svcCtx.ProjectModel.FindOne(l.ctx, quota.ProjectId)
	if err != nil {
		l.Logger.Errorf("查询项目失败: %v", err)
		return nil, code.GetProjectErr
	}
	if project.IsDefault == 1 {
		l.Logger.Errorf("默认项目不允许分配集群资源!")
		return nil, code.ProjectResourceNotAllowedErr
	}

	// 获取项目资源 分配的应用统计资源
	quotaTotal, err := l.svcCtx.ProjectQuotaModel.FindAllByProjectQuotaTotal(l.ctx, quota.ProjectId, quota.ClusterUuid)
	if err == nil || errors.Is(err, model.ErrNotFound) {
		// 判断传入进来的数据是否小与已经分配的总资源，如果小于则不允许修改
		if quotaTotal.CpuLimit > cpuLimit || quotaTotal.MemoryLimit > memoryLimit || quotaTotal.PvcLimit > in.PvcLimit || quotaTotal.ConfigmapLimit > in.ConfigmapLimit || quotaTotal.NodeportLimit > in.NodeportLimit || quotaTotal.PodLimit > in.PodLimit || quotaTotal.StorageLimit > in.StorageLimit || quotaTotal.SecretLimit > in.SecretLimit {
			l.Logger.Errorf("传入的资源小于已经分配的总资源: %v", err)
			return nil, code.ProjectQuotaNotEnoughErr
		}
	} else {
		l.Logger.Errorf("获取项目资源失败: %v", err)
		return nil, code.GetProjectQuotaErr
	}

	// 判断是否小于集群的可用资源
	cluster, err := l.svcCtx.ClusterModel.FindOneByUuid(l.ctx, quota.ClusterUuid)
	if err != nil {
		l.Logger.Errorf("查询集群失败: %v", err)
		return nil, code.GetClusterInfoErr
	}
	// 计算资源差异

	if in.CpuQuota > (cluster.CpuTotal-(cluster.CpuUsed-quota.CpuQuota)) || in.MemoryQuota > (cluster.MemoryTotal-(cluster.MemoryUsed-quota.MemoryQuota)) || in.PodLimit > (cluster.PodTotal-(cluster.PodUsed-quota.PodLimit)) {
		l.Logger.Errorf("集群资源不足: %v", err)
		return nil, code.ClusterResourceNotEnoughErr
	}
	cpuDiff := (cluster.CpuUsed - quota.CpuQuota) + in.CpuQuota
	memoryDiff := (cluster.MemoryUsed - quota.MemoryQuota) + in.MemoryQuota
	podDiff := (cluster.PodUsed - quota.PodLimit) + in.PodLimit
	err = l.svcCtx.ProjectQuotaModel.TransCtx(l.ctx, func(ctx context.Context, session sqlx.Session) error {
		updateSql := `UPDATE {table} 
SET 
    cpu_quota = ?, 
    cpu_limit = ?, 
    cpu_overcommit = ?, 
    memory_quota = ?, 
    memory_limit = ?, 
    memory_overcommit = ?, 
    pod_limit = ?, 
    configmap_limit = ?, 
    pvc_limit = ?, 
    storage_limit = ?, 
    nodeport_limit = ?, 
    secret_limit = ?, 
    updated_by = ?
WHERE id = ?`
		var params []interface{}
		params = append(params, in.CpuQuota)         // cpu_quota
		params = append(params, cpuLimit)            // cpu_limit
		params = append(params, in.CpuOvercommit)    // cpu_overcommit
		params = append(params, in.MemoryQuota)      // memory_quota
		params = append(params, memoryLimit)         // memory_limit
		params = append(params, in.MemoryOvercommit) // memory_overcommit
		params = append(params, in.PodLimit)         // pod_limit
		params = append(params, in.ConfigmapLimit)   // configmap_limit
		params = append(params, in.PvcLimit)         // pvc_limit
		params = append(params, in.StorageLimit)     // storage_limit
		params = append(params, in.NodeportLimit)    // nodeport_limit
		params = append(params, in.SecretLimit)      // secret_limit
		params = append(params, in.UpdatedBy)        // updated_by
		params = append(params, in.Id)               // WHERE 条件: id
		_, err := l.svcCtx.ProjectQuotaModel.TransOnSql(ctx, session, quota.Id, updateSql, params...)
		if err != nil {
			l.Logger.Errorf("更新项目资源失败: %v", err)
			return code.UpdateProjectQuotaErr
		}
		// 更新集群资源
		clusterUpdateSql := `UPDATE {table} 
SET 
    cpu_used = ?, 
    memory_used =  ?, 
    pod_used =  ?
WHERE uuid = ?`
		var clusterParams []interface{}
		clusterParams = append(clusterParams, cpuDiff)
		clusterParams = append(clusterParams, memoryDiff)
		clusterParams = append(clusterParams, podDiff)
		clusterParams = append(clusterParams, quota.ClusterUuid)
		_, err = l.svcCtx.ClusterModel.TransOnSql(ctx, session, cluster.Id, clusterUpdateSql, clusterParams...)
		if err != nil {
			l.Logger.Errorf("更新集群资源失败: %v", err)
			return code.UpdateClusterErr
		}
		return nil
	})
	if err != nil {
		l.Logger.Errorf("项目集群资源分配失败，事务执行失败: %v", err)
		return nil, code.ProjectQuotaAllocationErr
	}
	l.Logger.Infof("项目集群资源分配成功: %v", in.ProjectId)
	return &pb.UpdateOnecProjectQuotaResp{}, nil
}
