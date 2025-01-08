package shared

import (
	"context"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/internal/model"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

func ChangeNodeSyncStatus(ctx context.Context, svcCtx *svc.ServiceContext, node model.OnecNode, syncStatus int64, updatedBy string) {
	node.SyncStatus = syncStatus
	node.UpdatedBy = updatedBy
	if err := svcCtx.NodeModel.Update(ctx, &node); err != nil {
		logx.WithContext(ctx).Errorf("更新节点同步状态失败: %v, 集群: %v nodes: %v", err, node.ClusterUuid, node.NodeName)
	}
}
