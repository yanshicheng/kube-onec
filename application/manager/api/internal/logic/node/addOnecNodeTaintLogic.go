package node

import (
	"context"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/client/onecnodeservice"
	"github.com/yanshicheng/kube-onec/utils"

	"github.com/yanshicheng/kube-onec/application/manager/api/internal/svc"
	"github.com/yanshicheng/kube-onec/application/manager/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddOnecNodeTaintLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddOnecNodeTaintLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddOnecNodeTaintLogic {
	return &AddOnecNodeTaintLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddOnecNodeTaintLogic) AddOnecNodeTaint(req *types.AddOnecNodeTaintRequest) (resp string, err error) {

	_, err = l.svcCtx.NodeRpc.AddOnecNodeTaint(l.ctx, &onecnodeservice.AddOnecNodeTaintReq{
		UpdatedBy: utils.GetAccount(l.ctx),
		NodeId:    req.Id,
		Effect:    req.Effect,
		Key:       req.Key,
		Value:     req.Value,
	})
	if err != nil {
		l.Logger.Errorf("添加节点污点失败: %v", err)
		return
	}

	return "添加节点污点成功!", nil
}
