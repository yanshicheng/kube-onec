package node

import (
	"context"
	"github.com/yanshicheng/kube-onec/application/manager/api/internal/svc"
	"github.com/yanshicheng/kube-onec/application/manager/api/internal/types"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/client/onecnodeservice"
	"github.com/yanshicheng/kube-onec/pkg/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelOnecNodeTaintLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDelOnecNodeTaintLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelOnecNodeTaintLogic {
	return &DelOnecNodeTaintLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DelOnecNodeTaintLogic) DelOnecNodeTaint(req *types.DelOnecNodeTaintRequest) (resp string, err error) {
	_, err = l.svcCtx.NodeRpc.DelOnecNodeTaint(l.ctx, &onecnodeservice.DelOnecNodeTaintReq{
		TaintId:   req.Id,
		UpdatedBy: utils.GetAccount(l.ctx),
	})
	if err != nil {
		l.Logger.Errorf("删除节点污点失败: %v", err)
		return
	}

	return "删除节点污点成功!", nil
}
