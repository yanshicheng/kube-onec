package code

import "github.com/yanshicheng/kube-onec/common/handler/errorx"

var (
	// 集群相关错误 102
	UnsupportedConnTypeErr = errorx.New(102001, "不支持的连接类型!")
	SyncClusterErr         = errorx.New(102002, "集群同步失败!")
	AddClusterFailedErr    = errorx.New(102003, "集群添加失败!")
	// 获取集群client失败
	GetClusterClientErr = errorx.New(102004, "获取集群client失败!")
	ClusterConnectErr   = errorx.New(102005, "集群连接状态异常，请检查地址和Token信息!")
	GetClusterInfoErr   = errorx.New(102006, "获取集群信息失败!")
	GetNodeInfoErr      = errorx.New(102007, "获取节点信息失败!")
	AddNodeInfoErr      = errorx.New(102008, "节点信息添加失败!")
	// 获取集群客户端异常!
	UpdateClusterInfoErr = errorx.New(102010, "更新集群信息失败!")
	// 集群信息同步失败
	SyncClusterInfoErr = errorx.New(102011, "集群信息同步失败!")
)
