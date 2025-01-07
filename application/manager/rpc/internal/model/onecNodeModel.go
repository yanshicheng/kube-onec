package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ OnecNodeModel = (*customOnecNodeModel)(nil)

type (
	// OnecNodeModel is an interface to be customized, add more methods here,
	// and implement the added methods in customOnecNodeModel.
	OnecNodeModel interface {
		onecNodeModel
		FindOneClusterTotalInfo(ctx context.Context, clusterId string) (*clusterTotalInfo, error)
	}

	customOnecNodeModel struct {
		*defaultOnecNodeModel
	}
)

// NewOnecNodeModel returns a model for the database table.
func NewOnecNodeModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) OnecNodeModel {
	return &customOnecNodeModel{
		defaultOnecNodeModel: newOnecNodeModel(conn, c, opts...),
	}
}

type clusterTotalInfo struct {
	TotalNode   int64   `json:"totalNode"`
	TotalCpu    int64   `json:"totalCpu"`
	TotalMemory float64 `json:"totalMemory"`
	TotalPods   int64   `json:"totalPods"`
}

func (m *defaultOnecNodeModel) FindOneClusterTotalInfo(ctx context.Context, cluster_uuid string) (*clusterTotalInfo, error) {
	// 构建缓存键
	var resp clusterTotalInfo
	// 定义 SQL 查询语句
	query := fmt.Sprintf(`
		SELECT 
			COUNT(*) AS totalNode,
			SUM(cpu) AS totalCpu,
			SUM(memory) AS totalMemory,
			SUM(max_pods) AS totalPods
		FROM 
			%s
		WHERE 
			cluster_uuid = ? AND is_deleted = 0
	`, m.table)

	// 执行查询，直接从数据库获取数据，不使用缓存
	//m.ExecCtx()
	err := m.QueryRowNoCacheCtx(ctx, &resp, query, cluster_uuid)
	if err != nil {
		if err == sqlc.ErrNotFound {
			return nil, fmt.Errorf("cluster_uuid %s not found", cluster_uuid)
		}
		return nil, err
	}
	return &resp, nil
}
