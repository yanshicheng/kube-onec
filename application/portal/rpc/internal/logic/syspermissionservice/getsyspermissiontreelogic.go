package syspermissionservicelogic

import (
	"context"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/internal/model"
	"github.com/yanshicheng/kube-onec/common/handler/errorx"

	"github.com/yanshicheng/kube-onec/application/portal/rpc/internal/svc"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSysPermissionTreeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetSysPermissionTreeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSysPermissionTreeLogic {
	return &GetSysPermissionTreeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// SysPermissionTreeNode 是一个辅助结构体，用于在内存中构建树状权限结构
type SysPermissionTreeNode struct {
	*model.SysPermission
	Children []*SysPermissionTreeNode
}

func (l *GetSysPermissionTreeLogic) GetSysPermissionTree(in *pb.GetSysPermissionTreeReq) (*pb.GetSysPermissionTreeResp, error) {
	// todo: add your logic here and delete this line
	// Step 1: 查询所有权限数据
	permissions, err := l.svcCtx.SysPermission.SearchNoPage(l.ctx, "", true, "")
	if err != nil {
		l.Logger.Errorf("查询权限数据失败: %v", err)
		return nil, errorx.DatabaseQueryErr
	}

	if len(permissions) == 0 {
		l.Logger.Info("权限数据为空")
		return &pb.GetSysPermissionTreeResp{
			Data: []*pb.SysPermissionTree{},
		}, nil
	}

	// Step 2: 构建树状结构
	tree, err := buildPermissionTree(permissions)
	if err != nil {
		l.Logger.Errorf("构建权限树失败: %v", err)
		return nil, errorx.DatabaseProcessErr
	}

	// Step 3: 转换为 Protobuf 结构
	var protoTree []*pb.SysPermissionTree
	for _, node := range tree {
		protoNode, err := convertToProtoSysPermissionTree(node)
		if err != nil {
			l.Logger.Errorf("转换 Protobuf 结构失败: %v", err)
			return nil, errorx.DatabaseCopyErr
		}
		protoTree = append(protoTree, protoNode)
	}
	return &pb.GetSysPermissionTreeResp{
		Data: protoTree,
	}, nil
}

// buildPermissionTree 构建权限树状结构
func buildPermissionTree(permissions []*model.SysPermission) ([]*SysPermissionTreeNode, error) {
	// 创建一个 map 以便快速查找权限节点
	permissionMap := make(map[uint64]*SysPermissionTreeNode)
	for _, perm := range permissions {
		permissionMap[perm.Id] = &SysPermissionTreeNode{
			SysPermission: perm,
			Children:      []*SysPermissionTreeNode{},
		}
	}

	var roots []*SysPermissionTreeNode
	for _, node := range permissionMap {
		if node.ParentId == 0 {
			// 根节点
			roots = append(roots, node)
		} else {
			parentNode, exists := permissionMap[node.ParentId]
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

// convertToProtoSysPermissionTree 将 SysPermissionTreeNode 转换为 Protobuf SysPermissionTree
func convertToProtoSysPermissionTree(node *SysPermissionTreeNode) (*pb.SysPermissionTree, error) {
	if node == nil {
		return nil, nil
	}

	protoNode := &pb.SysPermissionTree{
		Id:   node.Id,
		Name: node.Name,
	}

	for _, child := range node.Children {
		protoChild, err := convertToProtoSysPermissionTree(child)
		if err != nil {
			return nil, err
		}
		protoNode.Children = append(protoNode.Children, protoChild)
	}

	return protoNode, nil
}
