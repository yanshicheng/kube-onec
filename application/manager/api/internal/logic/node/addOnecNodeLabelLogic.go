package node

import (
	"context"
	"github.com/yanshicheng/kube-onec/application/manager/api/internal/svc"
	"github.com/yanshicheng/kube-onec/application/manager/api/internal/types"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/client/onecnodeservice"
	"github.com/yanshicheng/kube-onec/pkg/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddOnecNodeLabelLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddOnecNodeLabelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddOnecNodeLabelLogic {
	return &AddOnecNodeLabelLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddOnecNodeLabelLogic) AddOnecNodeLabel(req *types.AddOnecNodeLabelRequest) (resp string, err error) {

	_, err = l.svcCtx.NodeRpc.AddOnecNodeLabel(l.ctx, &onecnodeservice.AddOnecNodeLabelReq{
		NodeId:    req.Id,
		Key:       req.Key,
		Value:     req.Value,
		UpdatedBy: utils.GetAccount(l.ctx),
	})
	if err != nil {
		l.Logger.Errorf("添加节点标签失败: %v", err)
		return
	}
	l.Logger.Infof("添加节点标签成功: %v", req)
	return "添加节点标签成功!", nil
}
