package shared

import (
	"context"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/internal/svc"
	"github.com/yanshicheng/kube-onec/common/k8swrapper/manager"
	"github.com/yanshicheng/kube-onec/utils"
	"github.com/zeromicro/go-zero/core/logx"
)

func GetK8sClient(ctx context.Context, svc *svc.ServiceContext, clusterUuid string) (*manager.OnecK8sClient, error) {
	// 如果已经存在则直接返回 client
	client, err := svc.OnecClient.GetOrCreateOnecK8sClient(ctx, clusterUuid, nil)
	if err != nil {
		// 如果不存在 则查询 cluster 进行创建 并返回 client
		cluster, err := svc.ClusterModel.FindOneByUuid(ctx, clusterUuid)
		config := utils.NewRestConfig(cluster.Host, cluster.Token, utils.IntToBool(cluster.SkipInsecure))
		if err != nil {
			logx.WithContext(ctx).Errorf("获取集群信息失败: %v", err)
			return nil, err
		}
		client, err = svc.OnecClient.GetOrCreateOnecK8sClient(ctx, clusterUuid, config)
		if err != nil {
			// 创建失败返回 error
			logx.WithContext(ctx).Errorf("集群客户端创建失败: %v", err)
			return nil, err
		}
		return client, nil
	}
	return client, nil
}
