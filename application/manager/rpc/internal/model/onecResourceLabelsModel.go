package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ OnecResourceLabelsModel = (*customOnecResourceLabelsModel)(nil)

type (
	// OnecResourceLabelsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customOnecResourceLabelsModel.
	OnecResourceLabelsModel interface {
		onecResourceLabelsModel
	}

	customOnecResourceLabelsModel struct {
		*defaultOnecResourceLabelsModel
	}
)

// NewOnecResourceLabelsModel returns a model for the database table.
func NewOnecResourceLabelsModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) OnecResourceLabelsModel {
	return &customOnecResourceLabelsModel{
		defaultOnecResourceLabelsModel: newOnecResourceLabelsModel(conn, c, opts...),
	}
}
