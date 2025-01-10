package sysuserservicelogic

import (
	"context"
	"errors"
	"github.com/yanshicheng/kube-onec/pkg/utils"

	"github.com/yanshicheng/kube-onec/application/portal/rpc/internal/code"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/internal/model"
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"github.com/yanshicheng/kube-onec/application/portal/rpc/internal/svc"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type BindRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBindRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BindRoleLogic {
	return &BindRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}
func (l *BindRoleLogic) BindRole(in *pb.BindRoleReq) (*pb.BindRoleResp, error) {
	// 查询账号是否存在
	account, err := l.svcCtx.SysUser.FindOne(l.ctx, in.Id)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			l.Logger.Errorf("账号不存在: %v", err)
			return nil, code.FindAccountErr
		}
		l.Logger.Errorf("查询账号失败: %v", err)
		return nil, code.FindAccountErr
	}

	roles, err := l.svcCtx.SysRole.SearchNoPage(l.ctx, "id", false, utils.BuildInCondition("id", len(in.RoleIds)), utils.ConvertUint64SliceToInterfaceSlice(in.RoleIds)...)
	if err != nil {
		l.Logger.Errorf("批量查询角色失败: %v", err)
		return nil, code.FindRoleErr
	}
	if len(roles) != len(in.RoleIds) {
		l.Logger.Errorf("角色不存在: %v", err)
		return nil, code.ParameterIllegal
	}
	// 将角色存入 map，方便后续查找
	// 开启事务
	err = l.svcCtx.SysUserRole.TransCtx(l.ctx, func(ctx context.Context, session sqlx.Session) error {
		// 批量删除
		deleteSql := "DELETE FROM {table} WHERE user_id = ?"
		_, err := l.svcCtx.SysUserRole.TransOnSql(ctx, session, 0, deleteSql, account.Id)
		if err != nil && !errors.Is(err, model.ErrNotFound) {
			l.Logger.Errorf("批量删除用户角色关系失败: %v", err)
			return code.BindRoleErr
		}
		// 批量插入
		insertSql := "INSERT INTO {table} (user_id, role_id) VALUES (?, ?)"
		for _, roleId := range roles {
			_, err := l.svcCtx.SysUserRole.TransOnSql(ctx, session, 0, insertSql, account.Id, roleId.Id)
			if err != nil {
				l.Logger.Errorf("批量插入用户角色关系失败: %v", err)
				return code.BindRoleErr
			}
		}
		return err
	})
	if err != nil {
		l.Logger.Errorf("绑定角色失败: %v", err)
		return nil, code.BindRoleErr
	}
	l.Logger.Infof("用户 %s 成功绑定角色: %v", account.Account, in.RoleIds)
	// 定义一个批量绑定角色的闭包函数
	return &pb.BindRoleResp{}, nil
}
