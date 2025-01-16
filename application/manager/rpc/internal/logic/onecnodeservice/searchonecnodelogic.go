package onecnodeservicelogic

import (
	"context"
	"errors"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/internal/code"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/internal/model"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/internal/shared"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/internal/svc"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/pb"
	"github.com/yanshicheng/kube-onec/pkg/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchOnecNodeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchOnecNodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchOnecNodeLogic {
	return &SearchOnecNodeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 搜索节点
func (l *SearchOnecNodeLogic) SearchOnecNode(in *pb.SearchOnecNodeReq) (*pb.SearchOnecNodeResp, error) {
	// 构建查询条件
	var queryParts []string
	var params []interface{}
	if in.ClusterUuid == "" {
		return nil, code.ClusterUuidEmptyErr
	}
	queryParts = append(queryParts, "`cluster_uuid` = ? AND")
	params = append(params, in.ClusterUuid)
	//queryStr.WriteString(" AND unschedulable = ? ")
	//params = append(params, in.Unschedulable)
	//queryStr.WriteString(" AND sync_status = ? ")
	//params = append(params, in.SyncStatus)
	// 添加查询条件
	if in.NodeName != "" {
		queryParts = append(queryParts, "`node_name` LIKE ? AND")
		params = append(params, "%"+in.NodeName+"%")
	}
	if in.NodeUid != "" {
		queryParts = append(queryParts, "`node_uid` = ? AND")
		params = append(params, in.NodeUid)
	}
	if in.Status != "" {
		queryParts = append(queryParts, "`status` = ? AND")
		params = append(params, in.Status)
	}

	if in.Roles != "" {
		queryParts = append(queryParts, "`roles` LIKE ? AND")
		params = append(params, "%"+in.Roles+"%")
	}
	if in.PodCidr != "" {
		queryParts = append(queryParts, "`pod_cidr` LIKE ? AND")
		params = append(params, "%"+in.PodCidr+"%")
	}

	if in.NodeIp != "" {
		queryParts = append(queryParts, "`node_ip` LIKE ? AND")
		params = append(params, "%"+in.NodeIp+"%")
	}
	if in.Os != "" {
		queryParts = append(queryParts, "`os` LIKE ? AND")
		params = append(params, "%"+in.Os+"%")
	}
	if in.ContainerRuntime != "" {
		queryParts = append(queryParts, "`container_runtime` LIKE ? AND")
		params = append(params, "%"+in.ContainerRuntime+"%")
	}
	if in.OperatingSystem != "" {
		queryParts = append(queryParts, "`operating_system` LIKE ? AND")
		params = append(params, "%"+in.OperatingSystem+"%")
	}
	if in.Architecture != "" {
		queryParts = append(queryParts, "`architecture` LIKE ? AND")
		params = append(params, "%"+in.Architecture+"%")
	}
	if in.CreatedBy != "" {
		queryParts = append(queryParts, "`created_by` = ? AND")
		params = append(params, in.CreatedBy)
	}
	if in.UpdatedBy != "" {
		queryParts = append(queryParts, "`updated_by` = ? AND")
		params = append(params, in.UpdatedBy)
	}
	// 去掉最后一个 " AND "，避免 SQL 语法错误
	// 使用正则表达式去掉结尾的 "AND" 或 "AND "
	query := utils.RemoveQueryADN(queryParts)

	// 调用 NodeModel 的搜索方法
	res, total, err := l.svcCtx.NodeModel.Search(
		l.ctx,
		in.OrderStr, // 排序字段
		in.IsAsc,    // 排序方式
		in.Page, in.PageSize,
		query, params...)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			l.Logger.Infof("未找到匹配的节点, sql: %v", query)
			return &pb.SearchOnecNodeResp{
				Data:  make([]*pb.OnecNode, 0),
				Total: 0,
			}, nil
		}
		l.Logger.Errorf("查询节点失败: %v", err)
		return nil, code.QueryNodeErr
	}

	// 转换数据库模型为 PB 模型
	var data []*pb.OnecNode
	for _, v := range res {
		data = append(data, shared.ConvertDbModelToPbModel(v))
	}

	// 返回响应
	return &pb.SearchOnecNodeResp{
		Data:  data,
		Total: total,
	}, nil
}
