package onecnodeservicelogic

import (
	"context"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/internal/code"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/internal/shared"

	"github.com/yanshicheng/kube-onec/application/manager/rpc/internal/svc"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type EvictNodePodLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewEvictNodePodLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EvictNodePodLogic {
	return &EvictNodePodLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 驱逐节点pod
func (l *EvictNodePodLogic) EvictNodePod(in *pb.EvictNodePodReq) (*pb.EvictNodePodResp, error) {
	node, err := l.svcCtx.NodeModel.FindOne(l.ctx, in.NodeId)
	if err != nil {
		l.Logger.Errorf("获取节点信息失败: %v, nodes: %v", err, in.NodeId)
		return nil, code.GetNodeInfoErr
	}

	// 获取 Kubernetes 客户端
	client, err := shared.GetK8sClient(l.ctx, l.svcCtx, node.ClusterUuid)
	if err != nil {
		l.Logger.Errorf("获取集群客户端异常: %v", err)
		return nil, code.GetClusterClientErr
	}

	err = client.GetNodeClient().ForceEvictAllPods(node.NodeName)
	if err != nil {
		l.Logger.Errorf("驱逐节点pod失败: %v", err)
		return nil, code.EvictNodePodErr
	}
	shared.ChangeNodeSyncStatus(l.ctx, l.svcCtx, *node, 1, in.UpdatedBy)
	return &pb.EvictNodePodResp{}, nil
}
