package kubeutils

import (
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

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

// filterZeroValueResources 过滤掉零值的资源限制
func filterZeroValueResources(limits ResourceLimits) corev1.ResourceList {
	resourceList := corev1.ResourceList{}

	if limits.CPU != "" {
		resourceList[corev1.ResourceCPU] = resource.MustParse(limits.CPU)
	}
	if limits.Memory != "" {
		resourceList[corev1.ResourceMemory] = resource.MustParse(limits.Memory)
	}
	if limits.Pods != "" {
		resourceList[corev1.ResourcePods] = resource.MustParse(limits.Pods)
	}
	if limits.Services != "" {
		resourceList[corev1.ResourceServices] = resource.MustParse(limits.Services)
	}
	if limits.Secrets != "" {
		resourceList[corev1.ResourceSecrets] = resource.MustParse(limits.Secrets)
	}
	if limits.ConfigMaps != "" {
		resourceList[corev1.ResourceConfigMaps] = resource.MustParse(limits.ConfigMaps)
	}
	if limits.ServicesNodePorts != "" {
		resourceList[corev1.ResourceServicesNodePorts] = resource.MustParse(limits.ServicesNodePorts)
	}
	// 修复存储字段的资源名称
	if limits.Storage != "" {
		resourceList["requests.storage"] = resource.MustParse(limits.Storage)
	}
	if limits.PersistentVolumeClaims != "" {
		resourceList[corev1.ResourcePersistentVolumeClaims] = resource.MustParse(limits.PersistentVolumeClaims)
	}

	return resourceList
}

// CreateNamespace 创建一个 Namespace
func CreateNamespace(name string, options ...Option) *corev1.Namespace {
	namespace := &corev1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
			// 默认的标签和注解
			Labels:      defaultLabels,
			Annotations: defaultAnnotations,
		},
	}

	// 应用所有 Option
	for _, option := range options {
		option(namespace)
	}

	return namespace
}

// CreateResourceQuota 创建一个 ResourceQuota
func CreateResourceQuota(name, namespace string, limits ResourceLimits, options ...Option) *corev1.ResourceQuota {
	resourceQuota := &corev1.ResourceQuota{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
			// 默认的标签和注解
			Labels:      defaultLabels,
			Annotations: defaultAnnotations,
		},
		Spec: corev1.ResourceQuotaSpec{
			Hard: filterZeroValueResources(limits),
		},
	}

	// 应用所有 Option
	for _, option := range options {
		option(resourceQuota)
	}

	return resourceQuota
}
