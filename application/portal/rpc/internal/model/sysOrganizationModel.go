package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ SysOrganizationModel = (*customSysOrganizationModel)(nil)

type (
	// SysOrganizationModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSysOrganizationModel.
	SysOrganizationModel interface {
		sysOrganizationModel
	}

	customSysOrganizationModel struct {
		*defaultSysOrganizationModel
	}
)

// NewSysOrganizationModel returns a model for the database table.
func NewSysOrganizationModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) SysOrganizationModel {
	return &customSysOrganizationModel{
		defaultSysOrganizationModel: newSysOrganizationModel(conn, c, opts...),
	}
}
