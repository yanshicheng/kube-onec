package auth

import (
	"context"
	"github.com/yanshicheng/kube-onec/application/portal/api/internal/code"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/pb"

	"github.com/yanshicheng/kube-onec/application/portal/api/internal/svc"
	"github.com/yanshicheng/kube-onec/application/portal/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LogoutLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLogoutLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LogoutLogic {
	return &LogoutLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LogoutLogic) Logout() (resp *types.LogoutResponse, err error) {
	account, ok := l.ctx.Value("account").(string)
	if !ok {
		return nil, code.UUIDNotExistErr
	}
	uuid, ok := l.ctx.Value("uuid").(string)
	if !ok {
		return nil, code.UUIDNotExistErr
	}
	_, err = l.svcCtx.SysAuthRpc.Logout(l.ctx, &pb.LogoutRequest{
		Account: account,
		Uuid:    uuid,
	})
	if err != nil {
		return nil, err
	}
	return &types.LogoutResponse{
		Message: "退出系统正常!",
		Success: true,
	}, nil
}
