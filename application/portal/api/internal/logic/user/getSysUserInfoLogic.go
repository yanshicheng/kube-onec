package user

import (
	"context"
	"github.com/yanshicheng/kube-onec/application/portal/api/internal/code"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/pb"

	"github.com/yanshicheng/kube-onec/application/portal/api/internal/svc"
	"github.com/yanshicheng/kube-onec/application/portal/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSysUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetSysUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSysUserInfoLogic {
	return &GetSysUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSysUserInfoLogic) GetSysUserInfo() (resp *types.SysUserInfoResponse, err error) {
	account, ok := l.ctx.Value("account").(string)
	if !ok {
		return nil, code.TokenContentErr
	}

	res, err := l.svcCtx.SysUserRpc.GetUserInfo(l.ctx, &pb.GetUserInfoReq{
		Account: account,
	})
	if err != nil {
		return nil, err
	}
	resp = &types.SysUserInfoResponse{
		Id:               res.Id,
		Account:          res.Account,
		UserName:         res.UserName,
		Mobile:           res.Mobile,
		Email:            res.Email,
		HireDate:         res.HireDate,
		Icon:             res.Icon,
		WorkNumber:       res.WorkNumber,
		OrganizationName: res.OrganizationName,
		PositionName:     res.PositionName,
		RoleNames:        res.RoleNames,
		LastLoginTime:    res.LastLoginTime,
		CreatedAt:        res.CreatedAt,
		UpdatedAt:        res.UpdatedAt,
	}
	return
}
