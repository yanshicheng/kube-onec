package user

import (
	"context"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/pb"

	"github.com/yanshicheng/kube-onec/application/portal/api/internal/svc"
	"github.com/yanshicheng/kube-onec/application/portal/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FrozenAccountsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFrozenAccountsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FrozenAccountsLogic {
	return &FrozenAccountsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FrozenAccountsLogic) FrozenAccounts(req *types.FrozenAccountsRequest) (resp string, err error) {
	_, err = l.svcCtx.SysUserRpc.FrozenAccounts(l.ctx, &pb.FrozenAccountsReq{
		Id:         req.Id,
		IsDisabled: req.IsDisabled,
	})
	if err != nil {
		l.Logger.Errorf("冻结/解冻账号失败: %v", err)
		return "", err
	}
	l.Logger.Infof("冻结/解冻账号成功: %v", req.Id)
	return "冻结/解冻账号成功", nil
}
