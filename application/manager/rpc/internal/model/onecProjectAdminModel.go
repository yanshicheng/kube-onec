package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ OnecProjectAdminModel = (*customOnecProjectAdminModel)(nil)

type (
	// OnecProjectAdminModel is an interface to be customized, add more methods here,
	// and implement the added methods in customOnecProjectAdminModel.
	OnecProjectAdminModel interface {
		onecProjectAdminModel
	}

	customOnecProjectAdminModel struct {
		*defaultOnecProjectAdminModel
	}
)

// NewOnecProjectAdminModel returns a model for the database table.
func NewOnecProjectAdminModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) OnecProjectAdminModel {
	return &customOnecProjectAdminModel{
		defaultOnecProjectAdminModel: newOnecProjectAdminModel(conn, c, opts...),
	}
}
