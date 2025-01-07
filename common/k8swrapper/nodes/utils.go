package nodes

import (
	"github.com/yanshicheng/kube-onec/common/k8swrapper/core"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	"net"
	"strings"
)

// getNodeTaints
func getNodeTaints(taint []corev1.Taint) []core.Taint {
	taints := make([]core.Taint, len(taint))
	for i, taint := range taint {
		taints[i] = core.Taint{
			Key:       taint.Key,
			Value:     taint.Value,
			Effect:    string(taint.Effect),
			TimeAdded: taint.TimeAdded.Time, // Kubernetes API 没有提供 Taint 添加时间
		}
	}
	return taints
}

// getNodeStatus 从节点的条件和调度状态中提取综合状态信息
func getNodeStatus(node *corev1.Node) string {
	var status string

	// 遍历节点的条件，查找 "Ready" 条件
	for _, condition := range node.Status.Conditions {
		if condition.Type == corev1.NodeReady {
			if condition.Status == corev1.ConditionTrue {
				status = "Ready"
			} else {
				status = "NotReady"
			}
			break
		}
	}

	// 检查节点是否被标记为不可调度
	if node.Spec.Unschedulable {
		if status != "" {
			status += ",SchedulingDisabled"
		} else {
			status = "SchedulingDisabled"
		}
	}

	return status
}

// getNodeRoles 从节点的标签中提取角色信息
func getNodeRoles(labels map[string]string) string {
	// 定义角色标签的前缀
	roleLabelPrefixes := []string{
		"node-role.kubernetes.io/", // 现代 Kubernetes 标签
		"kubernetes.io/role/",      // 自定义角色标签的前缀
	}

	// 使用 map 来存储角色，避免重复
	roleSet := make(map[string]struct{})

	// 遍历所有标签，处理键和值
	for key, value := range labels {
		// 处理基于前缀的标签
		for _, prefix := range roleLabelPrefixes {
			if strings.HasPrefix(key, prefix) {
				role := strings.TrimPrefix(key, prefix)
				if role != "" {
					roleSet[role] = struct{}{}
				}
			}
		}

		// 特殊处理 "kubernetes.io/role" 标签的值
		if key == "kubernetes.io/role" && value != "" {
			roleSet[value] = struct{}{}
		}
	}

	// 将 map 转换为切片
	var roles []string
	for role := range roleSet {
		roles = append(roles, role)
	}

	// 如果没有明确的角色标签，默认赋予 "node" 角色
	if len(roles) == 0 {
		roles = append(roles, "node")
	}

	// 将切片转换为以逗号分隔的字符串
	return strings.Join(roles, ",")
}

// 辅助函数：返回 int64 指针
func int64Ptr(i int64) *int64 {
	return &i
}

// getNodeAddresses 过滤并返回合法的 IPv4 或 IPv6 地址，返回逗号分隔的字符串
func getNodeAddresses(addresses []corev1.NodeAddress) string {
	var result []string

	for _, addr := range addresses {
		// 检查 Address 是否为合法的 IPv4 或 IPv6
		if net.ParseIP(addr.Address) != nil {
			result = append(result, addr.Address)
		}
	}

	// 将地址切片转换为逗号分隔的字符串
	return strings.Join(result, ",")
}

// 获取内存资源的整数值（以 Mi 为单位）
func getQuantityValueGi(q *resource.Quantity) float64 {
	value, ok := q.AsInt64()
	if !ok {
		return 0
	} // 以字节为单位
	return float64(value) / (1024 * 1024 * 1024) // 转换为 Gi
}

// 获取资源的整数值（如 CPU 核心数）
func getQuantityValueFloat(q *resource.Quantity) float64 {
	value, ok := q.AsInt64()
	if !ok {
		return 0
	}
	return float64(value)
}
func getQuantityValueInt(q *resource.Quantity) int64 {
	value, ok := q.AsInt64()
	if !ok {
		return 0
	}
	return value
}

// 生成 nodeInfo
func generateNodeInfo(node *corev1.Node) *core.NodeInfo {
	return &core.NodeInfo{
		NodeName:         node.Name,
		NodeUid:          string(node.UID),
		Status:           getNodeStatus(node),
		Roles:            getNodeRoles(node.Labels),
		Memory:           getQuantityValueGi(node.Status.Capacity.Memory()),
		MaxPods:          getQuantityValueInt(node.Status.Capacity.Pods()),
		Cpu:              getQuantityValueInt(node.Status.Capacity.Cpu()),
		JoinTime:         node.CreationTimestamp.Time,
		Labels:           node.Labels,
		Annotations:      node.Annotations,
		PodCIDR:          node.Spec.PodCIDR,
		Unschedulable:    node.Spec.Unschedulable,
		NodeIp:           getNodeAddresses(node.Status.Addresses),
		Taints:           getNodeTaints(node.Spec.Taints),
		Os:               node.Status.NodeInfo.OSImage,
		KernelVersion:    node.Status.NodeInfo.KernelVersion,
		ContainerRuntime: node.Status.NodeInfo.ContainerRuntimeVersion,
		KubeletVersion:   node.Status.NodeInfo.KubeletVersion,
		KubeletPort:      int64(node.Status.DaemonEndpoints.KubeletEndpoint.Port), // 默认 Kubelet 端口
		OperatingSystem:  node.Status.NodeInfo.OperatingSystem,
		Architecture:     node.Status.NodeInfo.Architecture,
	}
}
