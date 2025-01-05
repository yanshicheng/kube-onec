package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ OnecClusterModel = (*customOnecClusterModel)(nil)

type (
	// OnecClusterModel is an interface to be customized, add more methods here,
	// and implement the added methods in customOnecClusterModel.
	OnecClusterModel interface {
		onecClusterModel
	}

	customOnecClusterModel struct {
		*defaultOnecClusterModel
	}
)

// NewOnecClusterModel returns a model for the database table.
func NewOnecClusterModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) OnecClusterModel {
	return &customOnecClusterModel{
		defaultOnecClusterModel: newOnecClusterModel(conn, c, opts...),
	}
}
