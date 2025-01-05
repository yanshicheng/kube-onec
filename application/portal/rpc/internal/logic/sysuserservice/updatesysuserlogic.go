package sysuserservicelogic

import (
	"context"
	"errors"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/internal/code"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/internal/model"

	"github.com/yanshicheng/kube-onec/application/portal/rpc/internal/svc"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateSysUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateSysUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateSysUserLogic {
	return &UpdateSysUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateSysUserLogic) UpdateSysUser(in *pb.UpdateSysUserReq) (*pb.UpdateSysUserResp, error) {
	// 修改账户信息
	account, err := l.svcCtx.SysUser.FindOne(l.ctx, in.Id)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			l.Logger.Errorf("账号不存在: %v", err)
			return nil, code.FindAccountErr
		}
		l.Logger.Errorf("查询账号失败: %v", err)
		return nil, code.FindAccountErr
	}
	if in.UserName != "" {
		account.UserName = in.UserName
	}
	if in.Mobile != "" {
		account.Mobile = in.Mobile
	}
	if in.Email != "" {
		account.Email = in.Email
	}

	err = l.svcCtx.SysUser.Update(l.ctx, account)
	if err != nil {
		l.Logger.Errorf("修改账号信息失败: %v", err)
		return nil, code.UpdateSysUserErr
	}
	return &pb.UpdateSysUserResp{}, nil
}
