package cluster

import (
	"context"

	"github.com/yanshicheng/kube-onec/application/manager/api/internal/svc"
	"github.com/yanshicheng/kube-onec/application/manager/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelOnecClusterConnInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDelOnecClusterConnInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelOnecClusterConnInfoLogic {
	return &DelOnecClusterConnInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DelOnecClusterConnInfoLogic) DelOnecClusterConnInfo(req *types.DefaultIdRequest) (resp string, err error) {
	// todo: add your logic here and delete this line

	return
}
