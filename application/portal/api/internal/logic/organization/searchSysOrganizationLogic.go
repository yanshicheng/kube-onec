package organization

import (
	"context"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/client/sysorganizationservice"

	"github.com/yanshicheng/kube-onec/application/portal/api/internal/svc"
	"github.com/yanshicheng/kube-onec/application/portal/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchSysOrganizationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchSysOrganizationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchSysOrganizationLogic {
	return &SearchSysOrganizationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}
func (l *SearchSysOrganizationLogic) SearchSysOrganization(req *types.SearchSysOrganizationRequest) (resp *types.SearchSysOrganizationResponse, err error) {
	// 调用 RPC 接口，获取搜索结果
	res, err := l.svcCtx.SysOrganizationRpc.SearchSysOrganization(l.ctx, &sysorganizationservice.SearchSysOrganizationReq{
		Name:        req.Name,
		Description: req.Description,
		IsAsc:       req.IsAsc,
		OrderStr:    req.OrderStr,
		ParentId:    req.ParentId,
	})
	if err != nil {
		logx.WithContext(l.ctx).Errorf("搜索机构失败: 请求=%+v, 错误=%v", req, err)
		return nil, err
	}

	// 映射 RPC 响应到 API 响应
	resp = &types.SearchSysOrganizationResponse{
		Items: make([]types.OrganizationNode, len(res.Data)),
	}

	// 处理每个组织节点的映射
	for i, org := range res.Data {
		resp.Items[i] = mapSysOrganizationToOrganizationNode(org)
	}

	return resp, nil
}

// 将 SysOrganizationSearch 转换为 OrganizationNode
func mapSysOrganizationToOrganizationNode(org *sysorganizationservice.SysOrganizationSearch) types.OrganizationNode {
	return types.OrganizationNode{
		Id:          org.Id,
		Name:        org.Name,
		ParentId:    org.ParentId,
		Level:       org.Level,
		Description: org.Description,
		CreatedAt:   org.CreatedAt,
		UpdatedAt:   org.UpdatedAt,
		Children:    mapChildrenToOrganizationNode(org.Children),
	}
}

// 递归处理子机构
func mapChildrenToOrganizationNode(children []*sysorganizationservice.SysOrganizationSearch) []types.OrganizationNode {
	var mappedChildren []types.OrganizationNode
	for _, child := range children {
		mappedChildren = append(mappedChildren, mapSysOrganizationToOrganizationNode(child))
	}
	return mappedChildren
}
