package sysuserservicelogic

import (
	"context"
	"github.com/yanshicheng/kube-onec/common/handler/errorx"

	"github.com/yanshicheng/kube-onec/application/portal/rpc/internal/svc"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelSysUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelSysUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelSysUserLogic {
	return &DelSysUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelSysUserLogic) DelSysUser(in *pb.DelSysUserReq) (*pb.DelSysUserResp, error) {
	// 软删除 账户信息
	err := l.svcCtx.SysUser.DeleteSoft(l.ctx, in.Id)
	if err != nil {
		l.Logger.Errorf("删除账号失败: %v", err)
		return nil, errorx.DatabaseDeleteErr
	}
	return &pb.DelSysUserResp{}, nil
}
