package onecnodeservicelogic

import (
	"context"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/internal/code"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/internal/shared"

	"github.com/yanshicheng/kube-onec/application/manager/rpc/internal/svc"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ForbidScheduledLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewForbidScheduledLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ForbidScheduledLogic {
	return &ForbidScheduledLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 禁止调度
func (l *ForbidScheduledLogic) ForbidScheduled(in *pb.ForbidScheduledReq) (*pb.ForbidScheduledResp, error) {

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

	// 设置节点为不可调度
	err = client.GetNodeClient().ForbidScheduled(node.NodeName)
	if err != nil {
		l.Logger.Errorf("设置节点不可调度失败: %v", err)
		shared.ChangeNodeSyncStatus(l.ctx, l.svcCtx, *node, 0, in.UpdatedBy) // 更新节点同步状态为失败
		return nil, code.ForbidNodeScheduleErr
	}

	// 更新节点同步状态为成功
	node.SyncStatus = 1
	node.Unschedulable = 1
	node.UpdatedBy = in.UpdatedBy
	if err := l.svcCtx.NodeModel.Update(l.ctx, node); err != nil {
		l.Logger.Errorf("更新节点信息失败: %v", err)
		return nil, code.UpdateNodeInfoErr
	}
	l.Logger.Infof("节点禁止调度成功: %v", node.NodeName)

	return &pb.ForbidScheduledResp{}, nil
}
