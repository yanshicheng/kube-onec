package syspositionservicelogic

import (
	"context"
	"errors"
	"fmt"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/internal/model"
	"github.com/yanshicheng/kube-onec/common/handler/errorx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"github.com/yanshicheng/kube-onec/application/portal/rpc/internal/svc"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelSysPositionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelSysPositionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelSysPositionLogic {
	return &DelSysPositionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelSysPositionLogic) DelSysPosition(in *pb.DelSysPositionReq) (*pb.DelSysPositionResp, error) {
	// 查询 职位 和 角色 关联数据 同时清理
	//构造sql
	//queryStr := fmt.Sprintf("AND role_name LIKE '%%%s%%'", in.RoleName)
	sqlStr := "position_id = ?"
	userAll, err := l.svcCtx.SysUser.SearchNoPage(l.ctx, "", true, sqlStr, in.Id)
	if err != nil {
		if !errors.Is(err, model.ErrNotFound) {
			l.Logger.Errorf("查询关联用户失败: %v", err)
			return nil, errorx.DatabaseQueryErr
		}
	}
	// 如果没有相关用户，直接删除
	if len(userAll) == 0 {
		// 删除职位
		err = l.svcCtx.SysPosition.Delete(l.ctx, in.Id)
		if err != nil {
			if errors.Is(err, model.ErrNotFound) {
				return nil, errorx.DatabaseNotFound
			}
			l.Logger.Errorf("删除职位失败: %v", err)
			return nil, errorx.DatabaseDeleteErr
		}
		return &pb.DelSysPositionResp{}, nil
	}
	// 如果有相关用户，开启事务进行批量更新
	if txnErr := l.svcCtx.SysPosition.TransCtx(l.ctx, func(ctx context.Context, session sqlx.Session) error {
		// 遍历用户并清理职位数据
		for _, user := range userAll {
			// 编写 更新 职位 sql,使用占位符
			query := "UPDATE {table} SET position_id = ? WHERE id = ?"
			result, err := l.svcCtx.SysUser.TransOnSql(ctx, session, user.Id, query, 0, user.Id)
			if err != nil {
				logx.Errorf("清理用户职位失败, 用户ID: %d, 错误: %v", user.Id, err)
				return err
			}
			// 检查影响的行数
			rowsAffected, err := result.RowsAffected()
			if err != nil {
				logx.Errorf("无法获取影响的行数, 用户ID=%d, 错误: %v", user.Id, err)
				return err
			}
			if rowsAffected == 0 {
				logx.Errorf("清理用户职位失败, 没有影响任何行, 用户ID=%d", user.Id)
				return fmt.Errorf("no rows affected for user ID %d", user.Id)
			}
		}

		// 删除职位
		// 构造 sql
		query := "DELETE FROM {table} WHERE id = ?"
		_, err := l.svcCtx.SysPosition.TransOnSql(l.ctx, session, in.Id, query, in.Id)
		if err != nil {
			l.Logger.Errorf("删除职位失败: %v", err)
			return err
		}
		return nil
	}); txnErr != nil {
		l.Logger.Errorf("事务执行失败: %v", txnErr)
		return nil, errorx.DatabaseTransactionErr
	}
	return &pb.DelSysPositionResp{}, nil
}
