package core

import (
	corev1 "k8s.io/api/core/v1"
	"time"
)

// NodesInterface 定义节点模块的接口
type NodesInterface interface {
	GetAllNodesInfo() ([]*NodeInfo, error)          // 获取所有节点信息
	GetNodeInfo(nodeName string) (*NodeInfo, error) // 获取指定节点信息

	// 污点相关
	AddTaint(nodeName string, taint corev1.Taint) error    // 添加污点
	RemoveTaint(nodeName string, taint corev1.Taint) error // 移除污点

	AddLabel(nodeName, key, value string) error // 添加标签
	RemoveLabel(nodeName, key string) error     // 移除标签

	// 调度相关
	ForbidScheduled(nodeName string) error // 标记节点不可调度
	EnableScheduled(nodeName string) error // 启用调度

	// 驱逐当前节点所有pod
	ForceEvictAllPods(nodeName string) error // 强制驱逐所有Pod

	// 注解相关
	AddAnnotation(nodeName string, key, value string) error
	RemoveAnnotation(nodeName string, key string) error

	GetNodeWithRetry(nodeName string, maxRetries int) (*corev1.Node, error)
}

type NodeInfo struct {
	NodeName string `json:"nodeName"`
	NodeUid  string `json:"nodeUid"`
	Status   string `json:"status"`
	Roles    string `json:"roles"`
	// 加入集群时间
	JoinTime         time.Time `json:"joinTime"`
	PodCIDR          string    `json:"podCIDR"`
	Unschedulable    bool      `json:"unschedulable"` // 节点是否不可调度
	NodeIp           string    `json:"nodeIp"`
	Os               string    `json:"os"`
	Cpu              int64     `json:"cpu"`
	Memory           float64   `json:"memory"`
	MaxPods          int64     `json:"macPods"`
	KernelVersion    string    `json:"kernelVersion"`
	ContainerRuntime string    `json:"containerRuntime"`
	KubeletVersion   string    `json:"kubeletVersion"`
	KubeletPort      int64     `json:"kubeletPort"`
	OperatingSystem  string    `json:"operatingSystem"`
	Architecture     string    `json:"architecture"`

	Labels      map[string]string `json:"labels"`
	Annotations map[string]string `json:"annotations"`
	Taints      []Taint           `json:"taints"`
}

type Taint struct {
	Key       string    `json:"key"`       // 污点键
	Value     string    `json:"value"`     // 污点值
	Effect    string    `json:"effect"`    // 污点效果
	TimeAdded time.Time `json:"timeAdded"` // 污点添加时间
}
