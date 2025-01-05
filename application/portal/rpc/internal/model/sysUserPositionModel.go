package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ SysUserPositionModel = (*customSysUserPositionModel)(nil)

type (
	// SysUserPositionModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSysUserPositionModel.
	SysUserPositionModel interface {
		sysUserPositionModel
	}

	customSysUserPositionModel struct {
		*defaultSysUserPositionModel
	}
)

// NewSysUserPositionModel returns a model for the database table.
func NewSysUserPositionModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) SysUserPositionModel {
	return &customSysUserPositionModel{
		defaultSysUserPositionModel: newSysUserPositionModel(conn, c, opts...),
	}
}

// 通过 position_id 获取用户列表
