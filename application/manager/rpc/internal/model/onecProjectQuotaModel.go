package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ OnecProjectQuotaModel = (*customOnecProjectQuotaModel)(nil)

type (
	// OnecProjectQuotaModel is an interface to be customized, add more methods here,
	// and implement the added methods in customOnecProjectQuotaModel.
	OnecProjectQuotaModel interface {
		onecProjectQuotaModel
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
