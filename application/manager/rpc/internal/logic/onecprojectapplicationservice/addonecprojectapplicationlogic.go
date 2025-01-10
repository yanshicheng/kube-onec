package onecprojectapplicationservicelogic

import (
	"context"
	"fmt"
	"github.com/yanshicheng/kube-onec/common/handler/errorx"
	"github.com/yanshicheng/kube-onec/common/k8swrapper/core"
	"strings"
	"time"

	"github.com/yanshicheng/kube-onec/application/manager/rpc/internal/code"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/internal/model"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/internal/shared"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/internal/svc"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/pb"
	"github.com/yanshicheng/kube-onec/common/k8swrapper/kubeutils"
	"github.com/yanshicheng/kube-onec/pkg/utils"
	"github.com/zeromicro/go-zero/core/logx"
)

type AddOnecProjectApplicationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddOnecProjectApplicationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddOnecProjectApplicationLogic {
	return &AddOnecProjectApplicationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddOnecProjectApplicationLogic) AddOnecProjectApplication(in *pb.AddOnecProjectApplicationReq) (*pb.AddOnecProjectApplicationResp, error) {
	// 验证参数
	if err := utils.ValidateNamespaceName(in.Identifier); err != nil {
		l.Logger.Errorf("参数不合法: %v", err)
		return nil, code.IdentifierErr
	}

	// 获取项目和资源配额
	project, err := l.svcCtx.ProjectModel.FindOne(l.ctx, in.ProjectId)
	if err != nil {
		l.Logger.Errorf("查询项目失败: %v", err)
		return nil, code.GetProjectErr
	}

	quota, err := l.svcCtx.ProjectQuotaModel.FindOneByClusterUuidProjectId(l.ctx, in.ClusterUuid, in.ProjectId)
	if err != nil {
		l.Logger.Errorf("查询项目资源失败: %v", err)
		return nil, code.GetProjectQuotaErr
	}

	// 检查资源配额是否足够
	if err := l.isQuotaInsufficient(in, quota); err != nil {
		l.Logger.Errorf("资源不足，无法创建应用")
		return nil, err
	}

	// 获取 Kubernetes 客户端
	client, err := shared.GetK8sClient(l.ctx, l.svcCtx, in.ClusterUuid)
	if err != nil {
		l.Logger.Errorf("获取集群客户端异常: %v", err)
		return nil, code.GetClusterClientErr
	}

	nsClient := client.GetNamespaceClient()
	if nsClient.NamespaceExist(in.Identifier) {
		l.Logger.Errorf("命名空间 %s 已存在", in.Identifier)
		return nil, code.NamespaceExistErr
	}

	// 构建 Labels 和 Annotations
	labels, annotations := l.buildLabelsAndAnnotations(in, project)

	// 创建命名空间
	if err := l.createNamespace(nsClient, in.Identifier, labels, annotations); err != nil {
		return nil, code.CreateNamespaceErr
	}

	// 创建命名空间资源配额
	if err := l.createNamespaceQuota(nsClient, in, labels, annotations); err != nil {
		err := nsClient.DeleteNamespace(in.Identifier)
		if err != nil {
			l.Logger.Errorf("清理命名空间失败: %v", err)
		} // 清理命名空间
		return nil, code.AddProjectQuotaErr
	}

	// 插入应用数据
	if err := l.insertApplicationData(in, labels, annotations); err != nil {
		l.cleanUpNamespace(nsClient, in.Identifier, fmt.Sprintf("%s-onec-default", in.Identifier))
		return nil, code.AddProjectApplicationErr
	}

	// 更新项目资源配额
	if err := l.updateQuota(quota, in); err != nil {
		l.Logger.Errorf("更新项目资源配额失败: %v", err)
		return nil, code.ProjectQuotaUpdateErr
	}

	return &pb.AddOnecProjectApplicationResp{}, nil
}

func (l *AddOnecProjectApplicationLogic) isQuotaInsufficient(in *pb.AddOnecProjectApplicationReq, quota *model.OnecProjectQuota) error {
	// 定义一个存储超出配额的资源信息的切片
	var exceeded []string

	// 检查每种资源是否超出配额
	if in.CpuLimit > (quota.CpuLimit - quota.CpuUsed) {
		exceeded = append(exceeded, fmt.Sprintf("CPU 配额不足 (申请: %v, 剩余: %v)", in.CpuLimit, quota.CpuLimit-quota.CpuUsed))
	}
	if in.MemoryLimit > (quota.MemoryLimit - quota.MemoryUsed) {
		exceeded = append(exceeded, fmt.Sprintf("内存配额不足 (申请: %v, 剩余: %v)", in.MemoryLimit, quota.MemoryLimit-quota.MemoryUsed))
	}
	if in.StorageLimit > (quota.StorageLimit - quota.StorageUsed) {
		exceeded = append(exceeded, fmt.Sprintf("存储配额不足 (申请: %v, 剩余: %v)", in.StorageLimit, quota.StorageLimit-quota.StorageUsed))
	}
	if in.PvcLimit > (quota.PvcLimit - quota.PvcUsed) {
		exceeded = append(exceeded, fmt.Sprintf("PVC 配额不足 (申请: %v, 剩余: %v)", in.PvcLimit, quota.PvcLimit-quota.PvcUsed))
	}
	if in.PodLimit > (quota.PodLimit - quota.PodUsed) {
		exceeded = append(exceeded, fmt.Sprintf("Pod 配额不足 (申请: %v, 剩余: %v)", in.PodLimit, quota.PodLimit-quota.PodUsed))
	}
	if in.SecretLimit > (quota.SecretLimit - quota.SecretUsed) {
		exceeded = append(exceeded, fmt.Sprintf("Secret 配额不足 (申请: %v, 剩余: %v)", in.SecretLimit, quota.SecretLimit-quota.SecretUsed))
	}
	if in.NodeportLimit > (quota.NodeportLimit - quota.NodeportUsed) {
		exceeded = append(exceeded, fmt.Sprintf("NodePort 配额不足 (申请: %v, 剩余: %v)", in.NodeportLimit, quota.NodeportLimit-quota.NodeportUsed))
	}
	if in.ConfigmapLimit > (quota.ConfigmapLimit - quota.ConfigmapUsed) {
		exceeded = append(exceeded, fmt.Sprintf("ConfigMap 配额不足 (申请: %v, 剩余: %v)", in.ConfigmapLimit, quota.ConfigmapLimit-quota.ConfigmapUsed))
	}
	if in.ServiceLimit > (quota.ServiceLimit - quota.ServiceUsed) {
		exceeded = append(exceeded, fmt.Sprintf("Service 配额不足 (申请: %v, 剩余: %v)", in.ServiceLimit, quota.ServiceLimit-quota.ServiceUsed))
	}

	// 如果有超出配额的资源，返回错误信息
	if len(exceeded) > 0 {
		return errorx.New(102199, fmt.Sprintf("资源配额不足: %s", strings.Join(exceeded, "; ")))
	}

	// 如果没有超出配额，返回 nil
	return nil
}

func (l *AddOnecProjectApplicationLogic) buildLabelsAndAnnotations(in *pb.AddOnecProjectApplicationReq, project *model.OnecProject) (map[string]string, map[string]string) {
	labels := map[string]string{
		"created_by": in.CreatedBy,
	}
	annotations := map[string]string{
		"project_name":       project.Name,
		"project_Identifier": project.Identifier,
		"cluster_uuid":       in.ClusterUuid,
		"project_desc":       project.Description,
	}
	return labels, annotations
}

func (l *AddOnecProjectApplicationLogic) createNamespace(nsClient core.NamespacesInterface, identifier string, labels, annotations map[string]string) error {
	namespace := kubeutils.CreateNamespace(identifier, kubeutils.WithLabels(labels), kubeutils.WithAnnotations(annotations))
	if _, err := nsClient.CreateNamespace(namespace); err != nil {
		l.Logger.Errorf("创建命名空间失败: %v", err)
		return err
	}
	return nil
}

func (l *AddOnecProjectApplicationLogic) createNamespaceQuota(nsClient core.NamespacesInterface, in *pb.AddOnecProjectApplicationReq, labels, annotations map[string]string) error {
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
	resourceQuota := kubeutils.CreateResourceQuota(fmt.Sprintf("%s-onec-default", in.Identifier), in.Identifier, limits, kubeutils.WithLabels(labels), kubeutils.WithAnnotations(annotations))
	if err := nsClient.SetResourceQuota(in.Identifier, resourceQuota); err != nil {
		l.Logger.Errorf("创建命名空间资源配额失败: %v", err)
		return err
	}
	return nil
}

func (l *AddOnecProjectApplicationLogic) insertApplicationData(in *pb.AddOnecProjectApplicationReq, labels, annotations map[string]string) error {
	application := &model.OnecProjectApplication{
		ProjectId:      in.ProjectId,
		ClusterUuid:    in.ClusterUuid,
		Identifier:     in.Identifier,
		AppCreateTime:  time.Now(),
		Name:           in.Name,
		Description:    in.Description,
		CpuLimit:       in.CpuLimit,
		MemoryLimit:    in.MemoryLimit,
		StorageLimit:   in.StorageLimit,
		PvcLimit:       in.PvcLimit,
		PodLimit:       in.PodLimit,
		NodeportLimit:  in.NodeportLimit,
		SecretLimit:    in.SecretLimit,
		ServiceLimit:   in.ServiceLimit,
		ConfigmapLimit: in.ConfigmapLimit,
		CreatedBy:      in.CreatedBy,
		UpdatedBy:      in.UpdatedBy,
		Status:         "Active",
	}
	if _, err := l.svcCtx.ProjectApplicationModel.Insert(l.ctx, application); err != nil {
		l.Logger.Errorf("插入应用数据失败: %v", err)
		return err
	}
	return nil
}

func (l *AddOnecProjectApplicationLogic) updateQuota(quota *model.OnecProjectQuota, in *pb.AddOnecProjectApplicationReq) error {
	quota.CpuUsed += in.CpuLimit
	quota.MemoryUsed += in.MemoryLimit
	quota.PodUsed += in.PodLimit
	quota.StorageUsed += in.StorageLimit
	quota.PvcUsed += in.PvcLimit
	quota.SecretUsed += in.SecretLimit
	quota.NodeportUsed += in.NodeportLimit
	quota.ConfigmapUsed += in.ConfigmapLimit
	quota.ServiceUsed += in.ServiceLimit
	quota.UpdatedBy = in.UpdatedBy
	if err := l.svcCtx.ProjectQuotaModel.Update(l.ctx, quota); err != nil {
		l.Logger.Errorf("更新资源配额失败: %v", err)
		return err
	}
	return nil
}

func (l *AddOnecProjectApplicationLogic) cleanUpNamespace(nsClient core.NamespacesInterface, namespace, quotaName string) {
	if err := nsClient.DeleteResourceQuota(namespace, quotaName); err != nil {
		l.Logger.Errorf("清理资源配额失败: %v", err)
	}
	if err := nsClient.DeleteNamespace(namespace); err != nil {
		l.Logger.Errorf("清理命名空间失败: %v", err)
	}
}
