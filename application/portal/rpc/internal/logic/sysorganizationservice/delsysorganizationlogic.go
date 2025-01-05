package sysorganizationservicelogic

import (
	"context"
	"errors"
	"fmt"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/internal/code"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/internal/model"
	"github.com/yanshicheng/kube-onec/common/handler/errorx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"github.com/yanshicheng/kube-onec/application/portal/rpc/internal/svc"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelSysOrganizationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelSysOrganizationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelSysOrganizationLogic {
	return &DelSysOrganizationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelSysOrganizationLogic) DelSysOrganization(in *pb.DelSysOrganizationReq) (*pb.DelSysOrganizationResp, error) {
	// 查询这个机构下面是否还有其他机构
	OrSqlStr := "parent_id = ?"
	res, err := l.svcCtx.SysOrganization.SearchNoPage(l.ctx, "", true, OrSqlStr, in.Id)
	if err != nil {
		if !errors.Is(err, model.ErrNotFound) {
			l.Logger.Errorf("查询机构信息失败: %v", err)
			return nil, errorx.DatabaseQueryErr
		}
	}
	if len(res) > 0 {
		l.Logger.Errorf("删除机构失败, 机构下还有子机构")
		return nil, code.DeleteOrganizationNotNullErr
	}

	// 删除机构需要清理用户的机构信息
	sqlStr := "organization_id = ?"
	userAll, err := l.svcCtx.SysUser.SearchNoPage(l.ctx, "", true, sqlStr, in.Id)
	if err != nil {
		if !errors.Is(err, model.ErrNotFound) {
			l.Logger.Errorf("查询关联用户失败: %v", err)
			return nil, errorx.DatabaseQueryErr
		}
	}
	// 如果没有相关用户，直接删除
	if len(userAll) == 0 {
		// 软删除机构
		err = l.svcCtx.SysOrganization.DeleteSoft(l.ctx, in.Id)
		if err != nil {
			if errors.Is(err, model.ErrNotFound) {
				return nil, errorx.DatabaseQueryErr
			}
			l.Logger.Errorf("删除机构失败: %v", err)
			return nil, errorx.DatabaseDeleteErr
		}
		return &pb.DelSysOrganizationResp{}, nil
	}
	// 如果有相关用户，开启事务进行批量更新
	if txnErr := l.svcCtx.SysOrganization.TransCtx(l.ctx, func(ctx context.Context, session sqlx.Session) error {
		for _, user := range userAll {
			// 编写 更新 机构 sql,使用占位符
			query := "UPDATE {table} SET organization_id = ? WHERE id = ?"
			result, err := l.svcCtx.SysUser.TransOnSql(ctx, session, user.Id, query, 0, user.Id)
			if err != nil {
				l.Logger.Errorf("清理用户机构失败, 用户ID: %d, 错误: %v", user.Id, err)
				return err
			}
			// 检查影响的行数
			rowsAffected, err := result.RowsAffected()
			if err != nil {
				l.Logger.Errorf("无法获取影响的行数, 用户ID=%d, 错误: %v", user.Id, err)
				return err
			}
			if rowsAffected == 0 {
				l.Logger.Errorf("清理用户机构失败, 没有影响任何行, 用户ID=%d", user.Id)
				return fmt.Errorf("清理用户机构失败, 没有影响任何行, 用户ID=%d", user.Id)
			}
		}

		// 删除机构
		// 构造 sql
		query := "UPDATE {table} set delete_time = NOW() where `id` = ?"
		_, err := l.svcCtx.SysOrganization.TransOnSql(l.ctx, session, in.Id, query, in.Id)
		if err != nil {
			l.Logger.Errorf("删除职位失败: %v", err)
			return err
		}
		return nil
	}); txnErr != nil {
		return nil, txnErr
	}
	return &pb.DelSysOrganizationResp{}, nil
}
