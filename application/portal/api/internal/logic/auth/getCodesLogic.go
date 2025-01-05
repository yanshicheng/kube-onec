package auth

import (
	"context"

	"github.com/yanshicheng/kube-onec/application/portal/api/internal/svc"
	"github.com/yanshicheng/kube-onec/application/portal/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCodesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCodesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCodesLogic {
	return &GetCodesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCodesLogic) GetCodes() (resp *types.GetCodesResponse, err error) {
	return
}
