package onec_project

import (
	"context"
	"github.com/yanshicheng/kube-onec/application/manager/api/internal/svc"
	"github.com/yanshicheng/kube-onec/application/manager/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOnecProjectQuotaByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetOnecProjectQuotaByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOnecProjectQuotaByIdLogic {
	return &GetOnecProjectQuotaByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetOnecProjectQuotaByIdLogic) GetOnecProjectQuotaById(req *types.DefaultIdRequest) (resp string, err error) {

	return
}
