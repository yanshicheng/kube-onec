package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ OnecProjectApplicationModel = (*customOnecProjectApplicationModel)(nil)

type (
	// OnecProjectApplicationModel is an interface to be customized, add more methods here,
	// and implement the added methods in customOnecProjectApplicationModel.
	OnecProjectApplicationModel interface {
		onecProjectApplicationModel
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
