package sysroleservicelogic

import (
	"context"
	"errors"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/internal/code"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/internal/model"
	"github.com/yanshicheng/kube-onec/pkg/utils"
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"github.com/yanshicheng/kube-onec/application/portal/rpc/internal/svc"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type BindRolePermissionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBindRolePermissionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BindRolePermissionLogic {
	return &BindRolePermissionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *BindRolePermissionLogic) BindRolePermission(in *pb.BindRolePermissionReq) (*pb.BindRolePermissionResp, error) {
	// 先查询 role 是否存在
	role, err := l.svcCtx.SysRole.FindOne(l.ctx, in.RoleId)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			l.Logger.Errorf("角色不存在: %v", err)
			return nil, code.FindRoleErr
		}
		l.Logger.Errorf("查询角色失败: %v", err)
		return nil, code.FindRoleErr
	}

	//批量查询角色权限关系
	permissions, err := l.svcCtx.SysPermission.SearchNoPage(l.ctx, "id", false, utils.BuildInCondition("id", len(in.PermissionIds)), utils.ConvertUint64SliceToInterfaceSlice(in.PermissionIds)...)
	if err != nil {
		l.Logger.Errorf("批量查询角色权限关系失败: %v", err)
		return nil, code.FindRolePermissionListErr
	}

	// 开启事务
	err = l.svcCtx.SysRolePermission.TransCtx(l.ctx, func(ctx context.Context, session sqlx.Session) error {
		// 批量删除角色权限关系
		deleteSql := "DELETE FROM {table} WHERE role_id = ?"
		_, err := l.svcCtx.SysRolePermission.TransOnSql(ctx, session, 0, deleteSql, role.Id)
		if err != nil && !errors.Is(err, model.ErrNotFound) {
			l.Logger.Errorf("批量删除角色权限关系失败: %v", err)
			return code.DelBindRolePermissionErr
		}

		// 批量绑定角色权限关系
		for _, permission := range permissions {
			insertSql := "INSERT INTO {table} (role_id, permission_id) VALUES (?, ?)"
			_, err := l.svcCtx.SysRolePermission.TransOnSql(ctx, session, 0, insertSql, role.Id, permission.Id)
			if err != nil {
				l.Logger.Errorf("批量绑定角色权限关系失败: %v", err)
				return code.BindRolePermissionErr
			}
		}
		return nil
	})
	if err != nil {
		l.Logger.Errorf("绑定角色权限失败: %v", err)
		return nil, code.BindRolePermissionErr
	}
	return &pb.BindRolePermissionResp{}, nil
}
