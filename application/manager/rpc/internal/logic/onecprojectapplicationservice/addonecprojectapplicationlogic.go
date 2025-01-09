package onecprojectapplicationservicelogic

import (
	"context"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/internal/code"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/internal/shared"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/yanshicheng/kube-onec/application/manager/rpc/internal/svc"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/pb"

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

// -----------------------应用表，-----------------------
func (l *AddOnecProjectApplicationLogic) AddOnecProjectApplication(in *pb.AddOnecProjectApplicationReq) (*pb.AddOnecProjectApplicationResp, error) {
	name := "testsss"
	client, err := shared.GetK8sClient(l.ctx, l.svcCtx, "cee07e52-6516-4773-915d-0c762ed03d3e")
	if err != nil {
		l.Logger.Errorf("获取集群客户端异常: %v", err)
		return nil, code.GetClusterClientErr
	}

	ns := client.GetNamespaceClient()
	ns.CreateNamespace(&corev1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
	})
	ns.SetResourceQuota(name, &corev1.ResourceQuota{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test",
			Namespace: name,
		},
		Spec: corev1.ResourceQuotaSpec{
			Hard: corev1.ResourceList{
				corev1.ResourceCPU:                    resource.MustParse("100m"),
				corev1.ResourceMemory:                 resource.MustParse("100Mi"),
				corev1.ResourcePods:                   resource.MustParse("10"),
				corev1.ResourceStorage:                resource.MustParse("10Gi"),
				corev1.ResourceSecrets:                resource.MustParse("10"),
				corev1.ResourceServices:               resource.MustParse("10"),
				corev1.ResourceConfigMaps:             resource.MustParse("10"),
				corev1.ResourcePersistentVolumeClaims: resource.MustParse("10"),
				corev1.ResourceReplicationControllers: resource.MustParse("10"),
				corev1.ResourceQuotas:                 resource.MustParse("10"),
				corev1.ResourceServicesNodePorts:      resource.MustParse("10"),
				corev1.ResourceServicesLoadBalancers:  resource.MustParse("10"),
				corev1.ResourceLimitsCPU:              resource.MustParse("100m"),
				corev1.ResourceLimitsMemory:           resource.MustParse("100Mi"),
			},
		},
	})
	return &pb.AddOnecProjectApplicationResp{}, nil
}

func createResourceQuota(name, namespace string) *corev1.ResourceQuota {
	return &corev1.ResourceQuota{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
		},
		Spec: corev1.ResourceQuotaSpec{
			Hard: corev1.ResourceList{
				corev1.ResourceCPU:                    resource.MustParse("10"),   // 限制 10 CPU
				corev1.ResourceMemory:                 resource.MustParse("10Gi"), // 限制 10 Gi 内存
				corev1.ResourcePods:                   resource.MustParse("100"),  // 限制 100 个 Pod
				corev1.ResourceServices:               resource.MustParse("10"),   // 限制 10 个 Service
				corev1.ResourceSecrets:                resource.MustParse("10"),   // 限制 10 个 Secret
				corev1.ResourceConfigMaps:             resource.MustParse("10"),   // 限制 10 个 ConfigMap
				corev1.ResourceServicesNodePorts:      resource.MustParse("10"),   // 限制 10 个 NodePort
				corev1.ResourceStorage:                resource.MustParse("10Gi"), // 限制 10 Gi 存储
				corev1.ResourcePersistentVolumeClaims: resource.MustParse("20"),   // 限制 20 个 PVC
			},
		},
	}
}

func createNamespace(name string) *corev1.Namespace {
	return &corev1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: name, // 命名空间名称
			Labels: map[string]string{
				"project": "ikubeops", // 标签：project=ikubeops
			},
			Annotations: map[string]string{
				"Official website address": "https://www.ikubeops.com", // 注解：官网地址
				"作者":                       "yanshicheng",              // 注解：作者
			},
		},
	}
}
