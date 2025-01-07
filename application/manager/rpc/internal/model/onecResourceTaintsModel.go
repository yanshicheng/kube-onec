package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ OnecResourceTaintsModel = (*customOnecResourceTaintsModel)(nil)

type (
	// OnecResourceTaintsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customOnecResourceTaintsModel.
	OnecResourceTaintsModel interface {
		onecResourceTaintsModel
	}

	customOnecResourceTaintsModel struct {
		*defaultOnecResourceTaintsModel
	}
)

// NewOnecResourceTaintsModel returns a model for the database table.
func NewOnecResourceTaintsModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) OnecResourceTaintsModel {
	return &customOnecResourceTaintsModel{
		defaultOnecResourceTaintsModel: newOnecResourceTaintsModel(conn, c, opts...),
	}
}
