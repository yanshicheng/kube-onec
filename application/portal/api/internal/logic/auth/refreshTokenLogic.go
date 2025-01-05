package auth

import (
	"context"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/client/sysauthservice"

	"github.com/yanshicheng/kube-onec/application/portal/api/internal/svc"
	"github.com/yanshicheng/kube-onec/application/portal/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RefreshTokenLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRefreshTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RefreshTokenLogic {
	return &RefreshTokenLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RefreshTokenLogic) RefreshToken(req *types.RefreshTokenRequest) (resp *types.RefreshTokenResponse, err error) {
	res, err := l.svcCtx.SysAuthRpc.RefreshToken(l.ctx, &sysauthservice.RefreshTokenRequest{
		RefreshToken: req.RefreshToken,
	})
	if err != nil {
		l.Logger.Errorf("RefreshToken err: %v", err)
		return nil, err
	}
	return &types.RefreshTokenResponse{
		AccessToken:    res.AccessToken,
		AccessExpireIn: res.AccessExpiresIn,
	}, nil
}
