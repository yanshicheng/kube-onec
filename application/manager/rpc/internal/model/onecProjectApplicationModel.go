package model

import (
	"context"
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ OnecProjectApplicationModel = (*customOnecProjectApplicationModel)(nil)

type (
	// OnecProjectApplicationModel is an interface to be customized, add more methods here,
	// and implement the added methods in customOnecProjectApplicationModel.
	OnecProjectApplicationModel interface {
		onecProjectApplicationModel
		FindProjectResourceQuotaByCluster(ctx context.Context, projectId uint64, clusterUuid string) (*OnecProjectApplicationQuota, error)
	}

	customOnecProjectApplicationModel struct {
		*defaultOnecProjectApplicationModel
	}
)

// NewOnecProjectApplicationModel returns a model for the database table.
func NewOnecProjectApplicationModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) OnecProjectApplicationModel {
	return &customOnecProjectApplicationModel{
		defaultOnecProjectApplicationModel: newOnecProjectApplicationModel(conn, c, opts...),
	}
}

type OnecProjectApplicationQuota struct {
	CpuTotal       float64 `json:"cpuTotal"`
	MemoryTotal    float64 `json:"memoryTotal"`
	PodTotal       int64   `json:"podTotal"`
	StorageTotal   int64   `json:"storageTotal"`
	ConfigMapTotal int64   `json:"configMapTotal"`
	PvcTotal       int64   `json:"pvcTotal"`
	SecretTotal    int64   `json:"secretTotal"`
	ServiceTotal   int64   `json:"serviceTotal"`
	NodePortTotal  int64   `json:"nodePortTotal"`
}

func (m *defaultOnecProjectApplicationModel) FindProjectResourceQuotaByCluster(ctx context.Context, projectId uint64, clusterUuid string) (*OnecProjectApplicationQuota, error) {
	var quota OnecProjectApplicationQuota

	// 定义 SQL 查询语句
	query := fmt.Sprintf(`
SELECT 
    COALESCE(SUM(cpu_limit), 0) AS cpuTotal,
    COALESCE(SUM(memory_limit), 0) AS memoryTotal,
    COALESCE(SUM(pod_limit), 0) AS podTotal,
    COALESCE(SUM(storage_limit), 0) AS storageTotal,
    COALESCE(SUM(configmap_limit), 0) AS configMapTotal,
    COALESCE(SUM(pvc_limit), 0) AS pvcTotal,
    COALESCE(SUM(secret_limit), 0) AS secretTotal,
    COALESCE(SUM(service_limit), 0) AS serviceTotal,
    COALESCE(SUM(nodeport_limit), 0) AS nodePortTotal
FROM %s
WHERE project_id = ? AND cluster_uuid = ? AND is_deleted = 0
	`, m.table)

	// 执行查询，将结果绑定到 OnecProjectApplicationQuota 对象
	err := m.QueryRowNoCacheCtx(ctx, &quota, query, projectId, clusterUuid)
	switch {
	case err == nil:
		return &quota, nil
	case errors.Is(err, sqlx.ErrNotFound):
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
