// namespace/namespace.go

package namespace

import (
	"context"
	"fmt"
	"time"

	"github.com/yanshicheng/kube-onec/common/k8swrapper/core"
	"github.com/zeromicro/go-zero/core/logx"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Namespace 实现 core.NamespacesInterface
type Namespace struct {
	client core.K8sClientInterface
	ctx    context.Context
	Logger logx.Logger
}

// NewNamespace 初始化 Namespace
func NewNamespace(ctx context.Context, client core.K8sClientInterface) *Namespace {
	return &Namespace{
		client: client,
		ctx:    ctx,
		Logger: logx.WithContext(ctx),
	}
}

// CreateNamespace 创建命名空间
func (n *Namespace) CreateNamespace(namespace *corev1.Namespace) (*corev1.Namespace, error) {
	createdNamespace, err := n.client.GetClientset().CoreV1().Namespaces().Create(n.ctx, namespace, metav1.CreateOptions{})
	if err != nil {
		n.Logger.Errorf("创建命名空间 %s 失败: %v", namespace.Name, err)
		return nil, err
	}
	n.Logger.Infof("成功创建命名空间 %s", namespace.Name)
	return createdNamespace, nil
}

// DeleteNamespace 删除命名空间
func (n *Namespace) DeleteNamespace(namespaceName string) error {
	err := n.client.GetClientset().CoreV1().Namespaces().Delete(n.ctx, namespaceName, metav1.DeleteOptions{})
	if err != nil {
		n.Logger.Errorf("删除命名空间 %s 失败: %v", namespaceName, err)
		return err
	}
	n.Logger.Infof("成功删除命名空间 %s", namespaceName)
	return nil
}

// GetAllNamespaces 获取所有命名空间信息
func (n *Namespace) GetAllNamespaces() ([]*corev1.Namespace, error) {
	var namespaces []*corev1.Namespace
	options := metav1.ListOptions{}
	for {
		nsList, err := n.client.GetClientset().CoreV1().Namespaces().List(n.ctx, options)
		if err != nil {
			n.Logger.Errorf("获取所有命名空间失败: %v", err)
			return nil, err
		}

		for _, ns := range nsList.Items {
			namespaces = append(namespaces, &ns)
		}

		if nsList.Continue == "" {
			break
		}
		options.Continue = nsList.Continue
	}
	n.Logger.Infof("成功获取所有命名空间，共 %d 个", len(namespaces))
	return namespaces, nil
}

// GetNamespaceDetails 获取指定命名空间的详细信息
func (n *Namespace) GetNamespaceDetails(namespaceName string) (*corev1.Namespace, error) {
	ns, err := n.client.GetClientset().CoreV1().Namespaces().Get(n.ctx, namespaceName, metav1.GetOptions{})
	if err != nil {
		n.Logger.Errorf("获取命名空间 %s 详情失败: %v", namespaceName, err)
		return nil, err
	}
	n.Logger.Infof("成功获取命名空间 %s 的详情", namespaceName)
	return ns, nil
}

// SetResourceQuota 为命名空间设置 ResourceQuota
func (n *Namespace) SetResourceQuota(namespaceName string, quota *corev1.ResourceQuota) error {
	existingQuota, err := n.GetResourceQuota(namespaceName, quota.Name)
	if err != nil && !errors.IsNotFound(err) {
		return err
	}

	if existingQuota != nil {
		n.Logger.Infof("ResourceQuota %s 已存在于命名空间 %s，正在更新", quota.Name, namespaceName)
		return n.UpdateResourceQuota(namespaceName, quota)
	}

	_, err = n.client.GetClientset().CoreV1().ResourceQuotas(namespaceName).Create(n.ctx, quota, metav1.CreateOptions{})
	if err != nil {
		n.Logger.Errorf("为命名空间 %s 创建 ResourceQuota %s 失败: %v", namespaceName, quota.Name, err)
		return err
	}
	n.Logger.Infof("成功为命名空间 %s 创建 ResourceQuota %s", namespaceName, quota.Name)
	return nil
}

// GetResourceQuota 获取命名空间的 ResourceQuota
func (n *Namespace) GetResourceQuota(namespaceName, quotaName string) (*corev1.ResourceQuota, error) {
	quota, err := n.client.GetClientset().CoreV1().ResourceQuotas(namespaceName).Get(n.ctx, quotaName, metav1.GetOptions{})
	if err != nil {
		n.Logger.Errorf("获取命名空间 %s 的 ResourceQuota %s 失败: %v", namespaceName, quotaName, err)
		return nil, err
	}
	n.Logger.Infof("成功获取命名空间 %s 的 ResourceQuota %s", namespaceName, quotaName)
	return quota, nil
}

// DeleteResourceQuota 删除命名空间的 ResourceQuota
func (n *Namespace) DeleteResourceQuota(namespaceName, quotaName string) error {
	err := n.client.GetClientset().CoreV1().ResourceQuotas(namespaceName).Delete(n.ctx, quotaName, metav1.DeleteOptions{})
	if err != nil {
		n.Logger.Errorf("删除命名空间 %s 的 ResourceQuota %s 失败: %v", namespaceName, quotaName, err)
		return err
	}
	n.Logger.Infof("成功删除命名空间 %s 的 ResourceQuota %s", namespaceName, quotaName)
	return nil
}

// UpdateResourceQuota 修改命名空间的 ResourceQuota
func (n *Namespace) UpdateResourceQuota(namespaceName string, quota *corev1.ResourceQuota) error {
	_, err := n.client.GetClientset().CoreV1().ResourceQuotas(namespaceName).Update(n.ctx, quota, metav1.UpdateOptions{})
	if err != nil {
		n.Logger.Errorf("更新命名空间 %s 的 ResourceQuota %s 失败: %v", namespaceName, quota.Name, err)
		return err
	}
	n.Logger.Infof("成功更新命名空间 %s 的 ResourceQuota %s", namespaceName, quota.Name)
	return nil
}

// AddLabel 为命名空间添加标签
func (n *Namespace) AddLabel(namespaceName, key, value string) error {
	return n.updateNamespace(namespaceName, 3, func(ns *corev1.Namespace) error {
		if ns.Labels == nil {
			ns.Labels = make(map[string]string)
		}

		if existingValue, exists := ns.Labels[key]; exists && existingValue == value {
			n.Logger.Infof("命名空间 %s 已存在相同的标签: %s=%s，无需更新", namespaceName, key, value)
			return nil
		}

		ns.Labels[key] = value
		return nil
	})
}

// RemoveLabel 删除命名空间的指定标签
func (n *Namespace) RemoveLabel(namespaceName, key string) error {
	return n.updateNamespace(namespaceName, 3, func(ns *corev1.Namespace) error {
		if _, exists := ns.Labels[key]; !exists {
			n.Logger.Infof("命名空间 %s 不存在标签: %s，无需删除", namespaceName, key)
			return nil
		}

		delete(ns.Labels, key)
		return nil
	})
}

// AddAnnotation 为命名空间添加注解
func (n *Namespace) AddAnnotation(namespaceName, key, value string) error {
	return n.updateNamespace(namespaceName, 3, func(ns *corev1.Namespace) error {
		if ns.Annotations == nil {
			ns.Annotations = make(map[string]string)
		}

		if existingValue, exists := ns.Annotations[key]; exists && existingValue == value {
			n.Logger.Infof("命名空间 %s 已存在相同的注解: %s=%s，无需更新", namespaceName, key, value)
			return nil
		}

		ns.Annotations[key] = value
		return nil
	})
}

// RemoveAnnotation 删除命名空间的指定注解
func (n *Namespace) RemoveAnnotation(namespaceName, key string) error {
	return n.updateNamespace(namespaceName, 3, func(ns *corev1.Namespace) error {
		if _, exists := ns.Annotations[key]; !exists {
			n.Logger.Infof("命名空间 %s 不存在注解: %s，无需删除", namespaceName, key)
			return nil
		}

		delete(ns.Annotations, key)
		return nil
	})
}

// GetNamespaceWithRetry 获取命名空间并带有重试机制
func (n *Namespace) GetNamespaceWithRetry(namespaceName string, maxRetries int) (*corev1.Namespace, error) {
	var ns *corev1.Namespace
	var err error
	for i := 0; i < maxRetries; i++ {
		ns, err = n.client.GetClientset().CoreV1().Namespaces().Get(n.ctx, namespaceName, metav1.GetOptions{})
		if err == nil {
			return ns, nil
		}

		if !errors.IsNotFound(err) && !errors.IsConflict(err) && !isTemporaryError(err) {
			n.Logger.Errorf("获取命名空间 %s 失败: %v", namespaceName, err)
			return nil, err
		}

		n.Logger.Errorf("获取命名空间 %s 失败，正在重试（%d/%d）: %v", namespaceName, i+1, maxRetries, err)
		time.Sleep(100 * time.Millisecond)
	}

	return nil, fmt.Errorf("获取命名空间 %s 失败，已超出最大重试次数", namespaceName)
}

// updateNamespace 封装更新命名空间的逻辑，包含重试机制
func (n *Namespace) updateNamespace(namespaceName string, maxRetries int, updateFunc func(*corev1.Namespace) error) error {
	for i := 0; i < maxRetries; i++ {
		ns, err := n.GetNamespaceWithRetry(namespaceName, 3)
		if err != nil {
			return err
		}

		// 调用自定义更新逻辑
		if err := updateFunc(ns); err != nil {
			n.Logger.Errorf("更新命名空间 %s 失败: %v", namespaceName, err)
			return err
		}

		// 提交更新
		_, err = n.client.GetClientset().CoreV1().Namespaces().Update(n.ctx, ns, metav1.UpdateOptions{})
		if err == nil {
			n.Logger.Infof("成功更新命名空间 %s", namespaceName)
			return nil
		}

		if !errors.IsConflict(err) && !isTemporaryError(err) {
			n.Logger.Errorf("更新命名空间 %s 失败: %v", namespaceName, err)
			return err
		}

		n.Logger.Errorf("更新命名空间 %s 发生冲突或临时错误，正在重试（%d/%d）", namespaceName, i+1, maxRetries)
		time.Sleep(100 * time.Millisecond)
	}

	return fmt.Errorf("更新命名空间 %s 失败，已超出最大重试次数", namespaceName)
}

// isTemporaryError 判断错误是否为临时性错误
func isTemporaryError(err error) bool {
	// 根据需要扩展临时性错误的判断逻辑
	// 这里简单示例，您可以根据实际情况添加更多条件
	return errors.IsServerTimeout(err) || errors.IsTooManyRequests(err)
}

// NamespaceExist 查询名称空间是否存在
func (n *Namespace) NamespaceExist(namespaceName string) bool {
	_, err := n.client.GetClientset().CoreV1().Namespaces().Get(n.ctx, namespaceName, metav1.GetOptions{})
	if err != nil {
		if errors.IsNotFound(err) {
			n.Logger.Infof("命名空间 %s 不存在", namespaceName)
			return false
		}
		n.Logger.Errorf("查询命名空间 %s 是否存在时发生错误: %v", namespaceName, err)
		return true
	}
	n.Logger.Infof("命名空间 %s 存在", namespaceName)
	return true
}
