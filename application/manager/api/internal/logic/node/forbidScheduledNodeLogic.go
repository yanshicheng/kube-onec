package node

import (
	"context"
	"github.com/yanshicheng/kube-onec/application/manager/api/internal/svc"
	"github.com/yanshicheng/kube-onec/application/manager/api/internal/types"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/client/onecnodeservice"
	"github.com/yanshicheng/kube-onec/pkg/utils"
	"github.com/zeromicro/go-zero/core/logx"
)

type ForbidScheduledNodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewForbidScheduledNodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ForbidScheduledNodeLogic {
	return &ForbidScheduledNodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ForbidScheduledNodeLogic) ForbidScheduledNode(req *types.ForbidScheduledRequest) (resp string, err error) {
	_, err = l.svcCtx.NodeRpc.ForbidScheduled(l.ctx, &onecnodeservice.ForbidScheduledReq{
		NodeId:    req.Id,
		UpdatedBy: utils.GetAccount(l.ctx),
	})
	if err != nil {
		l.Logger.Errorf("禁用调度失败: %v", err)
		return
	}
	return "禁用调度成功!", nil
}
