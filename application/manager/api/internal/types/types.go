// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3

package types

type AddOnecClusterConnInfoRequest struct {
	ClusterUuid  string `json:"clusterUuid" validate:"required"`                // 关联集群 ID，必填
	ServiceCode  string `json:"serviceCode" validate:"required,max=100"`        // 服务名称，必填，最长 100 字符
	ServiceUrl   string `json:"serviceUrl" validate:"required,url"`             // 服务的 URL，必填，必须是有效 URL
	Username     string `json:"username,optional" validate:"omitempty,max=100"` // 用户名，可选
	Password     string `json:"password,optional" validate:"omitempty,max=100"` // 密码，可选
	Token        string `json:"token,optional" validate:"omitempty"`            // 令牌，可选
	SkipInsecure int64  `json:"skipInsecure,optional" validate:"oneof=0 1"`     // 是否忽略自签名证书验证，可选
	CaCert       string `json:"caCert,optional" validate:"omitempty"`           // CA 证书内容，可选
	ClientCert   string `json:"clientCert,optional" validate:"omitempty"`       // 客户端证书内容，可选
	ClientKey    string `json:"clientKey,optional" validate:"omitempty"`        // 客户端私钥内容，可选
}

type AddOnecClusterRequest struct {
	Name         string `json:"name,optional" validate:"required,min=2,max=50"`      // 集群名称，必填，长度 2-50 字符
	SkipInsecure int64  `json:"skipInsecure,optional" validate:"required,oneof=0 1"` // 是否跳过不安全连接，必填，值为 0 或 1
	Host         string `json:"host,optional" validate:"required,min=2,max=256"`     // 集群主机地址，必填，长度 2-256 字符
	Token        string `json:"token,optional" validate:"required"`                  // 集群访问令牌，必填
	ConnCode     string `json:"connCode,optional" validate:"required`                // 连接类型，必填
	EnvCode      string `json:"envCode,optional" validate:"omitempty,max=50"`        // 集群环境标签，可选，最长 50 字符
	Location     string `json:"location,optional" validate:"omitempty,max=255"`      // 集群所在地址，可选，最长 255 字符
	NodeLbIp     string `json:"nodeLbIp,optional" validate:"omitempty,max=100"`      // Node 负载均衡 IP，可选，最长 100 字符
	Description  string `json:"description,optional" validate:"omitempty,max=255"`   // 描述信息，可选，最长 255 字符
}

type AddOnecNodeAnnotationRequest struct {
	Id    uint64 `json:"id" validate:"required"`    // 节点ID
	Key   string `json:"key" validate:"required"`   // 注解键
	Value string `json:"value" validate:"required"` // 注解值
}

type AddOnecNodeLabelRequest struct {
	Id    uint64 `json:"id" validate:"required"`    // 节点ID
	Key   string `json:"key" validate:"required"`   // 标签键
	Value string `json:"value" validate:"required"` // 标签值
}

type AddOnecNodeRequest struct {
	ClusterUuid string `json:"clusterUuid" validate:"required"`            // 所属集群UUID，必填
	NodeName    string `json:"nodeName" validate:"required,min=2,max=50"`  // 节点名称，在同一集群中唯一，必填，长度 2-50 字符
	Ipaddr      string `json:"ipaddr" validate:"required,ip"`              // 节点 IP 地址，必填，必须是有效的 IP
	User        string `json:"user" validate:"required,max=50"`            // 节点用户名，必填，最长 50 字符
	Password    string `json:"password" validate:"required,min=6,max=100"` // 节点密码，必填，长度 6-100 字符
}

type AddOnecNodeTaintRequest struct {
	Id     uint64 `json:"id" validate:"required"`
	Key    string `json:"key" validate:"required"`
	Value  string `json:"value" validate:"required"`
	Effect string `json:"effect" validate:"required"`
}

type AddOnecProjectQuotaRequest struct {
	ClusterUuid      string  `json:"clusterUuid" validate:"required,max=100"`   // 关联的集群 ID，必填，最长 100 字符
	ProjectId        uint64  `json:"projectId" validate:"required,gt=0"`        // 关联的项目 ID，必填
	CpuQuota         int64   `json:"cpuQuota" validate:"required,min=0"`        // CPU 分配配额（单位：核），必填
	CpuOvercommit    float64 `json:"cpuOvercommit" validate:"required,gt=0"`    // CPU 超配比，必填，必须大于 0
	CpuLimit         float64 `json:"cpuLimit" validate:"required,min=0"`        // CPU 上限值（单位：核），必填
	MemoryQuota      float64 `json:"memoryQuota" validate:"required,min=0"`     // 内存分配配额（单位：GiB），必填
	MemoryOvercommit float64 `json:"memoryOvercommit" validate:"required,gt=0"` // 内存超配比，必填，必须大于 0
	MemoryLimit      float64 `json:"memoryLimit" validate:"required,min=0"`     // 内存上限值（单位：GiB），必填
	StorageLimit     int64   `json:"storageLimit" validate:"required,min=0"`    // 项目可使用的存储总量（单位：GiB），必填
	ConfigmapLimit   int64   `json:"configmapLimit" validate:"required,min=0"`  // ConfigMap 数量上限，必填
	SecretLimit      int64   `json:"secretLimit" validate:"required,min=0"`
	PvcLimit         int64   `json:"pvcLimit" validate:"required,min=0"`      // PVC 数量上限，必填
	PodLimit         int64   `json:"podLimit" validate:"required,min=0"`      // Pod 数量上限，必填
	NodeportLimit    int64   `json:"nodeportLimit" validate:"required,min=0"` // NodePort 数量上限，必填
}

type AddOnecProjectRequest struct {
	Name        string `json:"name,optional" validate:"required,max=255"`          // 项目的中文名称，必填，最长 255 字符
	Identifier  string `json:"identifier,optional" validate:"required,max=100"`    // 项目的唯一标识符（英文），必填，最长 100 字符
	Description string `json:"description,optional" validate:"omitempty,max=1000"` // 项目描述信息，可选，最长 1000 字符
}

type DefaultIdRequest struct {
	Id uint64 `path:"id"  validate:"required,gt=0"` // 用户 ID
}

type DelOnecNodeAnnotationRequest struct {
	Id uint64 `path:"id" validate:"required"` // 节点ID
}

type DelOnecNodeLabelRequest struct {
	Id uint64 `path:"id" validate:"required"` // 节点ID
}

type DelOnecNodeRequest struct {
	Id uint64 `path:"id" validate:"required"` // 节点ID
}

type DelOnecNodeTaintRequest struct {
	Id uint64 `json:"id" validate:"required"`
}

type EnableScheduledRequest struct {
	Id uint64 `path:"id" validate:"required"`
}

type ForbidScheduledRequest struct {
	Id uint64 `path:"id" validate:"required"`
}

type GetOnecProjectQuotaRequest struct {
	ProjectId   uint64 `path:"projectId" validate:"required,gt=0"`
	ClusterUuid string `path:"clusterUuid" validate:"required,max=100"`
}

type NodeAnnotation struct {
	Id           uint64 `json:"id"`
	ResourceId   uint64 `json:"resourceId"`
	ResourceType string `json:"resourceType"`
	Key          string `json:"key"`
	Value        string `json:"value"`
	CreatedAt    int64  `json:"createdAt"`
	UpdatedAt    int64  `json:"updatedAt"`
}

type NodeLabel struct {
	Id           uint64 `json:"id"`
	ResourceId   uint64 `json:"resourceId"`
	ResourceType string `json:"resourceType"`
	Key          string `json:"key"`
	Value        string `json:"value"`
	CreatedAt    int64  `json:"createdAt"`
	UpdatedAt    int64  `json:"updatedAt"`
}

type NodeTaint struct {
	Id         uint64 `json:"id"`
	NodeId     uint64 `json:"nodeId"`
	Key        string `json:"key"`
	Value      string `json:"value"`
	EffectCode string `json:"effectCode"`
	CreatedAt  int64  `json:"createdAt"`
	UpdatedAt  int64  `json:"updatedAt"`
}

type OnecCluster struct {
	Id               uint64  `json:"id"`               // 自增主键
	Name             string  `json:"name"`             // 集群名称
	UUID             string  `json:"uuid"`             // 集群唯一标识
	SkipInsecure     int64   `json:"skipInsecure"`     // 是否跳过不安全连接（0：否，1：是）
	Host             string  `json:"host"`             // 集群主机地址
	EnvName          string  `json:"envName"`          // 访问集群的令牌
	ConnCode         string  `json:"connType"`         // 连接类型
	EnvCode          string  `json:"envCode"`          // 集群环境标签 数据字典表
	Status           int64   `json:"status"`           // 集群状态
	Version          string  `json:"version"`          // 集群版本
	Commit           string  `json:"commit"`           // 集群提交版本
	Platform         string  `json:"platform"`         // 集群平台
	VersionBuildAt   int64   `json:"versionBuildAt"`   // 版本构建时间
	ClusterCreatedAt int64   `json:"clusterCreatedAt"` // 集群创建时间
	NodeCount        int64   `json:"nodeCount"`        // 节点数量
	CpuTotal         int64   `json:"cpuTotal"`         // 总 CPU
	MemoryTotal      float64 `json:"memoryTotal"`      // 总内存
	PodTotal         int64   `json:"podTotal"`         // 最大 Pod 数量
	CpuUsed          int64   `json:"cpuUsed"`          // 已使用的 CPU
	MemoryUsed       float64 `json:"memoryUsed"`       // 已使用的内存
	PodUsed          int64   `json:"podUsed"`          // 已使用的 Pod 数量
	Location         string  `json:"location"`         // 集群所在地址
	NodeLbIp         string  `json:"nodeLbIp"`         // Node 负载均衡 IP
	Description      string  `json:"description"`      // 集群描述信息
	CreatedBy        string  `json:"createdBy"`        // 记录创建人
	UpdatedBy        string  `json:"updatedBy"`        // 记录更新人
	CreatedAt        int64   `json:"createdAt"`        // 记录创建时间
	UpdatedAt        int64   `json:"updatedAt"`        // 记录更新时间
	Token            string  `json:"token"`            // 集群访问 token
}

type OnecClusterConnInfo struct {
	Id           uint64 `json:"id"`           // 自增主键
	ClusterUuid  string `json:"clusterUuid"`  // 关联集群 ID
	ServiceCode  string `json:"serviceCode"`  // 服务名称
	ServiceUrl   string `json:"serviceUrl"`   // 服务的 URL
	Username     string `json:"username"`     // 用户名（如果使用基本认证）
	Password     string `json:"password"`     // 密码（如果使用基本认证）
	Token        string `json:"token"`        // 令牌（如果使用 Token 认证）
	SkipInsecure int64  `json:"skipInsecure"` // 是否忽略自签名证书验证
	CaCert       string `json:"caCert"`       // CA 证书内容（以 PEM 格式存储）
	ClientCert   string `json:"clientCert"`   // 客户端证书内容（以 PEM 格式存储）
	ClientKey    string `json:"clientKey"`    // 客户端私钥内容（以 PEM 格式存储，仅客户端证书需要）
	CreatedBy    string `json:"createdBy"`    // 记录创建人
	UpdatedBy    string `json:"updatedBy"`    // 记录更新人
	CreatedAt    int64  `json:"createdAt"`    // 记录创建时间
	UpdatedAt    int64  `json:"updatedAt"`    // 记录更新时间
}

type OnecNode struct {
	Id               uint64  `json:"id"`            // 自增主键
	ClusterUuid      string  `json:"clusterUuid"`   // 所属集群ID
	NodeName         string  `json:"nodeName"`      // 节点名称
	Cpu              int64   `json:"cpu"`           // CPU核数
	Memory           float64 `json:"memory"`        // 内存大小 (Mi)
	MaxPods          int64   `json:"maxPods"`       // 最大Pod数量
	IsGpu            int64   `json:"isGpu"`         // 是否包含GPU
	NodeUid          string  `json:"nodeUid"`       // 节点UID
	Status           string  `json:"status"`        // 节点状态
	Roles            string  `json:"roles"`         // 节点角色
	JoinAt           int64   `json:"joinAt"`        // 加入集群时间
	PodCidr          string  `json:"podCidr"`       // Pod CIDR
	Unschedulable    int64   `json:"unschedulable"` // 是否不可调度
	SyncStatus       int64   `json:"syncStatus"`
	NodeIp           string  `json:"nodeIp"`           // 节点地址
	Os               string  `json:"os"`               // 操作系统
	KernelVersion    string  `json:"kernelVersion"`    // 内核版本
	ContainerRuntime string  `json:"containerRuntime"` // 容器运行时
	KubeletVersion   string  `json:"kubeletVersion"`   // Kubelet版本
	KubeletPort      int64   `json:"kubeletPort"`      // Kubelet端口
	OperatingSystem  string  `json:"operatingSystem"`  // 操作系统类型
	Architecture     string  `json:"architecture"`     // 架构类型
	CreatedBy        string  `json:"createdBy"`        // 创建人
	UpdatedBy        string  `json:"updatedBy"`        // 更新人
	CreatedAt        int64   `json:"createdAt"`        // 创建时间
	UpdatedAt        int64   `json:"updatedAt"`        // 更新时间
}

type OnecProject struct {
	Id          uint64 `json:"id"`          // 主键，自增 ID
	Name        string `json:"name"`        // 项目的中文名称
	Identifier  string `json:"identifier"`  // 项目的唯一标识符（英文）
	Description string `json:"description"` // 项目描述信息
	IsDefault   int64  `json:"isDefault"`   // 是否为默认项目
	CreatedBy   string `json:"createdBy"`   // 创建人
	UpdatedBy   string `json:"updatedBy"`   // 更新人
	CreatedAt   int64  `json:"createdAt"`   // 创建时间
	UpdatedAt   int64  `json:"updatedAt"`   // 最后更新时间
}

type OnecProjectQuota struct {
	Id                   uint64  `json:"id"`            // 主键，自增 ID
	ClusterUuid          string  `json:"clusterUuid"`   // 关联的集群 ID
	ProjectId            uint64  `json:"projectId"`     // 关联的项目 ID
	CpuQuota             int64   `json:"cpuQuota"`      // CPU 分配配额（单位：核）
	CpuOvercommit        float64 `json:"cpuOvercommit"` // CPU 超配比（如 1.5 表示允许超配 50%）
	CpuLimit             float64 `json:"cpuLimit"`      // CPU 上限值（单位：核）
	CpuUsed              float64 `json:"cpuUsed"`       // 已使用的 CPU 资源（单位：核）
	CpuLimitRemain       float64 `json:"cpuLimitRemain"`
	MemoryQuota          float64 `json:"memoryQuota"`      // 内存分配配额（单位：GiB）
	MemoryOvercommit     float64 `json:"memoryOvercommit"` // 内存超配比（如 1.2 表示允许超配 20%）
	MemoryLimit          float64 `json:"memoryLimit"`      // 内存上限值（单位：GiB）
	MemoryUsed           float64 `json:"memoryUsed"`       // 已使用的内存资源（单位：GiB）
	MemoryLimitRemain    float64 `json:"memoryLimitRemain"`
	StorageLimit         int64   `json:"storageLimit"` // 项目可使用的存储总量（单位：GiB）
	StorageUsed          int64   `json:"storageUsed"`
	StorageLimitRemain   int64   `json:"storageLimitRemain"`
	ConfigmapLimit       int64   `json:"configmapLimit"` // 项目允许创建的 ConfigMap 数量上限
	ConfigMapUsed        int64   `json:"configmapUsed"`
	ConfigMapLimitRemain int64   `json:"configmapLimitRemain"`
	SecretLimit          int64   `json:"secretLimit"`
	SecretUsed           int64   `json:"secretUsed"`
	SecretLimitRemain    int64   `json:"secretLimitRemain"`
	PvcLimit             int64   `json:"pvcLimit"` // 项目允许创建的 PVC（PersistentVolumeClaim）数量上限
	PvcUsed              int64   `json:"pvcUsed"`
	PvcLimitRemain       int64   `json:"pvcLimitRemain"`
	PodLimit             int64   `json:"podLimit"` // 项目允许创建的 Pod 数量上限
	PodUsed              int64   `json:"podUsed"`
	PodLimitRemain       int64   `json:"podLimitRemain"`
	NodeportLimit        int64   `json:"nodeportLimit"` // 项目允许使用的 NodePort 数量上限
	NodeportUsed         int64   `json:"nodeportUsed"`
	NodeportLimitRemain  int64   `json:"nodeportLimitRemain"`
	Status               string  `json:"status"`    // 项目状态（如 Active、Disabled、Archived）
	CreatedBy            string  `json:"createdBy"` // 创建人
	UpdatedBy            string  `json:"updatedBy"` // 更新人
	CreatedAt            int64   `json:"createdAt"` // 创建时间
	UpdatedAt            int64   `json:"updatedAt"` // 最后更新时间
}

type SearchNodeAnnotationsRequest struct {
	Page     uint64 `form:"page,optional" default:"1" validate:"required,min=1"`              // 当前页
	PageSize uint64 `form:"pageSize,optional" default:"10" validate:"required,min=1,max=200"` // 每页条数
	OrderStr string `form:"orderStr,optional" default:"id" validate:"omitempty,max=50"`       // 排序字段
	IsAsc    bool   `form:"isAsc,optional" default:"false"`
	NodeId   uint64 `form:"nodeId,optional" validate:"required"`
	Key      string `form:"key,optional" validate:"omitempty"`
}

type SearchNodeAnnotationsResponse struct {
	Items []NodeAnnotation `json:"items"` // 注解数据列表
	Total uint64           `json:"total"` // 总条数
}

type SearchNodeLabelsRequest struct {
	Page     uint64 `form:"page,optional" default:"1" validate:"required,min=1"`              // 当前页
	PageSize uint64 `form:"pageSize,optional" default:"10" validate:"required,min=1,max=200"` // 每页条数
	OrderStr string `form:"orderStr,optional" default:"id" validate:"omitempty,max=50"`       // 排序字段
	IsAsc    bool   `form:"isAsc,optional" default:"false"`
	NodeId   uint64 `form:"nodeId,optional" validate:"required"`
	Key      string `form:"key,optional" validate:"omitempty"`
}

type SearchNodeLabelsResponse struct {
	Items []NodeLabel `json:"items"` // 标签数据列表
	Total uint64      `json:"total"` // 总条数
}

type SearchNodeTaintsRequest struct {
	Page     uint64 `form:"page,optional" default:"1" validate:"required,min=1"`              // 当前页
	PageSize uint64 `form:"pageSize,optional" default:"10" validate:"required,min=1,max=200"` // 每页条数
	OrderStr string `form:"orderStr,optional" default:"id" validate:"omitempty,max=50"`       // 排序字段
	IsAsc    bool   `form:"isAsc,optional" default:"false"`
	NodeId   uint64 `form:"nodeId,optional" validate:"required"`
	Key      string `form:"key,optional" validate:"omitempty"`
}

type SearchNodeTaintsResponse struct {
	Items []NodeTaint `json:"items"` // 污点数据列表
	Total uint64      `json:"total"` // 总条数
}

type SearchOnecClusterConnInfoRequest struct {
	Page        uint64 `form:"page,optional" default:"1" validate:"required,min=1"`              // 当前页，必填，最小值 1
	PageSize    uint64 `form:"pageSize,optional" default:"10" validate:"required,min=1,max=200"` // 每页条数，必填，范围 1-200
	OrderStr    string `form:"orderStr,optional" default:"id" validate:"omitempty,max=50"`       // 排序字段，默认 "id"
	IsAsc       bool   `form:"isAsc,optional" default:"false"`                                   // 是否升序，默认 false
	ClusterUuid string `form:"clusterUuid,optional" validate:"omitempty,max=50"`                 // 关联集群 ID，可选
	ServiceCode string `form:"serviceCode,optional" validate:"omitempty,max=50"`                 // 服务名称，可选
}

type SearchOnecClusterConnInfoResponse struct {
	Items []OnecClusterConnInfo `json:"items"` // 服务连接信息数据列表
	Total uint64                `json:"total"` // 总条数
}

type SearchOnecClusterRequest struct {
	Page        uint64 `form:"page,optional" default:"1" validate:"required,min=1"`              // 当前页码，必填，最小值为 1，默认 1
	PageSize    uint64 `form:"pageSize,optional" default:"10" validate:"required,min=1,max=200"` // 每页条数，必填，范围 1-200，默认 10
	OrderStr    string `form:"orderStr,optional" default:"id" validate:"omitempty,max=50"`       // 排序字段，默认是 id
	IsAsc       bool   `form:"isAsc,optional" default:"false"`                                   // 是否升序，默认 false
	Name        string `form:"name,optional" validate:"omitempty,max=50"`                        // 集群名称，可选，最长 50 字符
	UUID        string `form:"uuid,optional" validate:"omitempty,max=50"`                        // 集群唯一标识，可选，最长 50 字符
	Host        string `form:"host,optional" validate:"omitempty,max=256"`                       // 集群主机地址，可选，最长 256 字符
	EnvCode     string `form:"envCode,optional" validate:"omitempty,max=50"`                     // 集群环境标签，可选，最长 50 字符
	Status      int64  `form:"status,optional" validate:"omitempty"`                             // 集群状态，可选
	Version     string `form:"version,optional" validate:"omitempty,max=50"`                     // 集群版本，可选，最长 50 字符
	Platform    string `form:"platform,optional" validate:"omitempty,max=50"`                    // 集群平台，可选，最长 50 字符
	Location    string `form:"location,optional" validate:"omitempty,max=255"`                   // 集群所在地址，可选，最长 255 字符
	NodeLbIp    string `form:"nodeLbIp,optional" validate:"omitempty,max=100"`                   // Node 负载均衡 IP，可选，最长 100 字符
	Description string `form:"description,optional" validate:"omitempty,max=255"`                // 描述信息，可选，最长 255 字符
	CreatedBy   string `form:"createdBy,optional" validate:"omitempty,max=100"`                  // 记录创建人，可选，最长 100 字符
	UpdatedBy   string `form:"updatedBy,optional" validate:"omitempty,max=100"`                  // 记录更新人，可选，最长 100 字符
}

type SearchOnecClusterResponse struct {
	Items []OnecCluster `json:"items"` // 集群数据列表
	Total uint64        `json:"total"` // 总条数
}

type SearchOnecNodeRequest struct {
	Page             uint64 `form:"page,optional" default:"1" validate:"required,min=1"`              // 当前页
	PageSize         uint64 `form:"pageSize,optional" default:"10" validate:"required,min=1,max=200"` // 每页条数
	OrderStr         string `form:"orderStr,optional" default:"id" validate:"omitempty,max=50"`       // 排序字段
	IsAsc            bool   `form:"isAsc,optional" default:"false"`                                   // 是否升序
	ClusterUuid      string `form:"clusterUuid,optional" validate:"required,max=50"`                  // 集群ID
	NodeName         string `form:"nodeName,optional" validate:"omitempty,max=50"`                    // 节点名称
	NodeUid          string `form:"nodeUid,optional" validate:"omitempty,max=50"`                     // 节点UID
	Status           string `form:"status,optional" validate:"omitempty"`
	SyncStatus       int64  `form:"syncStatus,optional" validate:"omitempty"`
	Roles            string `form:"roles,optional" validate:"omitempty"`
	PodCidr          string `form:"podCidr,optional" validate:"omitempty"`
	Unschedulable    int64  `form:"unschedulable,optional" validate:"omitempty"`
	NodeIp           string `form:"nodeIp,optional" validate:"omitempty"`
	Os               string `form:"os,optional" validate:"omitempty"`
	Architecture     string `form:"architecture,optional" validate:"omitempty"`
	ContainerRuntime string `form:"containerRuntime,optional" validate:"omitempty"`
	OperatingSystem  string `form:"operatingSystem,optional" validate:"omitempty"`
	CreatedBy        string `form:"createdBy,optional" validate:"omitempty"`
	UpdatedBy        string `form:"updatedBy,optional" validate:"omitempty"`
}

type SearchOnecNodeResponse struct {
	Items []OnecNode `json:"items"` // 节点数据列表
	Total uint64     `json:"total"` // 总条数
}

type SearchOnecProjectQuotaRequest struct {
	Page        uint64 `form:"page,optional" default:"1" validate:"required,min=1"`              // 当前页码，默认 1
	PageSize    uint64 `form:"pageSize,optional" default:"10" validate:"required,min=1,max=200"` // 每页条数，默认 10
	OrderStr    string `form:"orderStr,optional" default:"id" validate:"omitempty,max=50"`       // 排序字段，默认 id
	IsAsc       bool   `form:"isAsc,optional" default:"false"`                                   // 是否升序，默认 false
	ClusterUuid string `form:"clusterUuid,optional" validate:"omitempty,max=100"`                // 关联的集群 ID，可选，最长 100 字符
	ProjectId   uint64 `form:"projectId,optional" validate:"omitempty,gt=0"`                     // 关联的项目 ID，可选
	Status      string `form:"status,optional" validate:"omitempty,max=50"`                      // 项目状态，可选，最长 50 字符
	CreatedBy   string `form:"createdBy,optional" validate:"omitempty,max=50"`                   // 创建人，可选，最长 50 字符
	UpdatedBy   string `form:"updatedBy,optional" validate:"omitempty,max=50"`                   // 更新人，可选，最长 50 字符
}

type SearchOnecProjectQuotaResponse struct {
	Items []OnecProjectQuota `json:"items"` // 资源配额数据列表
	Total uint64             `json:"total"` // 总条数
}

type SearchOnecProjectRequest struct {
	Page       uint64 `form:"page,optional" default:"1" validate:"required,min=1"`              // 当前页码，默认 1
	PageSize   uint64 `form:"pageSize,optional" default:"10" validate:"required,min=1,max=200"` // 每页条数，默认 10
	OrderStr   string `form:"orderStr,optional" default:"id" validate:"omitempty,max=50"`       // 排序字段，默认 id
	IsAsc      bool   `form:"isAsc,optional" default:"false"`                                   // 是否升序，默认 false
	Name       string `form:"name,optional" validate:"omitempty,max=255"`                       // 项目的中文名称，可选，最长 255 字符
	Identifier string `form:"identifier,optional" validate:"omitempty,max=100"`                 // 项目的唯一标识符（英文），可选，最长 100 字符
	CreatedBy  string `form:"createdBy,optional" validate:"omitempty,max=50"`                   // 创建人，可选，最长 50 字符
	UpdatedBy  string `form:"updatedBy,optional" validate:"omitempty,max=50"`                   // 更新人，可选，最长 50 字符
}

type SearchOnecProjectResponse struct {
	Items []OnecProject `json:"items"` // 项目数据列表
	Total uint64        `json:"total"` // 总条数
}

type SyncOnecNodeRequest struct {
	Id uint64 `path:"id" validate:"required"`
}

type UpdateOnecClusterConnInfoRequest struct {
	Id           uint64 `path:"id" validate:"required"`                         // 自增主键，必填
	ClusterUuid  string `json:"clusterUuid" validate:"required"`                // 关联集群 ID，必填
	ServiceCode  string `json:"serviceCode" validate:"required,max=100"`        // 服务名称，必填，最长 100 字符
	ServiceUrl   string `json:"serviceUrl" validate:"required,url"`             // 服务的 URL，必填，必须是有效 URL
	Username     string `json:"username,optional" validate:"omitempty,max=100"` // 用户名，可选
	Password     string `json:"password,optional" validate:"omitempty,max=100"` // 密码，可选
	Token        string `json:"token,optional" validate:"omitempty"`            // 令牌，可选
	SkipInsecure int64  `json:"skipInsecure,optional" validate:"oneof=0 1"`     // 是否忽略自签名证书验证，可选
	CaCert       string `json:"caCert,optional" validate:"omitempty"`           // CA 证书内容，可选
	ClientCert   string `json:"clientCert,optional" validate:"omitempty"`       // 客户端证书内容，可选
	ClientKey    string `json:"clientKey,optional" validate:"omitempty"`        // 客户端私钥内容，可选
}

type UpdateOnecClusterRequest struct {
	Id           uint64 `path:"id" validate:"required,gt=0"`                                     // 集群 ID，必填，必须大于 0
	Name         string `json:"name" validate:"required,min=2,max=50"`                           // 集群名称，必填，长度 2-50 字符
	SkipInsecure int64  `json:"skipInsecure" validate:"required,oneof=0 1"`                      // 是否跳过不安全连接，必填，值为 0 或 1
	Host         string `json:"host" validate:"required,min=2,max=256"`                          // 集群主机地址，必填，长度 2-256 字符
	Token        string `json:"token" validate:"required"`                                       // 集群访问令牌，必填
	ConnCode     string `json:"connType" validate:"required,oneof=KUBECONFIG TOKEN AGENT OTHER"` // 连接类型，必填
	EnvCode      string `json:"envCode" validate:"omitempty,max=50"`                             // 集群环境标签，可选，最长 50 字符
	Location     string `json:"location,optional" validate:"omitempty,max=255"`                  // 集群所在地址，可选，最长 255 字符
	NodeLbIp     string `json:"nodeLbIp,optional" validate:"omitempty,max=100"`                  // Node 负载均衡 IP，可选，最长 100 字符
	Description  string `json:"description,optional" validate:"omitempty,max=255"`               // 描述信息，可选，最长 255 字符
}

type UpdateOnecProjectQuotaRequest struct {
	Id               uint64  `path:"id" validate:"required,gt=0"`               // 主键，自增 ID，必填
	ClusterUuid      string  `json:"clusterUuid" validate:"required,max=100"`   // 关联的集群 ID，必填，最长 100 字符
	ProjectId        uint64  `json:"projectId" validate:"required,gt=0"`        // 关联的项目 ID，必填
	CpuQuota         int64   `json:"cpuQuota" validate:"required,min=0"`        // CPU 分配配额（单位：核），必填
	CpuOvercommit    float64 `json:"cpuOvercommit" validate:"required,gt=0"`    // CPU 超配比，必填，必须大于 0
	CpuLimit         float64 `json:"cpuLimit" validate:"required,min=0"`        // CPU 上限值（单位：核），必填
	MemoryQuota      float64 `json:"memoryQuota" validate:"required,min=0"`     // 内存分配配额（单位：GiB），必填
	MemoryOvercommit float64 `json:"memoryOvercommit" validate:"required,gt=0"` // 内存超配比，必填，必须大于 0
	MemoryLimit      float64 `json:"memoryLimit" validate:"required,min=0"`     // 内存上限值（单位：GiB），必填
	StorageLimit     int64   `json:"storageLimit" validate:"required,min=0"`    // 项目可使用的存储总量（单位：GiB），必填
	SecretLimit      int64   `json:"secretLimit" validate:"required,min=0"`
	ConfigmapLimit   int64   `json:"configmapLimit" validate:"required,min=0"` // ConfigMap 数量上限，必填
	PvcLimit         int64   `json:"pvcLimit" validate:"required,min=0"`       // PVC 数量上限，必填
	PodLimit         int64   `json:"podLimit" validate:"required,min=0"`       // Pod 数量上限，必填
	NodeportLimit    int64   `json:"nodeportLimit" validate:"required,min=0"`  // NodePort 数量上限，必填
}

type UpdateOnecProjectRequest struct {
	Id          uint64 `path:"id" validate:"required,gt=0"`                        // 主键，自增 ID，必填
	Name        string `json:"name,optional" validate:"omitempty,max=255"`         // 项目的中文名称，可选，最长 255 字符
	Description string `json:"description,optional" validate:"omitempty,max=1000"` // 项目描述信息，可选，最长 1000 字符
}
