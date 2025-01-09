package namespace

import (
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func CreateNamespace(name string, options ...func(*corev1.Namespace)) *corev1.Namespace {
	// 默认的命名空间模板
	namespace := &corev1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: name, // 命名空间名称
			Labels: map[string]string{
				"project": "ikubeops", // 默认标签
			},
			Annotations: map[string]string{
				"Official website address": "https://www.ikubeops.com", // 默认注解：官网地址
				"作者":                       "yanshicheng",              // 默认注解：作者
			},
		},
	}

	// 应用传递的选项函数
	for _, option := range options {
		option(namespace)
	}

	return namespace
}

// Option 函数用于自定义 Labels
func WithLabels(labels map[string]string) func(*corev1.Namespace) {
	return func(namespace *corev1.Namespace) {
		if namespace.ObjectMeta.Labels == nil {
			namespace.ObjectMeta.Labels = make(map[string]string)
		}
		for key, value := range labels {
			namespace.ObjectMeta.Labels[key] = value
		}
	}
}

// Option 函数用于自定义 Annotations
func WithAnnotations(annotations map[string]string) func(*corev1.Namespace) {
	return func(namespace *corev1.Namespace) {
		if namespace.ObjectMeta.Annotations == nil {
			namespace.ObjectMeta.Annotations = make(map[string]string)
		}
		for key, value := range annotations {
			namespace.ObjectMeta.Annotations[key] = value
		}
	}
}

// ResourceLimits 定义资源限制的结构体
type ResourceLimits struct {
	CPU                    string
	Memory                 string
	Pods                   string
	Services               string
	Secrets                string
	ConfigMaps             string
	ServicesNodePorts      string
	Storage                string
	PersistentVolumeClaims string
}

// createResourceQuota 创建 ResourceQuota，支持传递 Labels 和 Annotations
func createResourceQuota(name, namespace string, limits ResourceLimits, options ...func(*corev1.ResourceQuota)) *corev1.ResourceQuota {
	// 构造 ResourceQuota
	resourceQuota := &corev1.ResourceQuota{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
		},
		Spec: corev1.ResourceQuotaSpec{
			Hard: corev1.ResourceList{
				corev1.ResourceCPU:                    resource.MustParse(limits.CPU),
				corev1.ResourceMemory:                 resource.MustParse(limits.Memory),
				corev1.ResourcePods:                   resource.MustParse(limits.Pods),
				corev1.ResourceServices:               resource.MustParse(limits.Services),
				corev1.ResourceSecrets:                resource.MustParse(limits.Secrets),
				corev1.ResourceConfigMaps:             resource.MustParse(limits.ConfigMaps),
				corev1.ResourceServicesNodePorts:      resource.MustParse(limits.ServicesNodePorts),
				corev1.ResourceStorage:                resource.MustParse(limits.Storage),
				corev1.ResourcePersistentVolumeClaims: resource.MustParse(limits.PersistentVolumeClaims),
			},
		},
	}

	// 应用传递的选项函数
	for _, option := range options {
		option(resourceQuota)
	}

	return resourceQuota
}
