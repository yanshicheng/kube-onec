package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ OnecClusterConnInfoModel = (*customOnecClusterConnInfoModel)(nil)

type (
	// OnecClusterConnInfoModel is an interface to be customized, add more methods here,
	// and implement the added methods in customOnecClusterConnInfoModel.
	OnecClusterConnInfoModel interface {
		onecClusterConnInfoModel
	}

	customOnecClusterConnInfoModel struct {
		*defaultOnecClusterConnInfoModel
	}
)

// NewOnecClusterConnInfoModel returns a model for the database table.
func NewOnecClusterConnInfoModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) OnecClusterConnInfoModel {
	return &customOnecClusterConnInfoModel{
		defaultOnecClusterConnInfoModel: newOnecClusterConnInfoModel(conn, c, opts...),
	}
}
