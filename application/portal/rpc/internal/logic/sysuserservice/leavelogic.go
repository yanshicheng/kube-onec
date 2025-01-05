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

type LeaveLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLeaveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LeaveLogic {
	return &LeaveLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LeaveLogic) Leave(in *pb.LeaveReq) (*pb.LeaveResp, error) {
	// 设置账号离职
	account, err := l.svcCtx.SysUser.FindOne(l.ctx, in.Id)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			l.Logger.Errorf("账号不存在: %v", err)
			return nil, code.FindAccountErr
		}
		l.Logger.Errorf("查询账号失败: %v", err)
		return nil, code.FindAccountErr
	}
	account.IsLeave = 1

	err = l.svcCtx.SysUser.Update(l.ctx, account)
	if err != nil {
		l.Logger.Errorf("离职账号失败: %v", err)
		return nil, code.LeaveErr
	}
	return &pb.LeaveResp{}, nil
}
