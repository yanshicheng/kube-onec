package onecclusterservicelogic

import (
	"context"
	"errors"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/internal/model"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/internal/svc"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/pb"
	"github.com/yanshicheng/kube-onec/common/handler/errorx"
	"github.com/yanshicheng/kube-onec/pkg/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchOnecClusterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchOnecClusterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchOnecClusterLogic {
	return &SearchOnecClusterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchOnecClusterLogic) SearchOnecCluster(in *pb.SearchOnecClusterReq) (*pb.SearchOnecClusterResp, error) {
	var queryParts []string
	var params []interface{}
	if in.Name != "" {
		queryParts = append(queryParts, "name = ? AND")
		params = append(params, "%"+in.Name+"%")
	}
	if in.Host != "" {
		queryParts = append(queryParts, "host = ? AND")
		params = append(params, "%"+in.Host+"%")
	}
	if in.EnvCode != "" {
		queryParts = append(queryParts, "env_tag = ? AND")
		params = append(params, in.EnvCode)
	}
	if in.Version != "" {
		queryParts = append(queryParts, "version = ? AND")
		params = append(params, "%"+in.Version+"%")
	}
	if in.Platform != "" {
		queryParts = append(queryParts, "platform = ? AND")
		params = append(params, "%"+in.Platform+"%")
	}
	if in.Version != "" {
		queryParts = append(queryParts, "version = ? AND")
		params = append(params, in.Version)
	}
	if in.UpdatedBy != "" {
		queryParts = append(queryParts, "update_by = ? AND")
		params = append(params, "%"+in.UpdatedBy+"%")
	}
	if in.CreatedBy != "" {
		queryParts = append(queryParts, "create_by = ? AND")
		params = append(params, "%"+in.CreatedBy+"%")
	}
	//if in.ConnType != 0 {
	//	queryStr.WriteString(" AND conn_type = ? ")
	//	params = append(params, pb.OnecClusterConnType_name[int32(in.ConnType)])
	//}
	query := utils.RemoveQueryADN(queryParts)
	matchedClusters, total, err := l.svcCtx.ClusterModel.Search(
		l.ctx,
		in.OrderStr,
		in.IsAsc,
		in.Page, in.PageSize,
		query, params...)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			l.Logger.Infof("查询集群为空:%v ,sql: %v", err, query)
			return &pb.SearchOnecClusterResp{
				Data:  make([]*pb.OnecCluster, 0),
				Total: 0,
			}, nil
		}
		l.Logger.Errorf("查询集群失败: %v", err)
		return nil, errorx.DatabaseQueryErr
	}
	var data []*pb.OnecCluster
	for _, v := range matchedClusters {
		data = append(data, ConvertModelToPBCluster(v))
	}
	return &pb.SearchOnecClusterResp{
		Data:  data,
		Total: total,
	}, nil
}
