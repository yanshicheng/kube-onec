// Code generated by goctl. DO NOT EDIT.
// versions:
//  goctl version: 1.7.3

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	sysUserFieldNames          = builder.RawFieldNames(&SysUser{})
	sysUserRows                = strings.Join(sysUserFieldNames, ",")
	sysUserRowsExpectAutoSet   = strings.Join(stringx.Remove(sysUserFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	sysUserRowsWithPlaceHolder = strings.Join(stringx.Remove(sysUserFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheKubeOnecSysUserIdPrefix         = "cache:kubeOnec:sysUser:id:"
	cacheKubeOnecSysUserAccountPrefix    = "cache:kubeOnec:sysUser:account:"
	cacheKubeOnecSysUserEmailPrefix      = "cache:kubeOnec:sysUser:email:"
	cacheKubeOnecSysUserMobilePrefix     = "cache:kubeOnec:sysUser:mobile:"
	cacheKubeOnecSysUserWorkNumberPrefix = "cache:kubeOnec:sysUser:workNumber:"
)

type (
	sysUserModel interface {
		Insert(ctx context.Context, data *SysUser) (sql.Result, error)

		FindOne(ctx context.Context, id uint64) (*SysUser, error)
		Search(ctx context.Context, orderStr string, isAsc bool, page, pageSize uint64, queryStr string, args ...any) ([]*SysUser, uint64, error)
		SearchNoPage(ctx context.Context, orderStr string, isAsc bool, queryStr string, args ...any) ([]*SysUser, error)
		FindOneByAccount(ctx context.Context, account string) (*SysUser, error)
		FindOneByEmail(ctx context.Context, email string) (*SysUser, error)
		FindOneByMobile(ctx context.Context, mobile string) (*SysUser, error)
		FindOneByWorkNumber(ctx context.Context, workNumber string) (*SysUser, error)
		Update(ctx context.Context, data *SysUser) error
		Delete(ctx context.Context, id uint64) error
		DeleteSoft(ctx context.Context, id uint64) error
		TransCtx(ctx context.Context, fn func(context.Context, sqlx.Session) error) error
		TransOnSql(ctx context.Context, session sqlx.Session, id uint64, sqlStr string, args ...any) (sql.Result, error)
		ExecSql(ctx context.Context, id uint64, sqlStr string, args ...any) (sql.Result, error)
	}

	defaultSysUserModel struct {
		sqlc.CachedConn
		table string
	}

	SysUser struct {
		Id              uint64    `db:"id"`                // 自增主键
		UserName        string    `db:"user_name"`         // 用户姓名
		Account         string    `db:"account"`           // 用户账号，唯一标识
		Password        string    `db:"password"`          // 用户密码，需加密存储
		Icon            string    `db:"icon"`              // 用户头像URL
		Mobile          string    `db:"mobile"`            // 用户手机号
		Email           string    `db:"email"`             // 用户邮箱地址
		WorkNumber      string    `db:"work_number"`       // 用户工号
		HireDate        time.Time `db:"hire_date"`         // 入职日期
		IsResetPassword int64     `db:"Is_reset_password"` // 是否需要重置密码,false否true是
		IsDisabled      int64     `db:"is_disabled"`       // 是否禁用,false否true是
		IsLeave         int64     `db:"is_leave"`          // 是否离职,false否true是
		PositionId      uint64    `db:"position_id"`       // 职位ID，关联职位表
		OrganizationId  uint64    `db:"organization_id"`   // 组织ID，关联组织表
		LastLoginTime   time.Time `db:"last_login_time"`   // 上次登录时间
		CreatedAt       time.Time `db:"created_at"`        // 记录创建时间
		UpdatedAt       time.Time `db:"updated_at"`        // 记录更新时间
		IsDeleted       int64     `db:"is_deleted"`        // 是否删除
	}
)

func newSysUserModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultSysUserModel {
	return &defaultSysUserModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      "`sys_user`",
	}
}

func (m *defaultSysUserModel) Delete(ctx context.Context, id uint64) error {
	data, err := m.FindOne(ctx, id)
	if err != nil {
		return err
	}

	kubeOnecSysUserAccountKey := fmt.Sprintf("%s%v", cacheKubeOnecSysUserAccountPrefix, data.Account)
	kubeOnecSysUserEmailKey := fmt.Sprintf("%s%v", cacheKubeOnecSysUserEmailPrefix, data.Email)
	kubeOnecSysUserIdKey := fmt.Sprintf("%s%v", cacheKubeOnecSysUserIdPrefix, id)
	kubeOnecSysUserMobileKey := fmt.Sprintf("%s%v", cacheKubeOnecSysUserMobilePrefix, data.Mobile)
	kubeOnecSysUserWorkNumberKey := fmt.Sprintf("%s%v", cacheKubeOnecSysUserWorkNumberPrefix, data.WorkNumber)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, kubeOnecSysUserAccountKey, kubeOnecSysUserEmailKey, kubeOnecSysUserIdKey, kubeOnecSysUserMobileKey, kubeOnecSysUserWorkNumberKey)
	return err
}

func (m *defaultSysUserModel) DeleteSoft(ctx context.Context, id uint64) error {
	data, err := m.FindOne(ctx, id)
	if err != nil {
		return err
	}
	// 如果记录已软删除，无需再次删除
	if data.IsDeleted == 1 {
		return nil
	}
	kubeOnecSysUserAccountKey := fmt.Sprintf("%s%v", cacheKubeOnecSysUserAccountPrefix, data.Account)
	kubeOnecSysUserEmailKey := fmt.Sprintf("%s%v", cacheKubeOnecSysUserEmailPrefix, data.Email)
	kubeOnecSysUserIdKey := fmt.Sprintf("%s%v", cacheKubeOnecSysUserIdPrefix, id)
	kubeOnecSysUserMobileKey := fmt.Sprintf("%s%v", cacheKubeOnecSysUserMobilePrefix, data.Mobile)
	kubeOnecSysUserWorkNumberKey := fmt.Sprintf("%s%v", cacheKubeOnecSysUserWorkNumberPrefix, data.WorkNumber)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set `is_deleted` = 1 where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, kubeOnecSysUserAccountKey, kubeOnecSysUserEmailKey, kubeOnecSysUserIdKey, kubeOnecSysUserMobileKey, kubeOnecSysUserWorkNumberKey)
	return err
}

func (m *defaultSysUserModel) TransCtx(ctx context.Context, fn func(context.Context, sqlx.Session) error) error {
	return m.TransactCtx(ctx, func(ctx context.Context, session sqlx.Session) error {
		return fn(ctx, session)
	})
}

func (m *defaultSysUserModel) TransOnSql(ctx context.Context, session sqlx.Session, id uint64, sqlStr string, args ...any) (sql.Result, error) {
	query := strings.ReplaceAll(sqlStr, "{table}", m.table)
	// 如果 id != 0 并且启用了缓存逻辑
	if !isZeroValue(id) {
		// 查询数据（如果需要，确保数据存在）
		data, err := m.FindOne(ctx, id)
		if err != nil {
			return nil, err
		}

		// 缓存相关处理

		kubeOnecSysUserAccountKey := fmt.Sprintf("%s%v", cacheKubeOnecSysUserAccountPrefix, data.Account)
		kubeOnecSysUserEmailKey := fmt.Sprintf("%s%v", cacheKubeOnecSysUserEmailPrefix, data.Email)
		kubeOnecSysUserIdKey := fmt.Sprintf("%s%v", cacheKubeOnecSysUserIdPrefix, id)
		kubeOnecSysUserMobileKey := fmt.Sprintf("%s%v", cacheKubeOnecSysUserMobilePrefix, data.Mobile)
		kubeOnecSysUserWorkNumberKey := fmt.Sprintf("%s%v", cacheKubeOnecSysUserWorkNumberPrefix, data.WorkNumber) // 处理缓存逻辑，例如删除或更新缓存
		// 执行带缓存处理的 SQL 操作
		return m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (sql.Result, error) {
			return session.ExecCtx(ctx, query, args...)
		}, kubeOnecSysUserAccountKey, kubeOnecSysUserEmailKey, kubeOnecSysUserIdKey, kubeOnecSysUserMobileKey, kubeOnecSysUserWorkNumberKey) // 传递缓存相关的键值

	}

	// 如果 id == 0 或不需要缓存，直接执行 SQL
	return m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (sql.Result, error) {
		return session.ExecCtx(ctx, query, args...)
	})
}

func (m *defaultSysUserModel) ExecSql(ctx context.Context, id uint64, sqlStr string, args ...any) (sql.Result, error) {
	// 如果 id != 0 并且启用了缓存逻辑
	query := strings.ReplaceAll(sqlStr, "{table}", m.table)
	if !isZeroValue(id) {
		// 缓存相关处理

		// 查询数据（如果需要，确保数据存在）
		data, err := m.FindOne(ctx, id)
		if err != nil {
			return nil, err
		}

		kubeOnecSysUserAccountKey := fmt.Sprintf("%s%v", cacheKubeOnecSysUserAccountPrefix, data.Account)
		kubeOnecSysUserEmailKey := fmt.Sprintf("%s%v", cacheKubeOnecSysUserEmailPrefix, data.Email)
		kubeOnecSysUserIdKey := fmt.Sprintf("%s%v", cacheKubeOnecSysUserIdPrefix, id)
		kubeOnecSysUserMobileKey := fmt.Sprintf("%s%v", cacheKubeOnecSysUserMobilePrefix, data.Mobile)
		kubeOnecSysUserWorkNumberKey := fmt.Sprintf("%s%v", cacheKubeOnecSysUserWorkNumberPrefix, data.WorkNumber) // 处理缓存逻辑，例如删除或更新缓存
		// 执行带缓存处理的 SQL 操作
		return m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (sql.Result, error) {
			return conn.ExecCtx(ctx, query, args...)
		}, kubeOnecSysUserAccountKey, kubeOnecSysUserEmailKey, kubeOnecSysUserIdKey, kubeOnecSysUserMobileKey, kubeOnecSysUserWorkNumberKey) // 传递缓存相关的键值

	}

	// 如果 id == 0 或不需要缓存，直接执行 SQL
	return m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (sql.Result, error) {
		return conn.ExecCtx(ctx, query, args...)
	})
}
func (m *defaultSysUserModel) FindOne(ctx context.Context, id uint64) (*SysUser, error) {
	kubeOnecSysUserIdKey := fmt.Sprintf("%s%v", cacheKubeOnecSysUserIdPrefix, id)
	var resp SysUser
	err := m.QueryRowCtx(ctx, &resp, kubeOnecSysUserIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? AND `is_deleted` = 0 limit 1", sysUserRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultSysUserModel) Search(ctx context.Context, orderStr string, isAsc bool, page, pageSize uint64, queryStr string, args ...any) ([]*SysUser, uint64, error) {
	// 确保分页参数有效
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 20
	}

	// 构造查询条件
	// 添加 `is_deleted` = 0 条件，保证只查询未软删除数据
	// 初始化 WHERE 子句
	where := "WHERE `is_deleted` = 0"
	if queryStr != "" {
		where = fmt.Sprintf("WHERE %s AND `is_deleted` = 0", queryStr)
	}

	// 根据 isAsc 参数确定排序方式
	sortDirection := "ASC"
	if !isAsc {
		sortDirection = "DESC"
	}

	// 如果用户未指定排序字段，则默认使用 id
	if orderStr == "" {
		orderStr = fmt.Sprintf("ORDER BY id %s", sortDirection)
	} else {
		orderStr = strings.TrimSpace(orderStr)
		if !strings.HasPrefix(strings.ToUpper(orderStr), "ORDER BY") {
			orderStr = "ORDER BY " + orderStr
		}
		orderStr = fmt.Sprintf("%s %s", orderStr, sortDirection)
	}

	countQuery := fmt.Sprintf("SELECT COUNT(1) FROM %s %s", m.table, where)

	var total uint64
	var resp []*SysUser
	err := m.QueryRowNoCacheCtx(ctx, &total, countQuery, args...)
	if err != nil {
		return nil, 0, err
	}
	if total == 0 {
		// 无匹配记录
		return resp, 0, ErrNotFound
	}
	offset := (page - 1) * pageSize
	dataQuery := fmt.Sprintf("SELECT %s FROM %s %s %s LIMIT %d,%d", sysUserRows, m.table, where, orderStr, offset, pageSize)

	err = m.QueryRowsNoCacheCtx(ctx, &resp, dataQuery, args...)
	if err != nil {
		return nil, 0, err
	}

	return resp, total, nil
}

func (m *defaultSysUserModel) SearchNoPage(ctx context.Context, orderStr string, isAsc bool, queryStr string, args ...any) ([]*SysUser, error) {
	// 初始化 WHERE 子句
	where := "WHERE `is_deleted` = 0"
	if queryStr != "" {
		where = fmt.Sprintf("WHERE %s AND `is_deleted` = 0", queryStr)
	}

	// 根据 isAsc 参数确定排序方式
	sortDirection := "ASC"
	if !isAsc {
		sortDirection = "DESC"
	}
	// 如果用户未指定排序字段，则默认使用 id
	if orderStr == "" {
		orderStr = fmt.Sprintf("ORDER BY id %s", sortDirection)
	} else {
		orderStr = strings.TrimSpace(orderStr)
		if !strings.HasPrefix(strings.ToUpper(orderStr), "ORDER BY") {
			orderStr = "ORDER BY " + orderStr
		}
		orderStr = fmt.Sprintf("%s %s", orderStr, sortDirection)
	}
	dataQuery := fmt.Sprintf("SELECT %s FROM %s %s %s", sysUserRows, m.table, where, orderStr)
	var resp []*SysUser
	err := m.QueryRowsNoCacheCtx(ctx, &resp, dataQuery, args...)
	switch err {
	case nil:
		return resp, nil
	case sqlx.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
func (m *defaultSysUserModel) FindOneByAccount(ctx context.Context, account string) (*SysUser, error) {
	kubeOnecSysUserAccountKey := fmt.Sprintf("%s%v", cacheKubeOnecSysUserAccountPrefix, account)
	var resp SysUser
	err := m.QueryRowIndexCtx(ctx, &resp, kubeOnecSysUserAccountKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v any) (i any, e error) {
		query := fmt.Sprintf("select %s from %s where `account` = ? AND `is_deleted` = 0  limit 1", sysUserRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, account); err != nil {
			return nil, err
		}
		return resp.Id, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultSysUserModel) FindOneByEmail(ctx context.Context, email string) (*SysUser, error) {
	kubeOnecSysUserEmailKey := fmt.Sprintf("%s%v", cacheKubeOnecSysUserEmailPrefix, email)
	var resp SysUser
	err := m.QueryRowIndexCtx(ctx, &resp, kubeOnecSysUserEmailKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v any) (i any, e error) {
		query := fmt.Sprintf("select %s from %s where `email` = ? AND `is_deleted` = 0  limit 1", sysUserRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, email); err != nil {
			return nil, err
		}
		return resp.Id, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultSysUserModel) FindOneByMobile(ctx context.Context, mobile string) (*SysUser, error) {
	kubeOnecSysUserMobileKey := fmt.Sprintf("%s%v", cacheKubeOnecSysUserMobilePrefix, mobile)
	var resp SysUser
	err := m.QueryRowIndexCtx(ctx, &resp, kubeOnecSysUserMobileKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v any) (i any, e error) {
		query := fmt.Sprintf("select %s from %s where `mobile` = ? AND `is_deleted` = 0  limit 1", sysUserRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, mobile); err != nil {
			return nil, err
		}
		return resp.Id, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultSysUserModel) FindOneByWorkNumber(ctx context.Context, workNumber string) (*SysUser, error) {
	kubeOnecSysUserWorkNumberKey := fmt.Sprintf("%s%v", cacheKubeOnecSysUserWorkNumberPrefix, workNumber)
	var resp SysUser
	err := m.QueryRowIndexCtx(ctx, &resp, kubeOnecSysUserWorkNumberKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v any) (i any, e error) {
		query := fmt.Sprintf("select %s from %s where `work_number` = ? AND `is_deleted` = 0  limit 1", sysUserRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, workNumber); err != nil {
			return nil, err
		}
		return resp.Id, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultSysUserModel) Insert(ctx context.Context, data *SysUser) (sql.Result, error) {
	kubeOnecSysUserAccountKey := fmt.Sprintf("%s%v", cacheKubeOnecSysUserAccountPrefix, data.Account)
	kubeOnecSysUserEmailKey := fmt.Sprintf("%s%v", cacheKubeOnecSysUserEmailPrefix, data.Email)
	kubeOnecSysUserIdKey := fmt.Sprintf("%s%v", cacheKubeOnecSysUserIdPrefix, data.Id)
	kubeOnecSysUserMobileKey := fmt.Sprintf("%s%v", cacheKubeOnecSysUserMobilePrefix, data.Mobile)
	kubeOnecSysUserWorkNumberKey := fmt.Sprintf("%s%v", cacheKubeOnecSysUserWorkNumberPrefix, data.WorkNumber)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, sysUserRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.UserName, data.Account, data.Password, data.Icon, data.Mobile, data.Email, data.WorkNumber, data.HireDate, data.IsResetPassword, data.IsDisabled, data.IsLeave, data.PositionId, data.OrganizationId, data.LastLoginTime, data.IsDeleted)
	}, kubeOnecSysUserAccountKey, kubeOnecSysUserEmailKey, kubeOnecSysUserIdKey, kubeOnecSysUserMobileKey, kubeOnecSysUserWorkNumberKey)
	return ret, err
}

func (m *defaultSysUserModel) Update(ctx context.Context, newData *SysUser) error {
	data, err := m.FindOne(ctx, newData.Id)
	if err != nil {
		return err
	}

	kubeOnecSysUserAccountKey := fmt.Sprintf("%s%v", cacheKubeOnecSysUserAccountPrefix, data.Account)
	kubeOnecSysUserEmailKey := fmt.Sprintf("%s%v", cacheKubeOnecSysUserEmailPrefix, data.Email)
	kubeOnecSysUserIdKey := fmt.Sprintf("%s%v", cacheKubeOnecSysUserIdPrefix, data.Id)
	kubeOnecSysUserMobileKey := fmt.Sprintf("%s%v", cacheKubeOnecSysUserMobilePrefix, data.Mobile)
	kubeOnecSysUserWorkNumberKey := fmt.Sprintf("%s%v", cacheKubeOnecSysUserWorkNumberPrefix, data.WorkNumber)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, sysUserRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, newData.UserName, newData.Account, newData.Password, newData.Icon, newData.Mobile, newData.Email, newData.WorkNumber, newData.HireDate, newData.IsResetPassword, newData.IsDisabled, newData.IsLeave, newData.PositionId, newData.OrganizationId, newData.LastLoginTime, newData.IsDeleted, newData.Id)
	}, kubeOnecSysUserAccountKey, kubeOnecSysUserEmailKey, kubeOnecSysUserIdKey, kubeOnecSysUserMobileKey, kubeOnecSysUserWorkNumberKey)
	return err
}

func (m *defaultSysUserModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheKubeOnecSysUserIdPrefix, primary)
}

func (m *defaultSysUserModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? AND `is_deleted` = 0 limit 1", sysUserRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultSysUserModel) tableName() string {
	return m.table
}
