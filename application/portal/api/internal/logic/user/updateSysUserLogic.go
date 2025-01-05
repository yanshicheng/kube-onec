package user

import (
	"context"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/client/sysuserservice"

	"github.com/yanshicheng/kube-onec/application/portal/api/internal/svc"
	"github.com/yanshicheng/kube-onec/application/portal/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateSysUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateSysUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateSysUserLogic {
	return &UpdateSysUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateSysUserLogic) UpdateSysUser(req *types.UpdateSysUserRequest) (resp string, err error) {
	_, err = l.svcCtx.SysUserRpc.UpdateSysUser(l.ctx, &sysuserservice.UpdateSysUserReq{
		Id:       req.Id,
		UserName: req.UserName,
		Email:    req.Email,
		Mobile:   req.Mobile,
	})
	if err != nil {
		l.Logger.Errorf("获取用户信息失败: %v", err)
		return
	}
	return "更新成功", nil
}
