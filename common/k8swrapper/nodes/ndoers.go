package nodes

import (
	"context"
	"fmt"
	"github.com/yanshicheng/kube-onec/common/k8swrapper/core"
	"github.com/zeromicro/go-zero/core/logx"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/api/policy/v1beta1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"time"
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
	var nodeInfos []*core.NodeInfo
	// 获取所有节点信息
	options := metav1.ListOptions{}
	for {
		nodeList, err := n.client.GetClientset().CoreV1().Nodes().List(n.ctx, options)
		if err != nil {
			n.Logger.Errorf("获取节点信息失败: %v", err)
			return nil, err
		}

		for _, node := range nodeList.Items {
			nodeInfos = append(nodeInfos, generateNodeInfo(&node))
		}

		// 如果没有下一页，跳出循环
		if nodeList.Continue == "" {
			break
		}
		options.Continue = nodeList.Continue
	}
	return nodeInfos, nil
}

// GetNodeInfo 查询单个节点并返回对应的 NodeInfo
func (n *Nodes) GetNodeInfo(nodeName string) (*core.NodeInfo, error) {
	// 使用 Kubernetes 客户端获取指定名称的节点
	node, err := n.client.GetClientset().CoreV1().Nodes().Get(n.ctx, nodeName, metav1.GetOptions{})
	if err != nil {
		n.Logger.Errorf("获取节点 %s 信息失败: %v", nodeName, err)
		return nil, err
	}
	// 构造 NodeInfo 实例
	return generateNodeInfo(node), nil
}

// AddLabel 为节点添加标签
func (n *Nodes) AddLabel(nodeName, key, value string) error {
	return n.updateNode(nodeName, 3, func(node *corev1.Node) error {
		if node.Labels == nil {
			node.Labels = make(map[string]string)
		}

		// 如果标签已存在且值相同，直接返回
		if existingValue, exists := node.Labels[key]; exists && existingValue == value {
			n.Logger.Infof("节点 %s 已存在相同的标签: %s=%s，无需更新", nodeName, key, value)
			return nil
		}

		node.Labels[key] = value
		return nil
	})
}

// RemoveLabel 删除节点的指定标签
func (n *Nodes) RemoveLabel(nodeName, key string) error {
	return n.updateNode(nodeName, 3, func(node *corev1.Node) error {
		// 如果标签不存在，直接返回
		if _, exists := node.Labels[key]; !exists {
			n.Logger.Infof("节点 %s 不存在标签: %s，无需删除", nodeName, key)
			return nil
		}

		delete(node.Labels, key)
		return nil
	})
}

// AddTaint 添加污点
func (n *Nodes) AddTaint(nodeName string, taint corev1.Taint) error {
	return n.updateNode(nodeName, 3, func(node *corev1.Node) error {
		// 检查污点是否已存在
		for _, t := range node.Spec.Taints {
			if t.MatchTaint(&taint) {
				n.Logger.Infof("节点 %s 已存在污点: %v，无需更新", nodeName, taint)
				return nil
			}
		}

		node.Spec.Taints = append(node.Spec.Taints, taint)
		return nil
	})
}

// RemoveTaint 删除污点
func (n *Nodes) RemoveTaint(nodeName string, taint corev1.Taint) error {
	return n.updateNode(nodeName, 3, func(node *corev1.Node) error {
		updatedTaints := []corev1.Taint{}
		found := false

		for _, t := range node.Spec.Taints {
			if t.MatchTaint(&taint) {
				found = true
				continue
			}
			updatedTaints = append(updatedTaints, t)
		}

		if !found {
			n.Logger.Infof("节点 %s 不存在污点: %v，无需删除", nodeName, taint)
			return nil
		}

		node.Spec.Taints = updatedTaints
		return nil
	})
}

func (n *Nodes) AddAnnotation(nodeName string, key, value string) error {
	return n.updateNode(nodeName, 3, func(node *corev1.Node) error {
		if node.Annotations == nil {
			node.Annotations = make(map[string]string)
		}

		// 如果注解已存在且值未变化，则无需更新
		if existingValue, exists := node.Annotations[key]; exists && existingValue == value {
			n.Logger.Infof("节点 %s 已存在相同的注解: %s=%s，无需更新", nodeName, key, value)
			return nil
		}

		node.Annotations[key] = value
		return nil
	})
}

func (n *Nodes) RemoveAnnotation(nodeName string, key string) error {
	return n.updateNode(nodeName, 3, func(node *corev1.Node) error {
		// 如果注解不存在，直接返回
		if _, exists := node.Annotations[key]; !exists {
			n.Logger.Infof("节点 %s 上不存在注解: %s，无需删除", nodeName, key)
			return nil
		}

		delete(node.Annotations, key)
		return nil
	})
}

// ForbidScheduled 禁止调度
func (n *Nodes) ForbidScheduled(nodeName string) error {
	return n.updateNode(nodeName, 3, func(node *corev1.Node) error {
		// 如果节点已经是不可调度状态，直接返回
		if node.Spec.Unschedulable {
			n.Logger.Infof("节点 %s 已经是不可调度状态，无需更新", nodeName)
			return nil
		}

		// 设置为不可调度
		node.Spec.Unschedulable = true
		return nil
	})
}

// EnableScheduled 允许调度
func (n *Nodes) EnableScheduled(nodeName string) error {
	return n.updateNode(nodeName, 3, func(node *corev1.Node) error {
		// 如果节点已经是可调度状态，直接返回
		if !node.Spec.Unschedulable {
			n.Logger.Infof("节点 %s 已经是可调度状态，无需更新", nodeName)
			return nil
		}

		// 设置为可调度
		node.Spec.Unschedulable = false
		return nil
	})
}

// ForceEvictAllPods 强制驱逐指定节点上的所有 Pod（忽略 PDB）
func (n *Nodes) ForceEvictAllPods(nodeName string) error {
	// 获取节点上所有的 Pod
	pods, err := n.client.GetClientset().CoreV1().Pods("").List(n.ctx, metav1.ListOptions{
		FieldSelector: "spec.nodeName=" + nodeName,
	})
	if err != nil {
		n.Logger.Errorf("获取节点 %s 上的 Pod 列表失败: %v", nodeName, err)
		return err
	}

	// 遍历每个 Pod，尝试驱逐
	for _, pod := range pods.Items {
		if err := n.evictOrDeletePod(&pod); err != nil {
			n.Logger.Errorf("处理 Pod %s/%s 失败: %v", pod.Namespace, pod.Name, err)
			return err
		}
	}

	n.Logger.Infof("成功强制驱逐节点 %s 上的所有 Pod", nodeName)
	return nil
}

// evictOrDeletePod 封装驱逐或删除 Pod 的逻辑
func (n *Nodes) evictOrDeletePod(pod *corev1.Pod) error {
	// 创建驱逐请求
	eviction := &v1beta1.Eviction{
		ObjectMeta: metav1.ObjectMeta{
			Name:      pod.Name,
			Namespace: pod.Namespace,
		},
		DeleteOptions: &metav1.DeleteOptions{
			GracePeriodSeconds: int64Ptr(0), // 设置宽限期为 0 秒
		},
	}

	// 尝试驱逐 Pod
	err := n.client.GetClientset().CoreV1().Pods(pod.Namespace).Evict(context.TODO(), eviction)
	if err == nil {
		n.Logger.Infof("成功驱逐 Pod %s/%s", pod.Namespace, pod.Name)
		return nil
	}

	// 如果驱逐失败，尝试直接删除 Pod
	n.Logger.Errorf("驱逐 Pod %s/%s 失败，尝试直接删除: %v", pod.Namespace, pod.Name, err)
	err = n.client.GetClientset().CoreV1().Pods(pod.Namespace).Delete(context.TODO(), pod.Name, metav1.DeleteOptions{
		GracePeriodSeconds: int64Ptr(0),
	})
	if err != nil {
		n.Logger.Errorf("直接删除 Pod %s/%s 失败: %v", pod.Namespace, pod.Name, err)
		return err
	}

	n.Logger.Infof("成功直接删除 Pod %s/%s", pod.Namespace, pod.Name)
	return nil
}

// updateNode 封装更新节点的逻辑，包含重试机制
func (n *Nodes) updateNode(nodeName string, maxRetries int, updateFunc func(*corev1.Node) error) error {
	for i := 0; i < maxRetries; i++ {
		node, err := n.GetNodeWithRetry(nodeName, 3) // 重用通用获取节点方法
		if err != nil {
			return err
		}

		// 调用自定义更新逻辑
		if err := updateFunc(node); err != nil {
			n.Logger.Errorf("更新节点 %s 失败: %v", nodeName, err)
			return err
		}

		// 提交更新
		_, err = n.client.GetClientset().CoreV1().Nodes().Update(n.ctx, node, metav1.UpdateOptions{})
		if err == nil {
			n.Logger.Infof("成功更新节点 %s", nodeName)
			return nil
		}

		// 如果不是冲突错误，直接返回
		if !errors.IsConflict(err) {
			n.Logger.Errorf("更新节点 %s 失败: %v", nodeName, err)
			return err
		}

		// 冲突错误，重试
		n.Logger.Errorf("更新节点 %s 发生冲突，正在重试（%d/%d）", nodeName, i+1, maxRetries)
		time.Sleep(100 * time.Millisecond) // 延迟后重试
	}

	return fmt.Errorf("更新节点 %s 失败，已超出最大重试次数", nodeName)
}

func (n *Nodes) GetNodeWithRetry(nodeName string, maxRetries int) (*corev1.Node, error) {
	for i := 0; i < maxRetries; i++ { // 重试逻辑
		node, err := n.client.GetClientset().CoreV1().Nodes().Get(n.ctx, nodeName, metav1.GetOptions{})
		if err == nil {
			return node, nil
		}

		// 如果不是 NotFound 错误且不是冲突，则直接返回
		if !errors.IsNotFound(err) && !errors.IsConflict(err) {
			n.Logger.Errorf("获取节点 %s 失败: %v", nodeName, err)
			return nil, err
		}

		// 如果是冲突或暂时性错误，则重试
		n.Logger.Errorf("获取节点 %s 失败，正在重试（%d/%d）: %v", nodeName, i+1, maxRetries, err)
		time.Sleep(100 * time.Millisecond) // 延迟后重试
	}

	return nil, fmt.Errorf("获取节点 %s 失败，已超出最大重试次数", nodeName)
}
