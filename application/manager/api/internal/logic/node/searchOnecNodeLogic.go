package node

import (
	"context"
	"github.com/yanshicheng/kube-onec/application/manager/api/internal/svc"
	"github.com/yanshicheng/kube-onec/application/manager/api/internal/types"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/client/onecnodeservice"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchOnecNodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchOnecNodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchOnecNodeLogic {
	return &SearchOnecNodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchOnecNodeLogic) SearchOnecNode(req *types.SearchOnecNodeRequest) (resp *types.SearchOnecNodeResponse, err error) {
	// 调用 RPC 方法，将请求参数映射到 RPC 请求结构
	res, err := l.svcCtx.NodeRpc.SearchOnecNode(l.ctx, &onecnodeservice.SearchOnecNodeReq{
		Page:             req.Page,             // 当前页码
		PageSize:         req.PageSize,         // 每页条数
		OrderStr:         req.OrderStr,         // 排序字段
		IsAsc:            req.IsAsc,            // 是否升序
		ClusterUuid:      req.ClusterUuid,      // 所属集群ID
		NodeName:         req.NodeName,         // 节点名称
		NodeUid:          req.NodeUid,          // 节点UID
		Status:           req.Status,           // 节点状态
		Roles:            req.Roles,            // 节点角色列表
		PodCidr:          req.PodCidr,          // Pod CIDR
		Unschedulable:    req.Unschedulable,    // 是否不可调度
		NodeIp:           req.NodeIp,           // 节点IP地址
		Os:               req.Os,               // 操作系统
		ContainerRuntime: req.ContainerRuntime, // 容器运行时
		OperatingSystem:  req.OperatingSystem,  // 操作系统类型
		Architecture:     req.Architecture,     // 节点架构
		CreatedBy:        req.CreatedBy,        // 创建人
		UpdatedBy:        req.UpdatedBy,        // 更新人
	})
	if err != nil {
		// 错误处理
		l.Logger.Errorf("搜索节点失败: %v", err)
		return nil, err
	}

	// 将 RPC 返回的节点数据列表映射到 API 响应结构
	resp = &types.SearchOnecNodeResponse{
		Total: res.Total,              // 总条数
		Items: convertNodes(res.Data), // 转换节点信息
	}

	return resp, nil
}

// 辅助函数：将 RPC 的节点信息列表转换为 API 返回的节点信息列表
func convertNodes(data []*onecnodeservice.OnecNode) []types.OnecNode {
	var nodes []types.OnecNode
	for _, n := range data {
		nodes = append(nodes, types.OnecNode{
			Id:               n.Id,
			ClusterUuid:      n.ClusterUuid,
			NodeName:         n.NodeName,
			Cpu:              n.Cpu,
			Memory:           n.Memory,
			MaxPods:          n.MaxPods,
			IsGpu:            n.IsGpu,
			NodeUid:          n.NodeUid,
			Status:           n.Status,
			Roles:            n.Roles,
			JoinAt:           n.JoinAt,
			PodCidr:          n.PodCidr,
			Unschedulable:    n.Unschedulable,
			NodeIp:           n.NodeIp,
			Os:               n.Os,
			KernelVersion:    n.KernelVersion,
			ContainerRuntime: n.ContainerRuntime,
			KubeletVersion:   n.KubeletVersion,
			KubeletPort:      n.KubeletPort,
			OperatingSystem:  n.OperatingSystem,
			Architecture:     n.Architecture,
			CreatedBy:        n.CreatedBy,
			UpdatedBy:        n.UpdatedBy,
			CreatedAt:        n.CreatedAt,
			UpdatedAt:        n.UpdatedAt,
		})
	}
	return nodes
}
