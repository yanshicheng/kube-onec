package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ SysPositionModel = (*customSysPositionModel)(nil)

type (
	// SysPositionModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSysPositionModel.
	SysPositionModel interface {
		sysPositionModel
	}

	customSysPositionModel struct {
		*defaultSysPositionModel
	}
)

// NewSysPositionModel returns a model for the database table.
func NewSysPositionModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) SysPositionModel {
	return &customSysPositionModel{
		defaultSysPositionModel: newSysPositionModel(conn, c, opts...),
	}
}
