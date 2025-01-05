package role

import (
	"context"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/client/sysroleservice"

	"github.com/yanshicheng/kube-onec/application/portal/api/internal/svc"
	"github.com/yanshicheng/kube-onec/application/portal/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type BindSysRolePermissionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBindSysRolePermissionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BindSysRolePermissionLogic {
	return &BindSysRolePermissionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BindSysRolePermissionLogic) BindSysRolePermission(req *types.BindSysRolePermissionRequest) (resp string, err error) {
	_, err = l.svcCtx.SysRoleRpc.BindRolePermission(l.ctx, &sysroleservice.BindRolePermissionReq{
		RoleId:        req.RoleId,
		PermissionIds: req.PermissionIds,
	})
	if err != nil {
		l.Logger.Errorf("绑定角色权限失败: 请求=%+v, 错误=%v", req, err)
		return
	}
	return
}
