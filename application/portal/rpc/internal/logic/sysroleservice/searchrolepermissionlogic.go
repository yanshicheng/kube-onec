package sysroleservicelogic

import (
	"context"
	"errors"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/internal/code"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/internal/model"
	"github.com/yanshicheng/kube-onec/utils"

	"github.com/yanshicheng/kube-onec/application/portal/rpc/internal/svc"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchRolePermissionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchRolePermissionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchRolePermissionLogic {
	return &SearchRolePermissionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchRolePermissionLogic) SearchRolePermission(in *pb.SearchRolePermissionReq) (*pb.SearchRolePermissionResp, error) {
	// 查询角色是否存在
	role, err := l.svcCtx.SysRole.FindOne(l.ctx, in.RoleId)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			l.Logger.Errorf("角色不存在: %v", err)
			return nil, code.FindRoleErr
		}
		l.Logger.Errorf("查询角色失败: %v", err)
		return nil, code.FindRoleErr
	}

	// 批量查询角色权限关系
	rolePermissions, err := l.svcCtx.SysRolePermission.SearchNoPage(l.ctx, "id", false, "`role_id`=?", role.Id)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			l.Logger.Errorf("角色权限关系不存在: %v", err)
			return &pb.SearchRolePermissionResp{}, nil
		}
		l.Logger.Errorf("批量查询角色权限关系失败: %v", err)
		return nil, code.FindRolePermissionListErr
	}
	rolePreIds := make([]uint64, len(rolePermissions))
	for i, rolePermission := range rolePermissions {
		rolePreIds[i] = rolePermission.PermissionId
	}
	// 通过 角色权限关系 批量查询权限
	permissions, err := l.svcCtx.SysPermission.SearchNoPage(l.ctx, "id", false, utils.BuildInCondition("id", len(rolePreIds)), utils.ConvertUint64SliceToInterfaceSlice(rolePreIds)...)

	if err != nil {
		l.Logger.Errorf("批量查询权限失败: %v", err)
		return nil, code.FindRolePermissionListErr
	}
	dataIds := make([]uint64, len(permissions))
	for i, permission := range permissions {
		dataIds[i] = permission.Id
	}
	return &pb.SearchRolePermissionResp{
		Data: dataIds,
	}, nil
}
