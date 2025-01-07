package cluster

import (
	"context"

	"github.com/yanshicheng/kube-onec/application/manager/api/internal/svc"
	"github.com/yanshicheng/kube-onec/application/manager/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddOnecNodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddOnecNodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddOnecNodeLogic {
	return &AddOnecNodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddOnecNodeLogic) AddOnecNode(req *types.AddOnecNodeRequest) (resp string, err error) {
	// todo: add your logic here and delete this line

	return
}
