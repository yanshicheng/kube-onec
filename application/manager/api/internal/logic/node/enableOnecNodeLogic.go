package node

import (
	"context"

	"github.com/yanshicheng/kube-onec/application/manager/api/internal/svc"
	"github.com/yanshicheng/kube-onec/application/manager/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type EnableOnecNodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEnableOnecNodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EnableOnecNodeLogic {
	return &EnableOnecNodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EnableOnecNodeLogic) EnableOnecNode(req *types.DefaultIdRequest) (resp string, err error) {
	// todo: add your logic here and delete this line

	return
}
