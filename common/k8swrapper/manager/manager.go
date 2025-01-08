package manager

import (
	"context"
	"fmt"
	"github.com/yanshicheng/kube-onec/common/k8swrapper/namespace"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sync"

	"github.com/yanshicheng/kube-onec/common/k8swrapper/cluster"
	"github.com/yanshicheng/kube-onec/common/k8swrapper/core"
	"github.com/yanshicheng/kube-onec/common/k8swrapper/nodes"
	"github.com/zeromicro/go-zero/core/logx"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

// OnecK8sClient 实现 core.K8sClientInterface
type OnecK8sClient struct {
	clientset *kubernetes.Clientset
	cluster   core.ClusterInterface
	nodes     core.NodesInterface
	namespace core.NamespacesInterface
	Logger    logx.Logger
	ctx       context.Context
}

// NewOnecK8sClient 创建一个新的 OnecK8sClient 实例
func NewOnecK8sClient(ctx context.Context, config *rest.Config) (*OnecK8sClient, error) {
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, fmt.Errorf("创建 Kubernetes 客户端失败: %v", err)
	}

	if ctx == nil {
		ctx = context.Background()
	}

	return &OnecK8sClient{
		clientset: clientset,
		ctx:       ctx,
		Logger:    logx.WithContext(ctx),
	}, nil
}

// SetContext 更新上下文
func (k *OnecK8sClient) SetContext(ctx context.Context) {
	k.ctx = ctx
	k.Logger = logx.WithContext(ctx)
}

// GetContext 获取上下文
func (k *OnecK8sClient) GetContext() context.Context {
	return k.ctx
}

// GetCluster 返回集群模块
func (k *OnecK8sClient) GetClusterClient() core.ClusterInterface {
	return cluster.NewCluster(k.ctx, k)
}

// GetNodes 返回节点模块
func (k *OnecK8sClient) GetNodeClient() core.NodesInterface {
	k.nodes = nodes.NewNodes(k.ctx, k)
	return k.nodes
}

// GetNamespace 返回命名空间模块
func (k *OnecK8sClient) GetNamespaceClient() core.NamespacesInterface {
	k.namespace = namespace.NewNamespace(k.ctx, k)
	return k.namespace
}

// GetClientset 返回 Kubernetes Clientset
func (k *OnecK8sClient) GetClientset() *kubernetes.Clientset {
	return k.clientset
}

// Ping 测试连接是否正常
func (k *OnecK8sClient) Ping() error {
	_, err := k.clientset.CoreV1().Namespaces().List(k.ctx, metav1.ListOptions{})
	return err
}

// OnecK8sClientManager 管理多个集群的客户端连接
type OnecK8sClientManager struct {
	clients sync.Map // 使用 sync.Map 缓存多个 K8sClient
	locks   sync.Map // 使用 sync.Map 为每个键创建一个独立的锁
}

// NewK8sClientManager 创建 OnecK8sClientManager 实例
func NewOnecK8sClientManager() *OnecK8sClientManager {
	return &OnecK8sClientManager{}
}

// getLock 获取特定键的锁
func (manager *OnecK8sClientManager) getLock(key string) *sync.Mutex {
	actual, _ := manager.locks.LoadOrStore(key, &sync.Mutex{})
	return actual.(*sync.Mutex)
}

// getClientWithContext 安全获取客户端并替换上下文
func (manager *OnecK8sClientManager) getClientWithContext(ctx context.Context, key string) (*OnecK8sClient, bool) {
	if client, ok := manager.clients.Load(key); ok {
		cachedClient := client.(*OnecK8sClient)
		cachedClient.SetContext(ctx)
		return cachedClient, true
	}
	return nil, false
}

// GetOrCreateOnecK8sClient 获取或创建集群客户端
func (manager *OnecK8sClientManager) GetOrCreateOnecK8sClient(ctx context.Context, key string, config *rest.Config) (*OnecK8sClient, error) {
	// 尝试从缓存中获取客户端
	if client, ok := manager.getClientWithContext(ctx, key); ok {
		return client, nil
	}

	// 获取锁，确保同一键的客户端只会被创建一次
	lock := manager.getLock(key)
	lock.Lock()
	defer lock.Unlock()

	// 再次检查缓存
	if client, ok := manager.getClientWithContext(ctx, key); ok {
		return client, nil
	}

	// 创建新的客户端
	if config != nil {
		client, err := NewOnecK8sClient(ctx, config)
		if err != nil {
			return nil, err
		}

		// 缓存客户端
		manager.clients.Store(key, client)
		return client, nil
	}
	return nil, fmt.Errorf("无法创建客户端，请检查配置")
}

// GetOrCreateK8sClientFromKubeConfig 根据 kubeconfig 文件获取或创建客户端
func (manager *OnecK8sClientManager) GetOrCreateK8sClientFromKubeConfig(ctx context.Context, key, kubeconfigPath string) (*OnecK8sClient, error) {
	// 尝试从缓存中获取客户端
	if client, ok := manager.getClientWithContext(ctx, key); ok {
		return client, nil
	}

	// 获取锁，确保同一键的客户端只会被创建一次
	lock := manager.getLock(key)
	lock.Lock()
	defer lock.Unlock()

	// 再次检查缓存
	if client, ok := manager.getClientWithContext(ctx, key); ok {
		return client, nil
	}

	// 从 kubeconfig 文件创建配置
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfigPath)
	if err != nil {
		return nil, fmt.Errorf("加载 kubeconfig 文件失败: %v", err)
	}

	// 创建新的客户端
	client, err := NewOnecK8sClient(ctx, config)
	if err != nil {
		return nil, err
	}

	// 缓存客户端
	manager.clients.Store(key, client)

	return client, nil
}

// AddK8sClient 手动添加一个新的客户端到连接池
func (manager *OnecK8sClientManager) AddK8sClient(key string, client *OnecK8sClient) bool {
	lock := manager.getLock(key)
	lock.Lock()
	defer lock.Unlock()

	// 检查是否已经存在
	if _, ok := manager.clients.Load(key); ok {
		fmt.Printf("客户端键 %s 已存在，无法添加\n", key)
		return false
	}

	// 插入新的客户端
	manager.clients.Store(key, client)
	fmt.Printf("成功添加客户端键 %s\n", key)
	return true
}

// RemoveK8sClient 从连接池中移除指定的客户端
func (manager *OnecK8sClientManager) RemoveK8sClient(key string) bool {
	lock := manager.getLock(key)
	lock.Lock()
	defer lock.Unlock()

	// 检查是否存在
	if _, ok := manager.clients.Load(key); !ok {
		fmt.Printf("客户端键 %s 不存在，无法移除\n", key)
		return false
	}

	// 移除客户端
	manager.clients.Delete(key)
	manager.locks.Delete(key)
	fmt.Printf("成功移除客户端键 %s\n", key)
	return true
}

// ListK8sClients 列出所有缓存的客户端键
func (manager *OnecK8sClientManager) ListK8sClients() []string {
	keys := []string{}
	manager.clients.Range(func(key, value interface{}) bool {
		keys = append(keys, key.(string))
		return true
	})
	fmt.Println("缓存的客户端键:", keys)
	return keys
}
