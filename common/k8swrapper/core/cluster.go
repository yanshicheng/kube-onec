package core

import "time"

// ClusterInterface 定义集群模块的接口
type ClusterInterface interface {
	GetClusterInfo() (*ClusterInfo, error)
}

type ClusterInfo struct {
	// 集群版本
	Version string
	// 集群Commit
	Commit string
	// 集群platform
	Platform string
	// 构建时间
	BuildTime time.Time
	// 创建时间
	CreateTime time.Time
}
