package core

import (
	"k8s.io/client-go/kubernetes"
)

// K8sClientInterface 定义一个接口，用于统一管理 Kubernetes 客户端的子模块
type K8sClientInterface interface {
	GetCluster() ClusterInterface
	GetNodes() NodesInterface
	GetClientset() *kubernetes.Clientset
}
