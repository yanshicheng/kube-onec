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
	onecProjectQuotaFieldNames          = builder.RawFieldNames(&OnecProjectQuota{})
	onecProjectQuotaRows                = strings.Join(onecProjectQuotaFieldNames, ",")
	onecProjectQuotaRowsExpectAutoSet   = strings.Join(stringx.Remove(onecProjectQuotaFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	onecProjectQuotaRowsWithPlaceHolder = strings.Join(stringx.Remove(onecProjectQuotaFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheKubeOnecOnecProjectQuotaIdPrefix                   = "cache:kubeOnec:onecProjectQuota:id:"
	cacheKubeOnecOnecProjectQuotaClusterUuidProjectIdPrefix = "cache:kubeOnec:onecProjectQuota:clusterUuid:projectId:"
)

type (
	onecProjectQuotaModel interface {
		Insert(ctx context.Context, data *OnecProjectQuota) (sql.Result, error)

		FindOne(ctx context.Context, id uint64) (*OnecProjectQuota, error)
		Search(ctx context.Context, orderStr string, isAsc bool, page, pageSize uint64, queryStr string, args ...any) ([]*OnecProjectQuota, uint64, error)
		SearchNoPage(ctx context.Context, orderStr string, isAsc bool, queryStr string, args ...any) ([]*OnecProjectQuota, error)
		FindOneByClusterUuidProjectId(ctx context.Context, clusterUuid string, projectId uint64) (*OnecProjectQuota, error)
		Update(ctx context.Context, data *OnecProjectQuota) error
		Delete(ctx context.Context, id uint64) error
		DeleteSoft(ctx context.Context, id uint64) error
		TransCtx(ctx context.Context, fn func(context.Context, sqlx.Session) error) error
		TransOnSql(ctx context.Context, session sqlx.Session, id uint64, sqlStr string, args ...any) (sql.Result, error)
		ExecSql(ctx context.Context, id uint64, sqlStr string, args ...any) (sql.Result, error)
	}

	defaultOnecProjectQuotaModel struct {
		sqlc.CachedConn
		table string
	}

	OnecProjectQuota struct {
		Id               uint64    `db:"id"`                // 主键，自增 ID
		ClusterUuid      string    `db:"cluster_uuid"`      // 关联的集群 ID
		ProjectId        uint64    `db:"project_id"`        // 关联的项目 ID
		CpuQuota         int64     `db:"cpu_quota"`         // CPU 分配配额（单位：核）
		CpuOvercommit    float64   `db:"cpu_overcommit"`    // CPU 超配比（如 1.5 表示允许超配 50%）
		CpuLimit         float64   `db:"cpu_limit"`         // CPU 上限值（单位：核）
		CpuUsed          float64   `db:"cpu_used"`          // 已使用的 CPU 资源（单位：核）
		MemoryQuota      float64   `db:"memory_quota"`      // 内存分配配额（单位：GiB）
		MemoryOvercommit float64   `db:"memory_overcommit"` // 内存超配比（如 1.2 表示允许超配 20%）
		MemoryLimit      float64   `db:"memory_limit"`      // 内存上限值（单位：GiB）
		MemoryUsed       float64   `db:"memory_used"`       // 已使用的内存资源（单位：GiB）
		StorageLimit     int64     `db:"storage_limit"`     // 项目可使用的存储总量（单位：GiB）
		StorageUsed      int64     `db:"storage_used"`      // 已使用的存储资源（单位：GiB）
		ConfigmapLimit   int64     `db:"configmap_limit"`   // 项目允许创建的 ConfigMap 数量上限
		ConfigmapUsed    int64     `db:"configmap_used"`    // 已使用的 ConfigMap 资源数量
		SecretUsed       int64     `db:"secret_used"`       // 已使用的 Secret 资源数量
		SecretLimit      int64     `db:"secret_limit"`      // 项目允许创建的 Secret 数量上限
		PvcLimit         int64     `db:"pvc_limit"`         // 项目允许创建的 PVC（PersistentVolumeClaim）数量上限
		PvcUsed          int64     `db:"pvc_used"`          // 已使用的 PVC 资源数量
		PodLimit         int64     `db:"pod_limit"`         // 项目允许创建的 Pod 数量上限
		PodUsed          int64     `db:"pod_used"`          // 已使用的 Pod 资源数量
		NodeportLimit    int64     `db:"nodeport_limit"`    // 项目允许使用的 NodePort 数量上限
		NodeportUsed     int64     `db:"nodeport_used"`     // 已使用的 NodePort 资源数量
		Status           string    `db:"status"`            // 项目状态（如 `Active`、`Disabled`、`Archived`）
		CreatedBy        string    `db:"created_by"`        // 记录创建人
		UpdatedBy        string    `db:"updated_by"`        // 记录更新人
		CreatedAt        time.Time `db:"created_at"`        // 记录创建时间
		UpdatedAt        time.Time `db:"updated_at"`        // 记录更新时间
		IsDeleted        int64     `db:"is_deleted"`        // 是否删除
		ServiceLimit     int64     `db:"service_limit"`     // 项目允许创建的 Service 数量上限
		ServiceUsed      int64     `db:"service_used"`      // 已使用的 Service 资源数量
	}
)

func newOnecProjectQuotaModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultOnecProjectQuotaModel {
	return &defaultOnecProjectQuotaModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      "`onec_project_quota`",
	}
}

func (m *defaultOnecProjectQuotaModel) Delete(ctx context.Context, id uint64) error {
	data, err := m.FindOne(ctx, id)
	if err != nil {
		return err
	}

	kubeOnecOnecProjectQuotaClusterUuidProjectIdKey := fmt.Sprintf("%s%v:%v", cacheKubeOnecOnecProjectQuotaClusterUuidProjectIdPrefix, data.ClusterUuid, data.ProjectId)
	kubeOnecOnecProjectQuotaIdKey := fmt.Sprintf("%s%v", cacheKubeOnecOnecProjectQuotaIdPrefix, id)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, kubeOnecOnecProjectQuotaClusterUuidProjectIdKey, kubeOnecOnecProjectQuotaIdKey)
	return err
}

func (m *defaultOnecProjectQuotaModel) DeleteSoft(ctx context.Context, id uint64) error {
	data, err := m.FindOne(ctx, id)
	if err != nil {
		return err
	}
	// 如果记录已软删除，无需再次删除
	if data.IsDeleted == 1 {
		return nil
	}
	kubeOnecOnecProjectQuotaClusterUuidProjectIdKey := fmt.Sprintf("%s%v:%v", cacheKubeOnecOnecProjectQuotaClusterUuidProjectIdPrefix, data.ClusterUuid, data.ProjectId)
	kubeOnecOnecProjectQuotaIdKey := fmt.Sprintf("%s%v", cacheKubeOnecOnecProjectQuotaIdPrefix, id)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set `is_deleted` = 1 where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, kubeOnecOnecProjectQuotaClusterUuidProjectIdKey, kubeOnecOnecProjectQuotaIdKey)
	return err
}

func (m *defaultOnecProjectQuotaModel) TransCtx(ctx context.Context, fn func(context.Context, sqlx.Session) error) error {
	return m.TransactCtx(ctx, func(ctx context.Context, session sqlx.Session) error {
		return fn(ctx, session)
	})
}

func (m *defaultOnecProjectQuotaModel) TransOnSql(ctx context.Context, session sqlx.Session, id uint64, sqlStr string, args ...any) (sql.Result, error) {
	query := strings.ReplaceAll(sqlStr, "{table}", m.table)
	// 如果 id != 0 并且启用了缓存逻辑
	if !isZeroValue(id) {
		// 查询数据（如果需要，确保数据存在）
		data, err := m.FindOne(ctx, id)
		if err != nil {
			return nil, err
		}

		// 缓存相关处理

		kubeOnecOnecProjectQuotaClusterUuidProjectIdKey := fmt.Sprintf("%s%v:%v", cacheKubeOnecOnecProjectQuotaClusterUuidProjectIdPrefix, data.ClusterUuid, data.ProjectId)
		kubeOnecOnecProjectQuotaIdKey := fmt.Sprintf("%s%v", cacheKubeOnecOnecProjectQuotaIdPrefix, id) // 处理缓存逻辑，例如删除或更新缓存
		// 执行带缓存处理的 SQL 操作
		return m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (sql.Result, error) {
			return session.ExecCtx(ctx, query, args...)
		}, kubeOnecOnecProjectQuotaClusterUuidProjectIdKey, kubeOnecOnecProjectQuotaIdKey) // 传递缓存相关的键值

	}

	// 如果 id == 0 或不需要缓存，直接执行 SQL
	return m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (sql.Result, error) {
		return session.ExecCtx(ctx, query, args...)
	})
}

func (m *defaultOnecProjectQuotaModel) ExecSql(ctx context.Context, id uint64, sqlStr string, args ...any) (sql.Result, error) {
	// 如果 id != 0 并且启用了缓存逻辑
	query := strings.ReplaceAll(sqlStr, "{table}", m.table)
	if !isZeroValue(id) {
		// 缓存相关处理

		// 查询数据（如果需要，确保数据存在）
		data, err := m.FindOne(ctx, id)
		if err != nil {
			return nil, err
		}

		kubeOnecOnecProjectQuotaClusterUuidProjectIdKey := fmt.Sprintf("%s%v:%v", cacheKubeOnecOnecProjectQuotaClusterUuidProjectIdPrefix, data.ClusterUuid, data.ProjectId)
		kubeOnecOnecProjectQuotaIdKey := fmt.Sprintf("%s%v", cacheKubeOnecOnecProjectQuotaIdPrefix, id) // 处理缓存逻辑，例如删除或更新缓存
		// 执行带缓存处理的 SQL 操作
		return m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (sql.Result, error) {
			return conn.ExecCtx(ctx, query, args...)
		}, kubeOnecOnecProjectQuotaClusterUuidProjectIdKey, kubeOnecOnecProjectQuotaIdKey) // 传递缓存相关的键值

	}

	// 如果 id == 0 或不需要缓存，直接执行 SQL
	return m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (sql.Result, error) {
		return conn.ExecCtx(ctx, query, args...)
	})
}
func (m *defaultOnecProjectQuotaModel) FindOne(ctx context.Context, id uint64) (*OnecProjectQuota, error) {
	kubeOnecOnecProjectQuotaIdKey := fmt.Sprintf("%s%v", cacheKubeOnecOnecProjectQuotaIdPrefix, id)
	var resp OnecProjectQuota
	err := m.QueryRowCtx(ctx, &resp, kubeOnecOnecProjectQuotaIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? AND `is_deleted` = 0 limit 1", onecProjectQuotaRows, m.table)
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

func (m *defaultOnecProjectQuotaModel) Search(ctx context.Context, orderStr string, isAsc bool, page, pageSize uint64, queryStr string, args ...any) ([]*OnecProjectQuota, uint64, error) {
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
	var resp []*OnecProjectQuota
	err := m.QueryRowNoCacheCtx(ctx, &total, countQuery, args...)
	if err != nil {
		return nil, 0, err
	}
	if total == 0 {
		// 无匹配记录
		return resp, 0, ErrNotFound
	}
	offset := (page - 1) * pageSize
	dataQuery := fmt.Sprintf("SELECT %s FROM %s %s %s LIMIT %d,%d", onecProjectQuotaRows, m.table, where, orderStr, offset, pageSize)

	err = m.QueryRowsNoCacheCtx(ctx, &resp, dataQuery, args...)
	if err != nil {
		return nil, 0, err
	}

	return resp, total, nil
}

func (m *defaultOnecProjectQuotaModel) SearchNoPage(ctx context.Context, orderStr string, isAsc bool, queryStr string, args ...any) ([]*OnecProjectQuota, error) {
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
	dataQuery := fmt.Sprintf("SELECT %s FROM %s %s %s", onecProjectQuotaRows, m.table, where, orderStr)
	var resp []*OnecProjectQuota
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
func (m *defaultOnecProjectQuotaModel) FindOneByClusterUuidProjectId(ctx context.Context, clusterUuid string, projectId uint64) (*OnecProjectQuota, error) {
	kubeOnecOnecProjectQuotaClusterUuidProjectIdKey := fmt.Sprintf("%s%v:%v", cacheKubeOnecOnecProjectQuotaClusterUuidProjectIdPrefix, clusterUuid, projectId)
	var resp OnecProjectQuota
	err := m.QueryRowIndexCtx(ctx, &resp, kubeOnecOnecProjectQuotaClusterUuidProjectIdKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v any) (i any, e error) {
		query := fmt.Sprintf("select %s from %s where `cluster_uuid` = ? and `project_id` = ? AND `is_deleted` = 0  limit 1", onecProjectQuotaRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, clusterUuid, projectId); err != nil {
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

func (m *defaultOnecProjectQuotaModel) Insert(ctx context.Context, data *OnecProjectQuota) (sql.Result, error) {
	kubeOnecOnecProjectQuotaClusterUuidProjectIdKey := fmt.Sprintf("%s%v:%v", cacheKubeOnecOnecProjectQuotaClusterUuidProjectIdPrefix, data.ClusterUuid, data.ProjectId)
	kubeOnecOnecProjectQuotaIdKey := fmt.Sprintf("%s%v", cacheKubeOnecOnecProjectQuotaIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, onecProjectQuotaRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.ClusterUuid, data.ProjectId, data.CpuQuota, data.CpuOvercommit, data.CpuLimit, data.CpuUsed, data.MemoryQuota, data.MemoryOvercommit, data.MemoryLimit, data.MemoryUsed, data.StorageLimit, data.StorageUsed, data.ConfigmapLimit, data.ConfigmapUsed, data.SecretUsed, data.SecretLimit, data.PvcLimit, data.PvcUsed, data.PodLimit, data.PodUsed, data.NodeportLimit, data.NodeportUsed, data.Status, data.CreatedBy, data.UpdatedBy, data.IsDeleted, data.ServiceLimit, data.ServiceUsed)
	}, kubeOnecOnecProjectQuotaClusterUuidProjectIdKey, kubeOnecOnecProjectQuotaIdKey)
	return ret, err
}

func (m *defaultOnecProjectQuotaModel) Update(ctx context.Context, newData *OnecProjectQuota) error {
	data, err := m.FindOne(ctx, newData.Id)
	if err != nil {
		return err
	}

	kubeOnecOnecProjectQuotaClusterUuidProjectIdKey := fmt.Sprintf("%s%v:%v", cacheKubeOnecOnecProjectQuotaClusterUuidProjectIdPrefix, data.ClusterUuid, data.ProjectId)
	kubeOnecOnecProjectQuotaIdKey := fmt.Sprintf("%s%v", cacheKubeOnecOnecProjectQuotaIdPrefix, data.Id)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, onecProjectQuotaRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, newData.ClusterUuid, newData.ProjectId, newData.CpuQuota, newData.CpuOvercommit, newData.CpuLimit, newData.CpuUsed, newData.MemoryQuota, newData.MemoryOvercommit, newData.MemoryLimit, newData.MemoryUsed, newData.StorageLimit, newData.StorageUsed, newData.ConfigmapLimit, newData.ConfigmapUsed, newData.SecretUsed, newData.SecretLimit, newData.PvcLimit, newData.PvcUsed, newData.PodLimit, newData.PodUsed, newData.NodeportLimit, newData.NodeportUsed, newData.Status, newData.CreatedBy, newData.UpdatedBy, newData.IsDeleted, newData.ServiceLimit, newData.ServiceUsed, newData.Id)
	}, kubeOnecOnecProjectQuotaClusterUuidProjectIdKey, kubeOnecOnecProjectQuotaIdKey)
	return err
}

func (m *defaultOnecProjectQuotaModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheKubeOnecOnecProjectQuotaIdPrefix, primary)
}

func (m *defaultOnecProjectQuotaModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? AND `is_deleted` = 0 limit 1", onecProjectQuotaRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultOnecProjectQuotaModel) tableName() string {
	return m.table
}
