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

type FrozenAccountsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFrozenAccountsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FrozenAccountsLogic {
	return &FrozenAccountsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FrozenAccountsLogic) FrozenAccounts(in *pb.FrozenAccountsReq) (*pb.FrozenAccountsResp, error) {
	// 先查询账号
	account, err := l.svcCtx.SysUser.FindOne(l.ctx, in.Id)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			l.Logger.Errorf("账号不存在: %v", err)
			return nil, code.FindAccountErr
		}
		l.Logger.Errorf("查询账号失败: %v", err)
		return nil, code.FindAccountErr
	}
	// 冻结 或 解冻账号
	account.IsDisabled = in.IsDisabled
	err = l.svcCtx.SysUser.Update(l.ctx, account)
	if err != nil {
		l.Logger.Errorf("冻结账号失败: %v", err)
		return nil, code.FrozenAccountsErr
	}
	return &pb.FrozenAccountsResp{}, nil
}
