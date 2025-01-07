package cluster

import (
	"context"
	"github.com/yanshicheng/kube-onec/application/manager/api/internal/svc"
	"github.com/yanshicheng/kube-onec/application/manager/api/internal/types"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/client/onecclusterservice"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchOnecClusterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchOnecClusterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchOnecClusterLogic {
	return &SearchOnecClusterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchOnecClusterLogic) SearchOnecCluster(req *types.SearchOnecClusterRequest) (resp *types.SearchOnecClusterResponse, err error) {
	// 调用 RPC 方法，将请求参数映射到 RPC 请求结构
	res, err := l.svcCtx.ClusterRpc.SearchOnecCluster(l.ctx, &onecclusterservice.SearchOnecClusterReq{
		Page:        req.Page,        // 当前页码
		PageSize:    req.PageSize,    // 每页条数
		OrderStr:    req.OrderStr,    // 排序字段
		IsAsc:       req.IsAsc,       // 是否升序
		Name:        req.Name,        // 集群名称
		Uuid:        req.UUID,        // 集群 UUID
		Host:        req.Host,        // 集群主机地址
		EnvCode:     req.EnvCode,     // 集群环境标签
		Status:      req.Status,      // 集群状态
		Version:     req.Version,     // 集群版本
		Platform:    req.Platform,    // 集群平台
		Location:    req.Location,    // 集群所在地址
		NodeLbIp:    req.NodeLbIp,    // Node 负载均衡 IP
		Description: req.Description, // 集群描述信息
		CreatedBy:    req.CreatedBy,    // 创建人
		UpdatedBy:    req.UpdatedBy,    // 更新人
	})
	if err != nil {
		l.Logger.Errorf("搜索集群失败: %v", err)
		return nil, err
	}
	data := make([]types.OnecCluster, len(res.Data))
	for i, c := range res.Data {
		data[i] = convertCluster(c)
	}
	// 映射 RPC 返回的结构到 API 响应结构
	resp = &types.SearchOnecClusterResponse{
		Total: res.Total, // 总条数
		Items: data,      // 转换每个集群信息
	}

	return resp, nil
}
