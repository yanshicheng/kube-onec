package onecprojectapplicationservicelogic

import (
	"context"
	"fmt"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/internal/code"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/internal/model"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/internal/shared"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/internal/svc"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/pb"
	"github.com/yanshicheng/kube-onec/common/handler/errorx"
	"github.com/yanshicheng/kube-onec/common/k8swrapper/core"
	"github.com/yanshicheng/kube-onec/common/k8swrapper/kubeutils"
	"strings"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateOnecProjectApplicationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateOnecProjectApplicationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateOnecProjectApplicationLogic {
	return &UpdateOnecProjectApplicationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateOnecProjectApplicationLogic) UpdateOnecProjectApplication(in *pb.UpdateOnecProjectApplicationReq) (*pb.UpdateOnecProjectApplicationResp, error) {
	// 查询应用信息
	application, err := l.svcCtx.ProjectApplicationModel.FindOne(l.ctx, in.Id)
	if err != nil {
		l.Logger.Errorf("查询应用失败: %v", err)
		return nil, code.GetApplicationErr
	}

	// 获取项目资源配额
	quota, err := l.svcCtx.ProjectQuotaModel.FindOneByClusterUuidProjectId(l.ctx, application.ClusterUuid, application.ProjectId)
	if err != nil {
		l.Logger.Errorf("查询项目资源配额失败: %v", err)
		return nil, code.GetProjectQuotaErr
	}
	// 检查更新的资源是否超出配额
	if err := l.checkQuotaForUpdate(application, in, quota); err != nil {
		return nil, err
	}

	// 获取 Kubernetes 客户端
	client, err := shared.GetK8sClient(l.ctx, l.svcCtx, application.ClusterUuid)
	if err != nil {
		l.Logger.Errorf("获取集群客户端失败: %v", err)
		return nil, code.GetClusterClientErr
	}
	nsClient := client.GetNamespaceClient()

	// 更新命名空间的资源配额
	if err := l.updateNamespaceQuota(nsClient, application.Identifier, in); err != nil {
		l.Logger.Errorf("更新命名空间资源配额失败: %v", err)
		return nil, code.UpdateNamespaceQuotaErr
	}
	// 更新应用数据库信息
	if err := l.updateApplicationData(application, in); err != nil {
		l.Logger.Errorf("更新应用信息失败: %v", err)
		return nil, code.UpdateApplicationErr
	}
	// 更新项目资源配额
	if err := l.updateQuota(quota, application, in); err != nil {
		l.Logger.Errorf("更新项目资源配额失败: %v", err)
		return nil, code.ProjectQuotaUpdateErr
	}
	return &pb.UpdateOnecProjectApplicationResp{}, nil
}

// 检查更新时资源配额是否足够
func (l *UpdateOnecProjectApplicationLogic) checkQuotaForUpdate(application *model.OnecProjectApplication, in *pb.UpdateOnecProjectApplicationReq, quota *model.OnecProjectQuota) error {
	// application 当前的应用 in 输入进来的资源
	// 检查是否超出配额
	// 定义一个存储超出配额信息的切片
	var exceeded []string

	// 检查每个资源是否超出配额，即：资源的剩余量 = 总配额 - (已用配额 - 当前应用已分配的配额)。
	if in.CpuLimit > (quota.CpuLimit - (quota.CpuUsed - application.CpuLimit)) {
		exceeded = append(exceeded, fmt.Sprintf("CPU 配额不足 (申请: %v, 剩余: %v)", in.CpuLimit, quota.CpuLimit-(quota.CpuUsed-application.CpuLimit)))
	}
	if in.MemoryLimit > (quota.MemoryLimit - (quota.MemoryUsed - application.MemoryLimit)) {
		exceeded = append(exceeded, fmt.Sprintf("内存配额不足 (申请: %v, 剩余: %v)", in.MemoryLimit, quota.MemoryLimit-(quota.MemoryUsed-application.MemoryLimit)))
	}
	if in.StorageLimit > (quota.StorageLimit - (quota.StorageUsed - application.StorageLimit)) {
		exceeded = append(exceeded, fmt.Sprintf("存储配额不足 (申请: %v, 剩余: %v)", in.StorageLimit, quota.StorageLimit-(quota.StorageUsed-application.StorageLimit)))
	}
	if in.PvcLimit > (quota.PvcLimit - (quota.PvcUsed - application.PvcLimit)) {
		exceeded = append(exceeded, fmt.Sprintf("PVC 配额不足 (申请: %v, 剩余: %v)", in.PvcLimit, quota.PvcLimit-(quota.PvcUsed-application.PvcLimit)))
	}
	if in.PodLimit > (quota.PodLimit - (quota.PodUsed - application.PodLimit)) {
		exceeded = append(exceeded, fmt.Sprintf("Pod 配额不足 (申请: %v, 剩余: %v)", in.PodLimit, quota.PodLimit-(quota.PodUsed-application.PodLimit)))
	}
	if in.SecretLimit > (quota.SecretLimit - (quota.SecretUsed - application.SecretLimit)) {
		exceeded = append(exceeded, fmt.Sprintf("Secret 配额不足 (申请: %v, 剩余: %v)", in.SecretLimit, quota.SecretLimit-(quota.SecretUsed-application.SecretLimit)))
	}
	if in.NodeportLimit > (quota.NodeportLimit - (quota.NodeportUsed - application.NodeportLimit)) {
		exceeded = append(exceeded, fmt.Sprintf("NodePort 配额不足 (申请: %v, 剩余: %v)", in.NodeportLimit, quota.NodeportLimit-(quota.NodeportUsed-application.NodeportLimit)))
	}
	if in.ConfigmapLimit > (quota.ConfigmapLimit - (quota.ConfigmapUsed - application.ConfigmapLimit)) {
		exceeded = append(exceeded, fmt.Sprintf("ConfigMap 配额不足 (申请: %v, 剩余: %v)", in.ConfigmapLimit, quota.ConfigmapLimit-(quota.ConfigmapUsed-application.ConfigmapLimit)))
	}
	if in.ServiceLimit > (quota.ServiceLimit - (quota.ServiceUsed - application.ServiceLimit)) {
		exceeded = append(exceeded, fmt.Sprintf("Service 配额不足 (申请: %v, 剩余: %v)", in.ServiceLimit, quota.ServiceLimit-(quota.ServiceUsed-application.ServiceLimit)))
	}

	// 如果有超出配额的资源，返回错误信息
	if len(exceeded) > 0 {
		return errorx.New(102199, fmt.Sprintf("资源配额不足: %s", strings.Join(exceeded, "; ")))
	}

	// 如果没有超出配额，返回 nil
	return nil
}

// 更新命名空间资源配额
func (l *UpdateOnecProjectApplicationLogic) updateNamespaceQuota(nsClient core.NamespacesInterface, identifier string, in *pb.UpdateOnecProjectApplicationReq) error {
	limits := kubeutils.ResourceLimits{
		CPU:                    fmt.Sprintf("%v", in.CpuLimit),
		Memory:                 fmt.Sprintf("%vGi", in.MemoryLimit),
		PersistentVolumeClaims: fmt.Sprintf("%v", in.PvcLimit),
		Pods:                   fmt.Sprintf("%v", in.PodLimit),
		Secrets:                fmt.Sprintf("%v", in.SecretLimit),
		Services:               fmt.Sprintf("%v", in.ServiceLimit),
		Storage:                fmt.Sprintf("%vGi", in.StorageLimit),
		ServicesNodePorts:      fmt.Sprintf("%v", in.NodeportLimit),
		ConfigMaps:             fmt.Sprintf("%v", in.ConfigmapLimit),
	}

	resourceQuota := kubeutils.CreateResourceQuota(
		fmt.Sprintf("%s-onec-default", identifier),
		identifier,
		limits,
	)

	if err := nsClient.UpdateResourceQuota(identifier, resourceQuota); err != nil {
		return err
	}
	return nil
}

// 更新数据库中的应用数据
func (l *UpdateOnecProjectApplicationLogic) updateApplicationData(application *model.OnecProjectApplication, in *pb.UpdateOnecProjectApplicationReq) error {
	application.Name = in.Name
	application.Description = in.Description
	application.CpuLimit = in.CpuLimit
	application.MemoryLimit = in.MemoryLimit
	application.StorageLimit = in.StorageLimit
	application.PvcLimit = in.PvcLimit
	application.PodLimit = in.PodLimit
	application.SecretLimit = in.SecretLimit
	application.NodeportLimit = in.NodeportLimit
	application.ServiceLimit = in.ServiceLimit
	application.ConfigmapLimit = in.ConfigmapLimit
	application.UpdatedBy = in.UpdatedBy

	if err := l.svcCtx.ProjectApplicationModel.Update(l.ctx, application); err != nil {
		return err
	}
	return nil
}

// 更新项目资源配额
func (l *UpdateOnecProjectApplicationLogic) updateQuota(quota *model.OnecProjectQuota, application *model.OnecProjectApplication, in *pb.UpdateOnecProjectApplicationReq) error {
	// 更新各项资源的已用配额
	quota.CpuUsed = (quota.CpuUsed - application.CpuLimit) + in.CpuLimit
	quota.MemoryUsed = (quota.MemoryUsed - application.MemoryLimit) + in.MemoryLimit
	quota.StorageUsed = (quota.StorageUsed - application.StorageLimit) + in.StorageLimit
	quota.PvcUsed = (quota.PvcUsed - application.PvcLimit) + in.PvcLimit
	quota.PodUsed = (quota.PodUsed - application.PodLimit) + in.PodLimit
	quota.SecretUsed = (quota.SecretUsed - application.SecretLimit) + in.SecretLimit
	quota.NodeportUsed = (quota.NodeportUsed - application.NodeportLimit) + in.NodeportLimit
	quota.ConfigmapUsed = (quota.ConfigmapUsed - application.ConfigmapLimit) + in.ConfigmapLimit
	quota.ServiceUsed = (quota.ServiceUsed - application.ServiceLimit) + in.ServiceLimit

	if err := l.svcCtx.ProjectQuotaModel.Update(l.ctx, quota); err != nil {
		return err
	}
	return nil
}
