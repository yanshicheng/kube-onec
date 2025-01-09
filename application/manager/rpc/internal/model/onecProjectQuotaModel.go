package model

import (
	"context"
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ OnecProjectQuotaModel = (*customOnecProjectQuotaModel)(nil)

type (
	// OnecProjectQuotaModel is an interface to be customized, add more methods here,
	// and implement the added methods in customOnecProjectQuotaModel.
	OnecProjectQuotaModel interface {
		onecProjectQuotaModel
		FindAllByProjectQuotaTotal(ctx context.Context, projectId uint64, clusterUuid string) (*ProjectQuota, error)
		FindClusterQuotasByUuid(ctx context.Context, clusterUuid string) (*ClusterQuota, error)
	}

	customOnecProjectQuotaModel struct {
		*defaultOnecProjectQuotaModel
	}
)

// NewOnecProjectQuotaModel returns a model for the database table.
func NewOnecProjectQuotaModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) OnecProjectQuotaModel {
	return &customOnecProjectQuotaModel{
		defaultOnecProjectQuotaModel: newOnecProjectQuotaModel(conn, c, opts...),
	}
}

type ProjectQuota struct {
	CpuLimit       float64 `json:"cpuLimit"`
	MemoryLimit    float64 `json:"memoryLimit"`
	StorageLimit   int64   `json:"storageLimit"`
	PodLimit       int64   `json:"podLimit"`
	NodeportLimit  int64   `json:"nodeportLimit"`
	PvcLimit       int64   `json:"pvcLimit"`
	ConfigmapLimit int64   `json:"configmapLimit"`
	SecretLimit    int64   `json:"secretLimit"`
}

func (m *defaultOnecProjectQuotaModel) FindAllByProjectQuotaTotal(ctx context.Context, projectId uint64, clusterUuid string) (*ProjectQuota, error) {
	// 定义返回的 ProjectQuota 对象
	var quota ProjectQuota
	table := "onec_project_application"
	// 定义 SQL 查询语句
	query := fmt.Sprintf(`
SELECT 
    COALESCE(SUM(cpu_limit), 0) AS cpuLimit,
    COALESCE(SUM(memory_limit), 0) AS memoryLimit,
    COALESCE(SUM(storage_limit), 0) AS storageLimit,
    COALESCE(SUM(pod_limit), 0) AS podLimit,
    COALESCE(SUM(nodeport_limit), 0) AS nodeportLimit,
    COALESCE(SUM(pvc_limit), 0) AS pvcLimit,
    COALESCE(SUM(configmap_limit), 0) AS configmapLimit,
    COALESCE(SUM(secret_limit), 0) AS secretLimit
FROM %s
WHERE project_id = ? AND cluster_uuid = ? AND is_deleted = 0
	`, table)

	// 执行查询，将结果绑定到 ProjectQuota 对象 	err := m.QueryRowNoCacheCtx(ctx, &resp, query, cluster_uuid)
	err := m.QueryRowNoCacheCtx(ctx, &quota, query, projectId, clusterUuid)
	switch {
	case err == nil:
		return &quota, nil
	case errors.Is(err, sqlc.ErrNotFound):
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

// 定义返回的 ClusterQuota 对象
type ClusterQuota struct {
	CPUQuota    int64   `json:"cpuQuota"`
	MemoryQuota float64 `json:"memoryQuota"`
	PodLimit    int64   `json:"podLimit"`
}

func (m *defaultOnecProjectQuotaModel) FindClusterQuotasByUuid(ctx context.Context, clusterUuid string) (*ClusterQuota, error) {

	var quota ClusterQuota
	// 定义 SQL 查询语句
	query := fmt.Sprintf(`
SELECT 
    COALESCE(SUM(cpu_quota), 0) AS cpuQuota,
    COALESCE(SUM(memory_quota), 0) AS memoryQuota,
    COALESCE(SUM(pod_limit), 0) AS podLimit
FROM %s
WHERE cluster_uuid = ? AND is_deleted = 0
	`, m.table)

	// 执行查询，将结果绑定到 ClusterQuota 对象
	err := m.QueryRowNoCacheCtx(ctx, &quota, query, clusterUuid)
	switch {
	case err == nil:
		return &quota, nil
	case errors.Is(err, sqlc.ErrNotFound):
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
