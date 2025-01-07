package user

import (
	"context"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/client/sysuserservice"

	"github.com/yanshicheng/kube-onec/application/portal/api/internal/svc"
	"github.com/yanshicheng/kube-onec/application/portal/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSysUserByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetSysUserByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSysUserByIdLogic {
	return &GetSysUserByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSysUserByIdLogic) GetSysUserById(req *types.DefaultIdRequest) (resp *types.SysUser, err error) {
	res, err := l.svcCtx.SysUserRpc.GetSysUserById(l.ctx, &sysuserservice.GetSysUserByIdReq{
		Id: req.Id,
	})

	if err != nil {
		l.Logger.Errorf("获取用户信息失败: %v", err)
		return
	}
	resp = &types.SysUser{
		Id:              res.Data.Id,
		UserName:        res.Data.UserName,
		Account:         res.Data.Account,
		Icon:            res.Data.Icon,
		Mobile:          res.Data.Mobile,
		Email:           res.Data.Email,
		WorkNumber:      res.Data.WorkNumber,
		HireDate:        res.Data.HireDate,
		IsResetPassword: res.Data.IsResetPassword,
		IsDisabled:      res.Data.IsDisabled,
		IsLeave:         res.Data.IsLeave,
		PositionId:      res.Data.PositionId,
		OrganizationId:  res.Data.OrganizationId,
		LastLoginTime:   res.Data.LastLoginTime,
		CreatedAt:       res.Data.CreatedAt,
		UpdatedAt:       res.Data.UpdatedAt,
	}
	return resp, nil
}
