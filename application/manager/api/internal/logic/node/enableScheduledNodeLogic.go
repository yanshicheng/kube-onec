package node

import (
	"context"
	"github.com/yanshicheng/kube-onec/application/manager/api/internal/svc"
	"github.com/yanshicheng/kube-onec/application/manager/api/internal/types"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/client/onecnodeservice"
	"github.com/yanshicheng/kube-onec/pkg/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type EnableScheduledNodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEnableScheduledNodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EnableScheduledNodeLogic {
	return &EnableScheduledNodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EnableScheduledNodeLogic) EnableScheduledNode(req *types.EnableScheduledRequest) (resp string, err error) {
	_, err = l.svcCtx.NodeRpc.EnableScheduledNode(l.ctx, &onecnodeservice.EnableScheduledNodeReq{
		NodeId:    req.Id,
		UpdatedBy: utils.GetAccount(l.ctx),
	})
	if err != nil {
		l.Logger.Errorf("启用调度失败: %v", err)
		return
	}
	return "启用调度成功!", nil

}
