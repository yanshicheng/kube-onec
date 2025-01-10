package core

import (
	corev1 "k8s.io/api/core/v1"
)

// NamespacesInterface 定义命名空间模块的接口
type NamespacesInterface interface {
	// 创建命名空间
	CreateNamespace(namespace *corev1.Namespace) (*corev1.Namespace, error)

	// 删除命名空间
	DeleteNamespace(namespaceName string) error
	// 查询名称空间是否存在
	NamespaceExist(namespaceName string) bool
	// 获取所有命名空间信息
	GetAllNamespaces() ([]*corev1.Namespace, error)

	// 获取指定命名空间的详细信息（类似于 describe）
	GetNamespaceDetails(namespaceName string) (*corev1.Namespace, error)

	// 设置命名空间的 ResourceQuota
	SetResourceQuota(namespaceName string, quota *corev1.ResourceQuota) error

	// 获取命名空间的 ResourceQuota
	GetResourceQuota(namespaceName, quotaName string) (*corev1.ResourceQuota, error)

	// 删除命名空间的 ResourceQuota
	DeleteResourceQuota(namespaceName, quotaName string) error

	// 修改命名空间的 ResourceQuota
	UpdateResourceQuota(namespaceName string, quota *corev1.ResourceQuota) error

	// 添加标签到命名空间
	AddLabel(namespaceName, key, value string) error

	// 移除命名空间的标签
	RemoveLabel(namespaceName, key string) error

	// 添加注解到命名空间
	AddAnnotation(namespaceName, key, value string) error

	// 移除命名空间的注解
	RemoveAnnotation(namespaceName, key string) error

	// 获取命名空间并带有重试机制
	GetNamespaceWithRetry(namespaceName string, maxRetries int) (*corev1.Namespace, error)
}
