package core

import (
	corev1 "k8s.io/api/core/v1"
	"time"
)

// NodesInterface 定义节点模块的接口
type NodesInterface interface {
	GetAllNodesInfo() ([]*NodeInfo, error)
	GetNodeInfo(nodeName string) (*NodeInfo, error)
	AddTaint(nodeName string, taint corev1.Taint) error
	AddLabel(nodeName, key, value string) error
	RemoveLabel(nodeName, key string) error
	RemoveTaint(nodeName string, taint corev1.Taint) error
	MarkUnschedulable(nodeName string) error
	ForceEvictAllPods(nodeName string) error
	EvictPod(namespace, podName string) error
}

type NodeInfo struct {
	NodeName string `json:"nodeName"`
	NodeUid  string `json:"nodeUid"`
	Status   string `json:"status"`
	Roles    string `json:"roles"`
	// 加入集群时间
	JoinTime time.Time `json:"joinTime"`
	// labels
	Labels map[string]string `json:"labels"`
	// annotations
	Annotations      map[string]string `json:"annotations"`
	PodCIDR          string            `json:"podCIDR"`
	Unschedulable    bool              `json:"unschedulable"` // 节点是否不可调度
	Taints           []Taint           `json:"taints"`
	NodeIp           string            `json:"nodeIp"`
	Os               string            `json:"os"`
	Cpu              int64             `json:"cpu"`
	Memory           int64             `json:"memory"`
	MaxPods          int64             `json:"macPods"`
	KernelVersion    string            `json:"kernelVersion"`
	ContainerRuntime string            `json:"containerRuntime"`
	KubeletVersion   string            `json:"kubeletVersion"`
	KubeletPort      int64             `json:"kubeletPort"`
	OperatingSystem  string            `json:"operatingSystem"`
	Architecture     string            `json:"architecture"`
}

type Taint struct {
	Key       string    `json:"key"`       // 污点键
	Value     string    `json:"value"`     // 污点值
	Effect    string    `json:"effect"`    // 污点效果
	TimeAdded time.Time `json:"timeAdded"` // 污点添加时间
}
