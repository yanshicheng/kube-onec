package user

import (
	"context"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/client/sysuserservice"

	"github.com/yanshicheng/kube-onec/application/portal/api/internal/svc"
	"github.com/yanshicheng/kube-onec/application/portal/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GlobalUpdateSysUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGlobalUpdateSysUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GlobalUpdateSysUserLogic {
	return &GlobalUpdateSysUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GlobalUpdateSysUserLogic) GlobalUpdateSysUser(req *types.UpdateGloabSysUserRequest) (resp string, err error) {
	_, err = l.svcCtx.SysUserRpc.UpdateGlobalSysUser(l.ctx, &sysuserservice.UpdateGlobalSysUserReq{
		Id:             req.Id,
		UserName:       req.UserName,
		Email:          req.Email,
		Mobile:         req.Mobile,
		WorkNumber:     req.WorkNumber,
		HireDate:       req.HireDate,
		PositionId:     req.PositionId,
		OrganizationId: req.OrganizationId,
	})
	if err != nil {
		l.Logger.Errorf("获取用户信息失败: %v", err)
		return
	}
	return "更新成功", nil
}
