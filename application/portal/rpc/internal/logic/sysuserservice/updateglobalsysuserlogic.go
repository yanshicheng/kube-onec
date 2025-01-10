package sysuserservicelogic

import (
	"context"
	"errors"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/internal/code"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/internal/model"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/internal/svc"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/pb"
	"github.com/yanshicheng/kube-onec/pkg/utils"
	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateGlobalSysUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateGlobalSysUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateGlobalSysUserLogic {
	return &UpdateGlobalSysUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateGlobalSysUserLogic) UpdateGlobalSysUser(in *pb.UpdateGlobalSysUserReq) (*pb.UpdateGlobalSysUserResp, error) {
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
	// 时间戳转为 time.time
	if in.UserName != "" {
		account.UserName = in.UserName
	}
	if in.Mobile != "" {
		account.Mobile = in.Mobile
	}
	if in.Email != "" {
		account.Email = in.Email
	}
	if in.WorkNumber != "" {
		account.WorkNumber = in.WorkNumber
	}
	if in.HireDate != 0 {
		account.HireDate = utils.FormattedDate(in.HireDate)
	}
	if in.PositionId != 0 {
		account.PositionId = in.PositionId
	}
	if in.OrganizationId != 0 {
		account.OrganizationId = in.OrganizationId
	}
	err = l.svcCtx.SysUser.Update(l.ctx, account)
	if err != nil {
		l.Logger.Errorf("修改账号信息失败: %v", err)
		return nil, code.UpdateSysUserErr
	}

	return &pb.UpdateGlobalSysUserResp{}, nil
}
