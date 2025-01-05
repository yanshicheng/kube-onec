package cluster

import (
	"context"
	"github.com/yanshicheng/kube-onec/common/k8swrapper/core"
	"github.com/yanshicheng/kube-onec/utils"
	"github.com/zeromicro/go-zero/core/logx"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"log"
	"time"
)

// Cluster 实现 core.ClusterInterface
type Cluster struct {
	client core.K8sClientInterface
	ctx    context.Context
	Logger logx.Logger
}

// NewCluster 初始化 Cluster
func NewCluster(ctx context.Context, client core.K8sClientInterface) *Cluster {
	return &Cluster{client: client, ctx: ctx, Logger: logx.WithContext(ctx)}
}

func (c *Cluster) GetClusterInfo() (*core.ClusterInfo, error) {
	version, err := c.client.GetClientset().Discovery().ServerVersion()
	if err != nil {
		c.Logger.Infof("获取集群信息失败, 错误: %v", err)
		return nil, err
	}
	buildTime, err := utils.ParseStringToTime(version.BuildDate, time.RFC3339)
	if err != nil {
		c.Logger.Infof("解析集群版本信息失败, 错误: %v", err)
		return nil, err
	}
	createTime := time.Now()
	tim, err := c.GetClusterCreateTime()
	if err == nil {
		createTime = tim
	}
	c.Logger.Infof("获取集群信息成功: %v", version.GoVersion)
	return &core.ClusterInfo{
		Version:    version.GitVersion,
		Commit:     version.GitCommit,
		Platform:   version.Platform,
		BuildTime:  buildTime,
		CreateTime: createTime,
	}, nil
}

// 获取集群创建的时间
func (c *Cluster) GetClusterCreateTime() (time.Time, error) {
	// 获取 kube-system Namespace 下的 kubernetes Service
	service, err := c.client.GetClientset().CoreV1().Services("default").Get(c.ctx, "kubernetes", metav1.GetOptions{})
	if err != nil {
		log.Fatalf("无法获取 kubernetes Service: %v", err)
	}
	return service.ObjectMeta.CreationTimestamp.Time, nil
}

func (c *Cluster) RemoveNode(nodeName string) error {
	// 调用 Kubernetes API 删除节点
	err := c.client.GetClientset().CoreV1().Nodes().Delete(c.ctx, nodeName, metav1.DeleteOptions{})
	if err != nil {
		c.Logger.Errorf("删除节点 %s 失败: %v", nodeName, err)
		return err
	}

	c.Logger.Infof("成功删除节点 %s", nodeName)
	return nil
}
