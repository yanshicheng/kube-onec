package syspermissionservicelogic

import (
	"context"
	"errors"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/internal/code"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/internal/model"
	"github.com/yanshicheng/kube-onec/common/handler/errorx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"github.com/yanshicheng/kube-onec/application/portal/rpc/internal/svc"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelSysPermissionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelSysPermissionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelSysPermissionLogic {
	return &DelSysPermissionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelSysPermissionLogic) DelSysPermission(in *pb.DelSysPermissionReq) (*pb.DelSysPermissionResp, error) {
	// 先查询是否有子集
	childPerms, err := l.svcCtx.SysPermission.SearchNoPage(l.ctx, "", true, "parent_id = ?", in.Id)
	if err != nil && !errors.Is(err, model.ErrNotFound) {
		l.Logger.Errorf("查询子集权限失败: %v", err)
		return nil, code.GetChildPermissionErr
	}
	if len(childPerms) > 0 {
		return nil, code.DeletePermissionHasChildErr
	}
	// 进行软删除权限，删除权限之前需要清理掉，角色和权限绑定的数据
	sqlStr := "permission_id = ?"
	userPerms, err := l.svcCtx.SysRolePermission.SearchNoPage(l.ctx, "", true, sqlStr, in.Id)
	if err != nil && !errors.Is(err, model.ErrNotFound) {
		l.Logger.Errorf("查询角色权限失败: %v", err)
		return nil, code.FindRolePermissionErr
	}
	// 如果没有关联则直接删除
	if len(userPerms) == 0 {
		err = l.svcCtx.SysPermission.DeleteSoft(l.ctx, in.Id)
		if err != nil {
			if errors.Is(err, model.ErrNotFound) {
				return nil, errorx.DatabaseNotFound
			}
			l.Logger.Errorf("删除权限失败: %v", err)
			return nil, errorx.DatabaseDeleteErr
		}
	}
	// 如果有关联则开始事务先清理关联数据
	if txnErr := l.svcCtx.SysPermission.TransCtx(l.ctx, func(ctx context.Context, session sqlx.Session) error {
		for _, user := range userPerms {
			sqlStr := "DELETE FROM {table} WHERE permission_id = ?"
			result, err := l.svcCtx.SysRolePermission.TransOnSql(ctx, session, user.Id, sqlStr, in.Id)
			if err != nil {
				return err
			}
			rowsAffected, err := result.RowsAffected()
			if err != nil {
				return err
			}
			if rowsAffected == 0 {
				return errors.New("删除角色权限失败")
			}
		}

		// 删除权限
		sqlStr := "DELETE FROM {table} WHERE id = ?"
		_, err := l.svcCtx.SysPermission.TransOnSql(l.ctx, session, in.Id, sqlStr, in.Id)
		if err != nil {
			return err
		}
		l.Logger.Infof("删除权限成功: %v", in.Id)
		return nil
	}); txnErr != nil {
		return nil, txnErr
	}
	return &pb.DelSysPermissionResp{}, nil
}
