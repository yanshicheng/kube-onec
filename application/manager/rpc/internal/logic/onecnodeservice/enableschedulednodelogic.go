package onecnodeservicelogic

import (
	"context"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/internal/code"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/internal/shared"

	"github.com/yanshicheng/kube-onec/application/manager/rpc/internal/svc"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type EnableScheduledNodeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewEnableScheduledNodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EnableScheduledNodeLogic {
	return &EnableScheduledNodeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 取消禁止调度
func (l *EnableScheduledNodeLogic) EnableScheduledNode(in *pb.EnableScheduledNodeReq) (*pb.EnableScheduledNodeResp, error) {
	// 获取节点信息
	node, err := l.svcCtx.NodeModel.FindOne(l.ctx, in.NodeId)
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

	// 设置节点为可调度
	err = client.GetNodeClient().EnableScheduled(node.NodeName)
	if err != nil {
		l.Logger.Errorf("设置节点可调度失败: %v", err)
		shared.ChangeNodeSyncStatus(l.ctx, l.svcCtx, *node, 0, in.UpdatedBy) // 更新节点同步状态为失败
		return nil, code.EnableNodeScheduleErr
	}

	node.SyncStatus = 1
	node.Unschedulable = 0
	node.UpdatedBy = in.UpdatedBy
	if err := l.svcCtx.NodeModel.Update(l.ctx, node); err != nil {
		l.Logger.Errorf("更新节点信息失败: %v", err)
		return nil, code.UpdateNodeInfoErr
	}
	// 更新节点同步状态为成功
	l.Logger.Infof("节点取消禁止调度成功: %v", node.NodeName)

	return &pb.EnableScheduledNodeResp{}, nil
}
