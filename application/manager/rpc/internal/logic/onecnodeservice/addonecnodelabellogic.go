package onecnodeservicelogic

import (
	"context"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/internal/code"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/internal/model"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/internal/shared"

	"github.com/yanshicheng/kube-onec/application/manager/rpc/internal/svc"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddOnecNodeLabelLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddOnecNodeLabelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddOnecNodeLabelLogic {
	return &AddOnecNodeLabelLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 节点添加标签
func (l *AddOnecNodeLabelLogic) AddOnecNodeLabel(in *pb.AddOnecNodeLabelReq) (*pb.AddOnecNodeLabelResp, error) {

	node, err := l.svcCtx.NodeModel.FindOne(l.ctx, in.NodeId)
	if err != nil {
		l.Logger.Errorf("获取节点信息失败: %v", err)
		return nil, code.GetNodeInfoErr
	}
	client, err := shared.GetK8sClient(l.ctx, l.svcCtx, node.ClusterUuid)
	if err != nil {
		l.Logger.Errorf("获取集群客户端异常: %v", err)
		return nil, code.GetClusterClientErr
	}

	err = client.GetNodeClient().AddLabel(node.NodeName, in.Key, in.Value)
	if err != nil {
		l.Logger.Errorf("节点标签数据同步失败: %v", err)
		shared.ChangeNodeSyncStatus(l.ctx, l.svcCtx, *node, 0, in.UpdatedBy)
		return nil, code.SyncNodeLabelErr
	}

	// 同步节点标签到数据库
	_, err = l.svcCtx.LabelsResourceModel.Insert(l.ctx, &model.OnecResourceLabels{
		Key:          in.Key,
		Value:        in.Value,
		ResourceType: "node",
		ResourceId:   node.Id,
	})
	if err != nil {
		l.Logger.Errorf("节点标签数据同步失败: %v", err)
		shared.ChangeNodeSyncStatus(l.ctx, l.svcCtx, *node, 0, in.UpdatedBy)
		return nil, code.SyncNodeLabelErr
	}
	shared.ChangeNodeSyncStatus(l.ctx, l.svcCtx, *node, 1, in.UpdatedBy)
	l.Logger.Infof("节点标签数据同步成功: %v", node.NodeName)
	return &pb.AddOnecNodeLabelResp{}, nil
}
