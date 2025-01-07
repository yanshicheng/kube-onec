package user

import (
	"context"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/pb"

	"github.com/yanshicheng/kube-onec/application/portal/api/internal/svc"
	"github.com/yanshicheng/kube-onec/application/portal/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchSysUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchSysUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchSysUserLogic {
	return &SearchSysUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchSysUserLogic) SearchSysUser(req *types.SearchSysUserRequest) (resp *types.SearchSysUserResponse, err error) {
	res, err := l.svcCtx.SysUserRpc.SearchSysUser(l.ctx, &pb.SearchSysUserReq{
		Page:               req.Page,
		PageSize:           req.PageSize,
		OrderStr:           req.OrderStr,
		IsAsc:              req.IsAsc,
		UserName:           req.UserName,
		Account:            req.Account,
		Mobile:             req.Mobile,
		Email:              req.Email,
		WorkNumber:         req.WorkNumber,
		HireDate:           req.HireDate,
		IsDisabled:         req.IsDisabled,
		IsLeave:            req.IsLeave,
		PositionId:         req.PositionId,
		OrganizationId:     req.OrganizationId,
		StartLastLoginTime: req.StartLastLoginTime,
		EndLastLoginTime:   req.EndLastLoginTime,
	})
	if err != nil {
		l.Logger.Errorf("获取用户信息失败: %v", err)
		return
	}
	items := make([]types.SysUser, len(res.Data))
	for i, v := range res.Data {
		items[i] = types.SysUser{
			Id:              v.Id,
			UserName:        v.UserName,
			Account:         v.Account,
			Icon:            v.Icon,
			Mobile:          v.Mobile,
			Email:           v.Email,
			WorkNumber:      v.WorkNumber,
			HireDate:        v.HireDate,
			IsResetPassword: v.IsResetPassword,
			IsDisabled:      v.IsDisabled,
			IsLeave:         v.IsLeave,
			PositionId:      v.PositionId,
			OrganizationId:  v.OrganizationId,
			LastLoginTime:   v.LastLoginTime,
			CreatedAt:       v.CreatedAt,
			UpdatedAt:       v.UpdatedAt,
		}
	}
	return &types.SearchSysUserResponse{
		Items: items,
		Total: res.Total,
	}, nil
}
