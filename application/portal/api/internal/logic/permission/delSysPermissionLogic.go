package permission

import (
	"context"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/client/syspermissionservice"

	"github.com/yanshicheng/kube-onec/application/portal/api/internal/svc"
	"github.com/yanshicheng/kube-onec/application/portal/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelSysPermissionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDelSysPermissionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelSysPermissionLogic {
	return &DelSysPermissionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DelSysPermissionLogic) DelSysPermission(req *types.DefaultIdRequest) (resp string, err error) {
	_, err = l.svcCtx.SysPermissionRpc.DelSysPermission(l.ctx, &syspermissionservice.DelSysPermissionReq{
		Id: req.Id,
	})
	if err != nil {
		l.Logger.Errorf("删除权限失败: %v", err)
		return
	}
	return "删除权限成功", nil
}
