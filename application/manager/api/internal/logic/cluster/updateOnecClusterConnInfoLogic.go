package cluster

import (
	"context"
	"github.com/yanshicheng/kube-onec/application/manager/api/internal/svc"
	"github.com/yanshicheng/kube-onec/application/manager/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateOnecClusterConnInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateOnecClusterConnInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateOnecClusterConnInfoLogic {
	return &UpdateOnecClusterConnInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateOnecClusterConnInfoLogic) UpdateOnecClusterConnInfo(req *types.UpdateOnecClusterConnInfoRequest) (resp string, err error) {
	return
}
