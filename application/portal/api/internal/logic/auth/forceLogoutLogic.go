package auth

import (
	"context"

	"github.com/yanshicheng/kube-onec/application/portal/api/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type ForceLogoutLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewForceLogoutLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ForceLogoutLogic {
	return &ForceLogoutLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ForceLogoutLogic) ForceLogout() (resp string, err error) {
	// todo: add your logic here and delete this line

	return
}
