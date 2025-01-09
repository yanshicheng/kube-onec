package onecprojectquotaservicelogic

import (
	"context"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/internal/code"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/internal/svc"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/pb"
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddOnecProjectQuotaLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddOnecProjectQuotaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddOnecProjectQuotaLogic {
	return &AddOnecProjectQuotaLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// -----------------------项目与集群的对应关系表，记录资源配额和使用情况-----------------------
func (l *AddOnecProjectQuotaLogic) AddOnecProjectQuota(in *pb.AddOnecProjectQuotaReq) (*pb.AddOnecProjectQuotaResp, error) {
	// 先查询集群信息

	cluster, err := l.svcCtx.ClusterModel.FindOneByUuid(l.ctx, in.ClusterUuid)
	if err != nil {
		l.Logger.Errorf("获取集群信息失败: %v", err)
		return nil, code.GetClusterInfoErr
	}

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

	//计算集群资源是否够分配
	if cpuLimit > (float64(cluster.CpuTotal)-float64(cluster.CpuUsed)) || memoryLimit > (cluster.MemoryTotal-cluster.MemoryUsed) || in.PodLimit > (cluster.PodTotal-cluster.PodUsed) {
		l.Logger.Errorf("集群资源不足")
		return nil, code.ClusterResourceNotEnoughErr
	}

	// 判断 项目是否为默认项目，默认项目不允许创建项目资源
	project, err := l.svcCtx.ProjectModel.FindOne(l.ctx, in.ProjectId)
	if err != nil {
		l.Logger.Errorf("获取项目信息失败: %v", err)
		return nil, code.GetProjectErr
	}

	if project.IsDefault == 1 {
		l.Logger.Errorf("默认项目不允许分配集群资源!")
		return nil, code.ProjectResourceNotAllowedErr
	}

	// 查询是否已经存在 项目资源
	if _, err := l.svcCtx.ProjectQuotaModel.FindOneByClusterUuidProjectId(l.ctx, cluster.Uuid, in.ProjectId); err == nil {
		l.Logger.Errorf("项目资源已存在: %v", in.ProjectId)
		return nil, code.ProjectQuotaExistErr
	}
	// 启动事务
	err = l.svcCtx.ProjectQuotaModel.TransCtx(l.ctx, func(ctx context.Context, session sqlx.Session) error {
		// 创建项目资源
		// 在数据库事务中插入一条项目资源记录
		interSql := `INSERT INTO {table} 
        (cluster_uuid, project_id, cpu_quota, cpu_limit, cpu_overcommit,  
        memory_quota, memory_limit, memory_overcommit, pod_limit, 
        configmap_limit, pvc_limit, storage_limit,nodeport_limit , secret_limit, status, 
        created_by, updated_by) 
        VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ? , ?, ?)`
		var params []interface{}
		params = append(params, in.ClusterUuid)
		params = append(params, in.ProjectId)
		params = append(params, in.CpuQuota)
		params = append(params, cpuLimit)
		params = append(params, in.CpuOvercommit)
		params = append(params, in.MemoryQuota)
		params = append(params, memoryLimit)
		params = append(params, in.MemoryOvercommit)
		params = append(params, in.PodLimit)
		if in.ConfigmapLimit != 0 {
			params = append(params, in.ConfigmapLimit)
		} else {
			params = append(params, 9999)
		}
		if in.PvcLimit != 0 {
			params = append(params, in.PvcLimit)
		} else {
			params = append(params, 9999)
		}
		if in.NodeportLimit != 0 {
			params = append(params, in.NodeportLimit)
		} else {
			params = append(params, 9999)
		}
		if in.StorageLimit != 0 {
			params = append(params, in.StorageLimit)
		} else {
			params = append(params, 9999)
		}
		if in.SecretLimit != 0 {
			params = append(params, in.SecretLimit)
		} else {
			params = append(params, 9999)
		}
		params = append(params, "Active")
		params = append(params, in.CreatedBy)
		params = append(params, in.UpdatedBy)
		_, err := l.svcCtx.ProjectQuotaModel.TransOnSql(ctx, session, 0, interSql, params...)
		if err != nil {
			l.Logger.Errorf("创建项目资源失败: %v", err)
			return code.AddProjectQuotaErr
		}
		l.Logger.Infof("创建项目资源成功: %v", in.ProjectId)
		//修改集群资源使用情况
		interSql = `UPDATE {table} SET cpu_used = ?, memory_used = ?, pod_used = ? WHERE id = ?`
		params = []interface{}{cluster.CpuUsed + in.CpuQuota, cluster.MemoryUsed + in.MemoryQuota, cluster.PodUsed + in.PodLimit, cluster.Id}
		_, err = l.svcCtx.ClusterModel.TransOnSql(ctx, session, cluster.Id, interSql, params...)
		if err != nil {
			l.Logger.Errorf("修改集群资源使用情况失败: %v", err)
			return code.UpdateClusterInfoErr
		}
		l.Logger.Infof("修改集群资源使用情况成功: %v", cluster.Id)
		return nil
	})
	if err != nil {
		l.Logger.Errorf("项目集群资源分配失败，事务执行失败: %v", err)
		return nil, code.ProjectQuotaAllocationErr
	}
	l.Logger.Infof("项目集群资源分配成功: %v", in.ProjectId)
	return &pb.AddOnecProjectQuotaResp{}, nil
}
