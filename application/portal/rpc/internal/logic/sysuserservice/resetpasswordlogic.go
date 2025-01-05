package sysuserservicelogic

import (
	"context"
	"errors"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/internal/code"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/internal/model"
	"github.com/yanshicheng/kube-onec/utils"

	"github.com/yanshicheng/kube-onec/application/portal/rpc/internal/svc"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ResetPasswordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewResetPasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ResetPasswordLogic {
	return &ResetPasswordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ResetPasswordLogic) ResetPassword(in *pb.ResetPasswordReq) (*pb.ResetPasswordResp, error) {
	// 查询账号
	account, err := l.svcCtx.SysUser.FindOne(l.ctx, in.Id)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			l.Logger.Errorf("账号不存在: %v", err)
			return nil, code.FindAccountErr
		}
		l.Logger.Errorf("查询账号失败: %v", err)
		return nil, code.FindAccountErr
	}

	//判断是否禁用
	if account.IsDisabled == 1 {
		l.Logger.Errorf("账号已禁用: %v", err)
		return nil, code.AccountLockedErr
	}
	// 查看账号是否离职
	if account.IsLeave == 1 {
		l.Logger.Infof("账号已离职, 禁止登陆。 账号: %s", account.Account)
		return nil, code.AccountLockedTip
	}
	encryptPassword, err := utils.EncryptPassword(utils.GeneratePassword())
	if err != nil {
		l.Logger.Errorf("密码加密失败: %v", err)
		return nil, code.EncryptPasswordErr
	}
	account.Password = encryptPassword
	account.IsResetPassword = 1
	err = l.svcCtx.SysUser.Update(l.ctx, account)
	if err != nil {
		l.Logger.Errorf("重置密码失败: %v", err)
		return nil, code.ResetPasswordErr
	}
	return &pb.ResetPasswordResp{}, nil
}
