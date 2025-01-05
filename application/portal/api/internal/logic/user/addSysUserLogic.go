package user

import (
	"context"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/pb"

	"github.com/yanshicheng/kube-onec/application/portal/api/internal/svc"
	"github.com/yanshicheng/kube-onec/application/portal/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddSysUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddSysUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddSysUserLogic {
	return &AddSysUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddSysUserLogic) AddSysUser(req *types.AddSysUserRequest) (resp string, err error) {
	_, err = l.svcCtx.SysUserRpc.AddSysUser(l.ctx, &pb.AddSysUserReq{
		UserName:       req.UserName,
		Account:        req.Account,
		Mobile:         req.Mobile,
		Email:          req.Email,
		WorkNumber:     req.WorkNumber,
		HireDate:       req.HireDate,
		PositionId:     req.PositionId,
		OrganizationId: req.OrganizationId,
	})
	if err != nil {
		l.Logger.Errorf("添加用户失败: %v", err)
		return "", err
	}
	l.Logger.Infof("添加用户成功: %v", req)
	return "添加用户成功", nil
}
