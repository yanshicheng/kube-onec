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
	SyncClusterInfoErr        = errorx.New(102011, "集群信息同步失败!")
	DictItemNotExistErr       = errorx.New(102012, "字典项不存在!")
	GetClusterIdErr           = errorx.New(102013, "获取集群ID失败!")
	SyncNodeInfoErr           = errorx.New(102009, "节点数据同步失败!")
	SyncNodeLabelErr          = errorx.New(102014, "节点标签数据同步失败!")
	SyncNodeAnnotationsErr    = errorx.New(102015, "节点注解数据同步失败!")
	SyncNodeTaintErr          = errorx.New(102016, "节点污点数据同步失败!")
	RemoveNodeLabelErr        = errorx.New(102017, "节点标签删除失败!")
	SearchNodeLabelDBErr      = errorx.New(102018, "节点标签查询失败!")
	DeleteNodeLabelDBErr      = errorx.New(102019, "节点标签删除失败!")
	SyncNodeAnnotationErr     = errorx.New(102020, "节点注解数据同步失败!")
	SearchNodeAnnotationDBErr = errorx.New(102021, "节点注解查询失败!")
	RemoveNodeAnnotationErr   = errorx.New(102022, "节点注解删除失败!")
	DeleteNodeAnnotationDBErr = errorx.New(102023, "节点注解删除失败!")
	AddNodeTaintErr           = errorx.New(102024, "节点污点数据添加失败!")
	SyncNodeTaintDBErr        = errorx.New(102025, "节点污点数据同步失败!")
	RemoveNodeTaintErr        = errorx.New(102026, "节点污点数据删除失败!")
	SearchNodeTaintDBErr      = errorx.New(102027, "节点污点数据查询失败!")
	DeleteNodeTaintDBErr      = errorx.New(102028, "节点污点数据删除失败!")
	ForbidNodeScheduleErr     = errorx.New(102029, "节点禁止调度失败!")
	EnableNodeScheduleErr     = errorx.New(102030, "节点允许调度失败!")
	NodeNotExistErr           = errorx.New(102031, "未找到节点信息!")
	QueryNodeErr              = errorx.New(102032, "查询节点信息失败!")
	//ClusterUuid 不能为空
	ClusterUuidEmptyErr = errorx.New(102033, "ClusterUuid 不能为空!")
	UpdateNodeInfoErr   = errorx.New(102034, "更新节点信息失败!")
	// 信息查询不到，请先同步信息
	NodeInfoNotExistErr = errorx.New(102035, "节点信息查询不到，请先同步信息!")
	// 节点ID 不能 为空

	NodeIdEmptyErr  = errorx.New(102036, "节点ID 不能 为空!")
	EvictNodePodErr = errorx.New(102037, "驱逐节点Pod失败!")

	// 项目相关 1021**
	CreateProjectErr = errorx.New(102101, "项目创建失败!")
	DeleteProjectErr = errorx.New(102102, "项目删除失败!")
	UpdateProjectErr = errorx.New(102103, "项目更新失败!")
	GetProjectErr    = errorx.New(102104, "项目查询失败!")
)
