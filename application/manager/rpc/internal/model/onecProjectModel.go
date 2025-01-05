package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ OnecProjectModel = (*customOnecProjectModel)(nil)

type (
	// OnecProjectModel is an interface to be customized, add more methods here,
	// and implement the added methods in customOnecProjectModel.
	OnecProjectModel interface {
		onecProjectModel
	}

	customOnecProjectModel struct {
		*defaultOnecProjectModel
	}
)

// NewOnecProjectModel returns a model for the database table.
func NewOnecProjectModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) OnecProjectModel {
	return &customOnecProjectModel{
		defaultOnecProjectModel: newOnecProjectModel(conn, c, opts...),
	}
}
