package node

import (
	"context"
	"github.com/yanshicheng/kube-onec/application/manager/api/internal/svc"
	"github.com/yanshicheng/kube-onec/application/manager/api/internal/types"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/pb"
	"github.com/yanshicheng/kube-onec/pkg/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type EvictOnecNodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEvictOnecNodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EvictOnecNodeLogic {
	return &EvictOnecNodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EvictOnecNodeLogic) EvictOnecNode(req *types.DefaultIdRequest) (resp string, err error) {
	_, err = l.svcCtx.NodeRpc.EvictNodePod(l.ctx, &pb.EvictNodePodReq{
		NodeId:    req.Id,
		UpdatedBy: utils.GetAccount(l.ctx),
	})
	if err != nil {
		l.Logger.Errorf("驱逐节点失败: %v", err)
		return
	}
	return "节点 pod 驱逐成功!", nil
}
