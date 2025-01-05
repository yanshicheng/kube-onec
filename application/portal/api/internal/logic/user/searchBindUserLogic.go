package user

import (
	"context"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/client/sysuserservice"

	"github.com/yanshicheng/kube-onec/application/portal/api/internal/svc"
	"github.com/yanshicheng/kube-onec/application/portal/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchBindUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchBindUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchBindUserLogic {
	return &SearchBindUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchBindUserLogic) SearchBindUser(req *types.SearchBindUserRequest) (resp *types.SearchBindUserResponse, err error) {
	res, err := l.svcCtx.SysUserRpc.GetRoleByUserId(l.ctx, &sysuserservice.GetRoleByUserIdReq{
		Id: req.Id,
	})
	if err != nil {
		l.Logger.Errorf("获取用户信息失败: %v", err)
		return
	}

	return &types.SearchBindUserResponse{
		RoleIds:   res.RoleIds,
		RoleNames: res.RoleNames,
	}, nil
}
