package sysorganizationservicelogic

import (
	"context"
	"errors"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/internal/model"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/internal/svc"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/pb"
	"github.com/yanshicheng/kube-onec/common/handler/errorx"
	"github.com/yanshicheng/kube-onec/pkg/utils"
	"github.com/zeromicro/go-zero/core/logx"
)

type SearchSysOrganizationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchSysOrganizationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchSysOrganizationLogic {
	return &SearchSysOrganizationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

type OrgNode struct {
	*model.SysOrganization
	Children []*OrgNode
}

func (l *SearchSysOrganizationLogic) SearchSysOrganization(in *pb.SearchSysOrganizationReq) (*pb.SearchSysOrganizationResp, error) {
	// Step 1: 构建查询字符串和参数
	// 构建动态 SQL 查询条件
	var queryParts []string
	var params []interface{}

	// 动态拼接条件
	if in.Name != "" {
		queryParts = append(queryParts, "`name` LIKE ? AND ")
		params = append(params, "%"+in.Name+"%")
	}
	if in.Description != "" {
		queryParts = append(queryParts, "`description` LIKE ? AND ")
		params = append(params, "%"+in.Description+"%")
	}
	if in.ParentId != 0 {
		queryParts = append(queryParts, "`parent_id` = ? AND ")
		params = append(params, in.ParentId)
	}
	// 去掉最后一个 " AND "，避免 SQL 语法错误
	query := utils.RemoveQueryADN(queryParts)
	// Step 2: 执行搜索（不分页）
	matchedOrgs, err := l.svcCtx.SysOrganization.SearchNoPage(
		l.ctx,
		in.OrderStr, // 使用请求中的 orderStr
		in.IsAsc,    // 使用请求中的 isAsc
		query, params...)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			l.Logger.Infof("查询组织为空: %v, sql: %v", err, query)
			return &pb.SearchSysOrganizationResp{
				Data: make([]*pb.SysOrganizationSearch, 0),
			}, nil
		}
		l.Logger.Errorf("查询组织失败: %v", err)
		return nil, errorx.DatabaseQueryErr
	}

	if len(matchedOrgs) == 0 {
		return &pb.SearchSysOrganizationResp{
			Data: []*pb.SysOrganizationSearch{},
		}, nil
	}

	// Step 3: 获取所有匹配组织的祖先
	ancestorOrgs, err := l.getAllAncestors(matchedOrgs)
	if err != nil {
		l.Logger.Errorf("获取祖先组织失败: %v", err)
		return nil, errorx.DatabaseQueryErr
	}

	// Step 4: 合并匹配组织和其祖先，去重
	allOrgsMap := make(map[uint64]*model.SysOrganization)
	for _, org := range matchedOrgs {
		allOrgsMap[org.Id] = org
	}
	for _, org := range ancestorOrgs {
		allOrgsMap[org.Id] = org
	}

	// 将合并后的组织转换为切片
	var allOrgs []*model.SysOrganization
	for _, org := range allOrgsMap {
		allOrgs = append(allOrgs, org)
	}

	// Step 5: 构建树状结构
	tree, err := buildOrgTree(allOrgs)
	if err != nil {
		l.Logger.Errorf("构建组织树失败: %v", err)
		return nil, errorx.DatabaseProcessErr
	}

	// Step 6: 将树状结构转换为 Protobuf 结构并转换时间字段
	var data []*pb.SysOrganizationSearch
	for _, node := range tree {
		pbOrg, err := convertToProto(node)
		if err != nil {
			l.Logger.Errorf("转换 Protobuf 结构失败: %v", err)
			return nil, errorx.DatabaseCopyErr
		}
		data = append(data, pbOrg)
	}

	return &pb.SearchSysOrganizationResp{
		Data: data,
	}, nil
}

// getAllAncestors 获取所有匹配组织的祖先组织
func (l *SearchSysOrganizationLogic) getAllAncestors(matchedOrgs []*model.SysOrganization) ([]*model.SysOrganization, error) {
	ancestorMap := make(map[uint64]*model.SysOrganization)
	for _, org := range matchedOrgs {
		currentParentId := org.ParentId
		for currentParentId != 0 { // 假设 parent_id 为 0 表示根节点
			if _, exists := ancestorMap[currentParentId]; exists {
				break // 已经处理过该祖先
			}
			parentOrg, err := l.svcCtx.SysOrganization.FindOne(l.ctx, currentParentId)
			if err != nil {
				if errors.Is(err, model.ErrNotFound) {
					break // 找不到父组织，结束递归
				}
				return nil, err
			}
			ancestorMap[parentOrg.Id] = parentOrg
			currentParentId = parentOrg.ParentId
		}
	}

	var ancestors []*model.SysOrganization
	for _, org := range ancestorMap {
		ancestors = append(ancestors, org)
	}
	return ancestors, nil
}

// buildOrgTree 构建组织树
func buildOrgTree(orgs []*model.SysOrganization) ([]*OrgNode, error) {
	// 创建一个 map 以便快速查找组织
	orgMap := make(map[uint64]*OrgNode)
	for _, org := range orgs {
		orgMap[org.Id] = &OrgNode{
			SysOrganization: org,
			Children:        []*OrgNode{},
		}
	}

	var roots []*OrgNode
	for _, node := range orgMap {
		if node.ParentId == 0 {
			// 根节点
			roots = append(roots, node)
		} else {
			parentNode, exists := orgMap[node.ParentId]
			if exists {
				parentNode.Children = append(parentNode.Children, node)
			} else {
				// 如果父节点不存在（可能是数据问题），将其视为根节点
				roots = append(roots, node)
			}
		}
	}

	return roots, nil
}

// convertToProto 将 OrgNode 转换为 Protobuf SysOrganizationSearch，并处理时间字段
func convertToProto(node *OrgNode) (*pb.SysOrganizationSearch, error) {
	if node == nil {
		return nil, nil
	}

	pbOrg := &pb.SysOrganizationSearch{
		Id:          node.Id,
		Name:        node.Name,
		ParentId:    node.ParentId,
		Level:       node.Level,
		Description: node.Description,
		CreatedAt:   node.CreatedAt.Unix(),
		UpdatedAt:   node.UpdatedAt.Unix(),
	}

	for _, child := range node.Children {
		pbChild, err := convertToProto(child)
		if err != nil {
			return nil, err
		}
		pbOrg.Children = append(pbOrg.Children, pbChild)
	}

	return pbOrg, nil
}
