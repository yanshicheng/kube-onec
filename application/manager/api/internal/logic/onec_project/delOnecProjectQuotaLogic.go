package onec_project

import (
	"context"

	"github.com/yanshicheng/kube-onec/application/manager/api/internal/svc"
	"github.com/yanshicheng/kube-onec/application/manager/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelOnecProjectQuotaLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDelOnecProjectQuotaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelOnecProjectQuotaLogic {
	return &DelOnecProjectQuotaLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DelOnecProjectQuotaLogic) DelOnecProjectQuota(req *types.DefaultIdRequest) (resp string, err error) {
	// todo: add your logic here and delete this line

	return
}
