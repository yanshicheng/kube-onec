package role

import (
	"context"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/client/sysroleservice"

	"github.com/yanshicheng/kube-onec/application/portal/api/internal/svc"
	"github.com/yanshicheng/kube-onec/application/portal/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelSysRoleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDelSysRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelSysRoleLogic {
	return &DelSysRoleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DelSysRoleLogic) DelSysRole(req *types.DefaultIdRequest) (resp string, err error) {

	res, err := l.svcCtx.SysRoleRpc.DelSysRole(l.ctx, &sysroleservice.DelSysRoleReq{
		Id: req.Id,
	})
	if err != nil {
		logx.WithContext(l.ctx).Errorf("删除角色失败: %v", err)
		return "", err
	}
	logx.WithContext(l.ctx).Infof("删除角色成功: %v", res)
	return "删除用户成功", nil
}
