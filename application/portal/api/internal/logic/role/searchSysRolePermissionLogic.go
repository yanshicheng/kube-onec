package role

import (
	"context"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/client/sysroleservice"

	"github.com/yanshicheng/kube-onec/application/portal/api/internal/svc"
	"github.com/yanshicheng/kube-onec/application/portal/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchSysRolePermissionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchSysRolePermissionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchSysRolePermissionLogic {
	return &SearchSysRolePermissionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchSysRolePermissionLogic) SearchSysRolePermission(req *types.SearchSysRolePermissionRequest) (resp *types.SearchSysRolePermissionResponse, err error) {
	res, err := l.svcCtx.SysRoleRpc.SearchRolePermission(l.ctx, &sysroleservice.SearchRolePermissionReq{
		RoleId: req.RoleId,
	})
	if err != nil {
		l.Logger.Errorf("查询角色权限失败: 请求=%+v, 错误=%v", req, err)
		return
	}
	return &types.SearchSysRolePermissionResponse{
		PermissionIds: res.Data,
	}, nil
}
