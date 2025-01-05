package auth

import (
	"context"
	"github.com/yanshicheng/kube-onec/application/portal/api/internal/code"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/client/sysuserservice"

	"github.com/yanshicheng/kube-onec/application/portal/api/internal/svc"
	"github.com/yanshicheng/kube-onec/application/portal/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginRequest) (resp *types.LoginResponse, err error) {
	switch req.LoginType {
	case 1:
		// 账号密码登录
		login, err := l.svcCtx.SysAuthRpc.GetToken(l.ctx, &sysuserservice.GetTokenRequest{
			Account:  req.Account,
			Password: req.Password,
		})
		if err != nil {
			return nil, err
		}
		var resp types.LoginResponse
		resp.UserName = login.UserName
		resp.AccountId = login.AccountId
		resp.Account = login.Account
		resp.Roles = login.Roles
		resp.Token = types.Token{
			AccessToken:     login.Token.AccessToken,
			RefreshToken:    login.Token.RefreshToken,
			AccessExpireIn:  login.Token.AccessExpiresIn,
			RefreshExpireIn: login.Token.RefreshExpiresIn,
		}
		return &resp, nil
	default:
		return nil, code.LoginTypeNotSupport
		//
		//}
	}
}
