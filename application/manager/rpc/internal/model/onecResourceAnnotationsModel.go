package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ OnecResourceAnnotationsModel = (*customOnecResourceAnnotationsModel)(nil)

type (
	// OnecResourceAnnotationsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customOnecResourceAnnotationsModel.
	OnecResourceAnnotationsModel interface {
		onecResourceAnnotationsModel
	}

	customOnecResourceAnnotationsModel struct {
		*defaultOnecResourceAnnotationsModel
	}
)

// NewOnecResourceAnnotationsModel returns a model for the database table.
func NewOnecResourceAnnotationsModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) OnecResourceAnnotationsModel {
	return &customOnecResourceAnnotationsModel{
		defaultOnecResourceAnnotationsModel: newOnecResourceAnnotationsModel(conn, c, opts...),
	}
}
