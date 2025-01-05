package user

import (
	"context"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/client/sysuserservice"

	"github.com/yanshicheng/kube-onec/application/portal/api/internal/svc"
	"github.com/yanshicheng/kube-onec/application/portal/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelSysUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDelSysUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelSysUserLogic {
	return &DelSysUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DelSysUserLogic) DelSysUser(req *types.DefaultIdRequest) (resp string, err error) {
	_, err = l.svcCtx.SysUserRpc.DelSysUser(l.ctx, &sysuserservice.DelSysUserReq{
		Id: req.Id,
	})
	if err != nil {
		l.Logger.Errorf("删除用户失败: %v", err)
		return "", err
	}
	l.Logger.Infof("删除用户成功: %v", req)
	return
}
