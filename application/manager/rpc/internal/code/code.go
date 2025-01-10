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

	NodeIdEmptyErr              = errorx.New(102036, "节点ID 不能 为空!")
	EvictNodePodErr             = errorx.New(102037, "驱逐节点Pod失败!")
	ClusterResourceNotEnoughErr = errorx.New(102038, "集群资源不足!")

	// 项目相关 1021**
	CreateProjectErr                       = errorx.New(102101, "项目创建失败!")
	DeleteProjectErr                       = errorx.New(102102, "项目删除失败!")
	UpdateProjectErr                       = errorx.New(102103, "项目更新失败!")
	GetProjectErr                          = errorx.New(102104, "项目查询失败!")
	DefaultProjectNotAllowCreateProjectErr = errorx.New(102105, "默认项目不允许创建项目资源!")
	ProjectResourceNotAllowedErr           = errorx.New(102106, "项目资源不允许分配集群资源!")
	AddProjectQuotaErr                     = errorx.New(102107, "项目资源添加失败!")
	// 项目集群资源分配失败

	ProjectQuotaAllocationErr = errorx.New(102108, "项目资源分配失败!")
	ProjectQuotaExistErr      = errorx.New(102109, "项目资源已存在!")
	GetProjectQuotaErr        = errorx.New(102110, "项目资源查询失败!")
	//传入的资源小于已经分配的总资源
	ProjectQuotaNotEnoughErr = errorx.New(102111, "传入的资源小于已经分配的总资源!")
	UpdateProjectQuotaErr    = errorx.New(102112, "项目资源更新失败!")
	UpdateClusterErr         = errorx.New(102113, "集群更新失败!")
	IdentifierErr            = errorx.New(102114, "标识符不合法，只允许小写英文和‘-’ 并且只能英文字母开头!")
	// 项目资源不足，无法创建应用
	ProjectQuotaNotEnoughCreateAppErr = errorx.New(102115, "项目资源不足，无法创建应用!")
	NamespaceExistErr                 = errorx.New(102116, "命名空间已存在或查询异常!")
	CreateNamespaceErr                = errorx.New(102117, "命名空间创建失败!")

	AddProjectApplicationErr = errorx.New(102118, "项目应用添加失败!")
	// 应用创建成功，但是项目资源修改失败，请稍后手动同步项目资源
	ProjectQuotaUpdateErr   = errorx.New(102119, "应用创建成功，但是项目资源修改失败，请稍后手动同步项目资源!")
	GetApplicationErr       = errorx.New(102120, "应用查询失败!")
	UpdateNamespaceQuotaErr = errorx.New(102121, "命名空间资源更新失败!")
	UpdateApplicationErr    = errorx.New(102122, "应用更新失败!")
	DeleteNamespaceQuotaErr = errorx.New(102123, "命名空间资源删除失败!")
	DeleteNamespaceErr      = errorx.New(102124, "命名空间删除失败!")
	DeleteApplicationErr    = errorx.New(102125, "应用删除失败!")
	ProjectIdEmptyErr       = errorx.New(102126, "项目ID 不能为空!")
	ProjectNotExistErr      = errorx.New(102127, "项目不存在!")
)
