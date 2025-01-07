package nodes

import (
	"context"
	"github.com/yanshicheng/kube-onec/common/k8swrapper/core"
	"github.com/zeromicro/go-zero/core/logx"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/api/policy/v1beta1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"net"
	"strings"
)

// Nodes 实现 core.NodesInterface
type Nodes struct {
	client core.K8sClientInterface
	ctx    context.Context
	Logger logx.Logger
}

// NewNodes 初始化 Nodes
func NewNodes(ctx context.Context, client core.K8sClientInterface) *Nodes {
	return &Nodes{client: client, ctx: ctx, Logger: logx.WithContext(ctx)}
}

func (n *Nodes) GetAllNodesInfo() ([]*core.NodeInfo, error) {
	// 获取所有节点信息
	nodes, err := n.client.GetClientset().CoreV1().Nodes().List(n.ctx, metav1.ListOptions{})
	if err != nil {
		n.Logger.Errorf("获取节点信息失败: %v", err)
		return nil, err
	}

	// 定义存储 NodeInfo 的切片
	var nodeInfos []*core.NodeInfo

	// 遍历所有节点
	for _, node := range nodes.Items {
		// 构造 NodeInfo 实例
		// 将 NodeInfo 添加到列表
		nodeInfos = append(nodeInfos, generateNodeInfo(&node))
	}

	return nodeInfos, nil
}

// GetNodeInfo 查询单个节点并返回对应的 NodeInfo
func (n *Nodes) GetNodeInfo(nodeName string) (*core.NodeInfo, error) {
	// 使用 Kubernetes 客户端获取指定名称的节点
	node, err := n.client.GetClientset().CoreV1().Nodes().Get(n.ctx, nodeName, metav1.GetOptions{})
	if err != nil {
		n.Logger.Errorf("获取节点信息失败: %v", err)
		return nil, err
	}

	// 构造 NodeInfo 实例
	return generateNodeInfo(node), nil
}

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

// AddLabel 为节点添加标签
func (n *Nodes) AddLabel(nodeName, key, value string) error {
	node, err := n.client.GetClientset().CoreV1().Nodes().Get(n.ctx, nodeName, metav1.GetOptions{})
	if err != nil {
		n.Logger.Errorf("获取节点 %s 失败: %v", nodeName, err)
		return err
	}

	// 添加或更新标签
	if node.Labels == nil {
		node.Labels = make(map[string]string)
	}
	node.Labels[key] = value

	_, err = n.client.GetClientset().CoreV1().Nodes().Update(n.ctx, node, metav1.UpdateOptions{})
	if err != nil {
		n.Logger.Errorf("更新节点 %s 标签失败: %v", nodeName, err)
		return err
	}
	n.Logger.Infof("成功为节点 %s 添加标签: %s=%s", nodeName, key, value)
	return nil
}

// RemoveLabel 删除节点的指定标签
func (n *Nodes) RemoveLabel(nodeName, key string) error {
	node, err := n.client.GetClientset().CoreV1().Nodes().Get(n.ctx, nodeName, metav1.GetOptions{})
	if err != nil {
		n.Logger.Errorf("获取节点 %s 失败: %v", nodeName, err)
		return err
	}

	// 删除标签
	delete(node.Labels, key)

	_, err = n.client.GetClientset().CoreV1().Nodes().Update(n.ctx, node, metav1.UpdateOptions{})
	if err != nil {
		n.Logger.Errorf("删除节点 %s 标签失败: %v", nodeName, err)
		return err
	}
	n.Logger.Infof("成功删除节点 %s 的标签: %s", nodeName, key)
	return nil
}

// AddTaint 添加污点
func (n *Nodes) AddTaint(nodeName string, taint corev1.Taint) error {
	node, err := n.client.GetClientset().CoreV1().Nodes().Get(n.ctx, nodeName, metav1.GetOptions{})
	if err != nil {
		n.Logger.Errorf("获取节点 %s 失败: %v", nodeName, err)
		return err
	}

	// 添加污点
	node.Spec.Taints = append(node.Spec.Taints, taint)

	_, err = n.client.GetClientset().CoreV1().Nodes().Update(n.ctx, node, metav1.UpdateOptions{})
	if err != nil {
		n.Logger.Errorf("为节点 %s 添加污点失败: %v", nodeName, err)
		return err
	}
	n.Logger.Infof("成功为节点 %s 添加污点: %v", nodeName, taint)
	return nil
}

// RemoveTaint 删除污点
func (n *Nodes) RemoveTaint(nodeName string, taint corev1.Taint) error {
	node, err := n.client.GetClientset().CoreV1().Nodes().Get(n.ctx, nodeName, metav1.GetOptions{})
	if err != nil {
		n.Logger.Errorf("获取节点 %s 失败: %v", nodeName, err)
		return err
	}

	// 删除污点
	var updatedTaints []corev1.Taint
	for _, t := range node.Spec.Taints {
		if t.Key != taint.Key || t.Effect != taint.Effect {
			updatedTaints = append(updatedTaints, t)
		}
	}
	node.Spec.Taints = updatedTaints

	_, err = n.client.GetClientset().CoreV1().Nodes().Update(n.ctx, node, metav1.UpdateOptions{})
	if err != nil {
		n.Logger.Errorf("删除节点 %s 污点失败: %v", nodeName, err)
		return err
	}
	n.Logger.Infof("成功删除节点 %s 的污点: %v", nodeName, taint)
	return nil
}

// MarkUnschedulable 禁止调度
func (n *Nodes) MarkUnschedulable(nodeName string) error {
	node, err := n.client.GetClientset().CoreV1().Nodes().Get(n.ctx, nodeName, metav1.GetOptions{})
	if err != nil {
		n.Logger.Errorf("获取节点 %s 失败: %v", nodeName, err)
		return err
	}

	node.Spec.Unschedulable = true

	_, err = n.client.GetClientset().CoreV1().Nodes().Update(n.ctx, node, metav1.UpdateOptions{})
	if err != nil {
		n.Logger.Errorf("禁止节点 %s 调度失败: %v", nodeName, err)
		return err
	}
	n.Logger.Infof("成功将节点 %s 设置为不可调度", nodeName)
	return nil
}

// UnmarkUnschedulable 取消禁止调度
func (n *Nodes) UnmarkUnschedulable(nodeName string) error {
	node, err := n.client.GetClientset().CoreV1().Nodes().Get(n.ctx, nodeName, metav1.GetOptions{})
	if err != nil {
		n.Logger.Errorf("获取节点 %s 失败: %v", nodeName, err)
		return err
	}

	node.Spec.Unschedulable = false

	_, err = n.client.GetClientset().CoreV1().Nodes().Update(n.ctx, node, metav1.UpdateOptions{})
	if err != nil {
		n.Logger.Errorf("取消节点 %s 的不可调度状态失败: %v", nodeName, err)
		return err
	}
	n.Logger.Infof("成功将节点 %s 设置为可调度", nodeName)
	return nil
}

// ForceEvictAllPods 强制驱逐指定节点上的所有 Pod（忽略 PDB）
func (n *Nodes) ForceEvictAllPods(nodeName string) error {
	// 获取指定节点上的所有 Pod
	pods, err := n.client.GetClientset().CoreV1().Pods("").List(n.ctx, metav1.ListOptions{
		FieldSelector: "spec.nodeName=" + nodeName,
	})
	if err != nil {
		n.Logger.Errorf("获取节点 %s 上的 Pod 列表失败: %v", nodeName, err)
		return err
	}

	// 遍历 Pod 列表并逐一强制驱逐
	for _, pod := range pods.Items {
		// 创建强制 Eviction 对象
		eviction := &v1beta1.Eviction{
			ObjectMeta: metav1.ObjectMeta{
				Name:      pod.Name,
				Namespace: pod.Namespace,
			},
			DeleteOptions: &metav1.DeleteOptions{
				GracePeriodSeconds: int64Ptr(0), // 设置宽限期为 0 秒，立即删除
			},
		}

		// 调用 Evict 方法强制驱逐 Pod
		err := n.client.GetClientset().CoreV1().Pods(pod.Namespace).Evict(context.TODO(), eviction)
		if err != nil {
			// 如果强制驱逐失败，尝试直接删除 Pod
			n.Logger.Errorf("强制驱逐 Pod %s/%s 失败，尝试直接删除: %v", pod.Namespace, pod.Name, err)

			err := n.client.GetClientset().CoreV1().Pods(pod.Namespace).Delete(context.TODO(), pod.Name, metav1.DeleteOptions{
				GracePeriodSeconds: int64Ptr(0), // 再次确保直接删除
			})
			if err != nil {
				n.Logger.Errorf("直接删除 Pod %s/%s 失败: %v", pod.Namespace, pod.Name, err)
				return err
			}
			n.Logger.Infof("成功直接删除 Pod %s/%s", pod.Namespace, pod.Name)
		} else {
			n.Logger.Infof("成功强制驱逐 Pod %s/%s", pod.Namespace, pod.Name)
		}
	}

	n.Logger.Infof("成功强制驱逐节点 %s 上的所有 Pod", nodeName)
	return nil
}

// 辅助函数：返回 int64 指针
func int64Ptr(i int64) *int64 {
	return &i
}

// EvictPod 驱逐节点上的 Pod
func (n *Nodes) EvictPod(namespace, podName string) error {

	eviction := &v1beta1.Eviction{
		ObjectMeta: metav1.ObjectMeta{
			Name:      podName,
			Namespace: namespace,
		},
	}

	err := n.client.GetClientset().CoreV1().Pods(namespace).Evict(n.ctx, eviction)
	if err != nil {
		n.Logger.Errorf("驱逐 Pod %s/%s 失败: %v", namespace, podName, err)
		return err
	}
	n.Logger.Infof("成功驱逐 Pod %s/%s", namespace, podName)
	return nil
}

// 方法实现细节
// 添加标签和删除标签
//
// 直接操作节点的 Labels 字段，修改后调用 Update 方法更新节点。
// 添加污点和删除污点
//
// 使用污点（Taints）字段，添加时直接 append，删除时过滤出不需要删除的污点。
// 禁止调度和取消禁止调度
//
// 操作节点的 Unschedulable 字段，将其设置为 true 或 false。
// 驱逐 Pod
//
// 调用 Kubernetes 提供的 Evict 方法，需要传递 namespace 和 podName。
// 测试示例
// 以下是如何调用这些方法的示例：
//
// go
// 复制代码
//
//	func main() {
//		nodeManager := &Nodes{/* 初始化 */}
//
//		// 添加标签
//		err := nodeManager.AddLabel("node1", "env", "production")
//		if err != nil {
//			fmt.Printf("添加标签失败: %v\n", err)
//		}
//
//		// 删除标签
//		err = nodeManager.RemoveLabel("node1", "env")
//		if err != nil {
//			fmt.Printf("删除标签失败: %v\n", err)
//		}
//
//		// 添加污点
//		taint := v1.Taint{
//			Key:    "key1",
//			Value:  "value1",
//			Effect: v1.TaintEffectNoSchedule,
//		}
//		err = nodeManager.AddTaint("node1", taint)
//		if err != nil {
//			fmt.Printf("添加污点失败: %v\n", err)
//		}
//
//		// 删除污点
//		err = nodeManager.RemoveTaint("node1", taint)
//		if err != nil {
//			fmt.Printf("删除污点失败: %v\n", err)
//		}
//
//		// 禁止调度
//		err = nodeManager.MarkUnschedulable("node1")
//		if err != nil {
//			fmt.Printf("禁止调度失败: %v\n", err)
//		}
//
//		// 取消禁止调度
//		err = nodeManager.UnmarkUnschedulable("node1")
//		if err != nil {
//			fmt.Printf("取消禁止调度失败: %v\n", err)
//		}
//
//		// 驱逐 Pod
//		err = nodeManager.EvictPod("default", "my-pod")
//		if err != nil {
//			fmt.Printf("驱逐 Pod 失败: %v\n", err)
//		}
//	}
//
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
