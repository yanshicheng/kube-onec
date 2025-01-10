package onecprojectapplicationservicelogic

import (
	"context"
	"fmt"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/internal/code"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/internal/model"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/internal/shared"

	"github.com/yanshicheng/kube-onec/application/manager/rpc/internal/svc"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelOnecProjectApplicationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelOnecProjectApplicationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelOnecProjectApplicationLogic {
	return &DelOnecProjectApplicationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelOnecProjectApplicationLogic) DelOnecProjectApplication(in *pb.DelOnecProjectApplicationReq) (*pb.DelOnecProjectApplicationResp, error) {
	// 1. 查询应用信息
	application, err := l.svcCtx.ProjectApplicationModel.FindOne(l.ctx, in.Id)
	if err != nil {
		l.Logger.Errorf("查询应用失败: %v", err)
		return nil, code.GetApplicationErr
	}

	// 2. 获取项目资源配额
	quota, err := l.svcCtx.ProjectQuotaModel.FindOneByClusterUuidProjectId(l.ctx, application.ClusterUuid, application.ProjectId)
	if err != nil {
		l.Logger.Errorf("查询项目资源配额失败: %v", err)
		return nil, code.GetProjectQuotaErr
	}

	// 3. 获取 Kubernetes 客户端
	client, err := shared.GetK8sClient(l.ctx, l.svcCtx, application.ClusterUuid)
	if err != nil {
		l.Logger.Errorf("获取集群客户端失败: %v", err)
		return nil, code.GetClusterClientErr
	}
	nsClient := client.GetNamespaceClient()

	// 4. 删除命名空间资源配额
	resourceQuotaName := fmt.Sprintf("%s-onec-default", application.Identifier)
	if err := nsClient.DeleteResourceQuota(application.Identifier, resourceQuotaName); err != nil {
		l.Logger.Errorf("删除命名空间资源配额失败: %v", err)
		// 返回错误，但可以根据需求决定是否继续删除其他资源
		return nil, code.DeleteNamespaceQuotaErr
	}

	// 5. 删除命名空间
	if err := nsClient.DeleteNamespace(application.Identifier); err != nil {
		l.Logger.Errorf("删除命名空间失败: %v", err)
		return nil, code.DeleteNamespaceErr
	}

	// 6. 更新项目资源配额
	if err := l.updateQuotaOnDelete(quota, application); err != nil {
		l.Logger.Errorf("更新项目资源配额失败: %v", err)
		return nil, code.ProjectQuotaUpdateErr
	}
	// 7. 删除数据库中的应用记录
	if err := l.svcCtx.ProjectApplicationModel.Delete(l.ctx, application.Id); err != nil {
		l.Logger.Errorf("删除应用记录失败: %v", err)
		return nil, code.DeleteApplicationErr
	}

	return &pb.DelOnecProjectApplicationResp{}, nil
}

func (l *DelOnecProjectApplicationLogic) updateQuotaOnDelete(quota *model.OnecProjectQuota, application *model.OnecProjectApplication) error {
	// 恢复已用资源配额
	quota.CpuUsed -= application.CpuLimit
	quota.MemoryUsed -= application.MemoryLimit
	quota.StorageUsed -= application.StorageLimit
	quota.PvcUsed -= application.PvcLimit
	quota.PodUsed -= application.PodLimit
	quota.SecretUsed -= application.SecretLimit
	quota.NodeportUsed -= application.NodeportLimit
	quota.ConfigmapUsed -= application.ConfigmapLimit
	quota.ServiceUsed -= application.ServiceLimit
	quota.UpdatedBy = application.UpdatedBy

	// 调用数据库更新
	if err := l.svcCtx.ProjectQuotaModel.Update(l.ctx, quota); err != nil {
		return err
	}
	return nil
}
