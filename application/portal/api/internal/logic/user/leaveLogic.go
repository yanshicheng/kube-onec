package user

import (
	"context"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/client/sysuserservice"

	"github.com/yanshicheng/kube-onec/application/portal/api/internal/svc"
	"github.com/yanshicheng/kube-onec/application/portal/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LeaveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLeaveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LeaveLogic {
	return &LeaveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LeaveLogic) Leave(req *types.LeaveRequest) (resp string, err error) {
	_, err = l.svcCtx.SysUserRpc.Leave(l.ctx, &sysuserservice.LeaveReq{Id: req.Id})
	if err != nil {
		l.Logger.Errorf("设置用户离职状态失败: %v", err)
		return
	}
	return "离职状态设置成功!", nil
}
