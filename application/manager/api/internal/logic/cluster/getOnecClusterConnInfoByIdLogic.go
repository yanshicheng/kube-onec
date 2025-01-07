package cluster

import (
	"context"

	"github.com/yanshicheng/kube-onec/application/manager/api/internal/svc"
	"github.com/yanshicheng/kube-onec/application/manager/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOnecClusterConnInfoByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetOnecClusterConnInfoByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOnecClusterConnInfoByIdLogic {
	return &GetOnecClusterConnInfoByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetOnecClusterConnInfoByIdLogic) GetOnecClusterConnInfoById(req *types.DefaultIdRequest) (resp *types.OnecClusterConnInfo, err error) {
	// todo: add your logic here and delete this line

	return
}
