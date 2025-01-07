package node

import (
	"context"

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
	// todo: add your logic here and delete this line

	return
}
