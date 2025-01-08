package onecnodeservicelogic

import (
	"context"
	"errors"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/internal/code"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/internal/model"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/internal/shared"

	"github.com/yanshicheng/kube-onec/application/manager/rpc/internal/svc"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelOnecNodeLabelLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelOnecNodeLabelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelOnecNodeLabelLogic {
	return &DelOnecNodeLabelLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 节点删除标签
// 节点删除标签
func (l *DelOnecNodeLabelLogic) DelOnecNodeLabel(in *pb.DelOnecNodeLabelReq) (*pb.DelOnecNodeLabelResp, error) {
	// 从数据库删除标签记录
	label, err := l.svcCtx.LabelsResourceModel.FindOne(l.ctx, in.LabelId)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			l.Logger.Infof("污点记录不存在: %v", err)
			return nil, code.NodeInfoNotExistErr
		}
		l.Logger.Errorf("查询节点标签数据库记录失败: %v", err)
		return nil, code.SearchNodeLabelDBErr
	}
	// 获取节点信息
	node, err := l.svcCtx.NodeModel.FindOne(l.ctx, label.ResourceId)
	if err != nil {
		l.Logger.Errorf("获取节点信息失败: %v", err)
		return nil, code.GetNodeInfoErr
	}
	// 获取 Kubernetes 客户端
	client, err := shared.GetK8sClient(l.ctx, l.svcCtx, node.ClusterUuid)
	if err != nil {
		l.Logger.Errorf("获取集群客户端异常: %v", err)
		return nil, code.GetClusterClientErr
	}

	// 删除节点标签
	err = client.GetNodeClient().RemoveLabel(node.NodeName, label.Key)
	if err != nil {
		l.Logger.Errorf("节点标签删除失败: %v", err)
		shared.ChangeNodeSyncStatus(l.ctx, l.svcCtx, *node, 0, in.UpdatedBy) // 更新节点同步状态为失败
		return nil, code.RemoveNodeLabelErr
	}

	if err := l.svcCtx.LabelsResourceModel.Delete(l.ctx, label.Id); err != nil {
		l.Logger.Errorf("删除节点标签数据库记录失败: %v", err)
		shared.ChangeNodeSyncStatus(l.ctx, l.svcCtx, *node, 0, in.UpdatedBy) // 更新节点同步状态为失败
		return nil, code.DeleteNodeLabelDBErr
	}
	// 更新同步状态为成功
	shared.ChangeNodeSyncStatus(l.ctx, l.svcCtx, *node, 1, in.UpdatedBy)
	l.Logger.Infof("节点标签删除成功: %v", node.NodeName)
	return &pb.DelOnecNodeLabelResp{}, nil
}
