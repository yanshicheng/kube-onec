package onecclusterservicelogic

import (
	"context"
	"github.com/yanshicheng/kube-onec/common/handler/errorx"
	"strings"

	"github.com/yanshicheng/kube-onec/application/manager/rpc/internal/svc"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/pb"

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
	var queryStr strings.Builder
	var params []interface{}
	if in.Name != "" {
		queryStr.WriteString(" AND name = ? ")
		params = append(params, "%"+in.Name+"%")
	}
	if in.Host != "" {
		queryStr.WriteString(" AND host = ? ")
		params = append(params, "%"+in.Host+"%")
	}
	if in.EnvCode != "" {
		queryStr.WriteString(" AND env_tag = ? ")
		params = append(params, in.EnvCode)
	}
	if in.Version != "" {
		queryStr.WriteString(" AND version = ? ")
		params = append(params, "%"+in.Version+"%")
	}
	if in.Platform != "" {
		queryStr.WriteString(" AND platform = ? ")
		params = append(params, "%"+in.Platform+"%")
	}
	if in.Version != "" {
		queryStr.WriteString(" AND version = ? ")
		params = append(params, in.Version)
	}
	if in.UpdatedBy != "" {
		queryStr.WriteString(" AND update_by = ? ")
		params = append(params, "%"+in.UpdatedBy+"%")
	}
	if in.CreatedBy != "" {
		queryStr.WriteString(" AND create_by = ? ")
		params = append(params, "%"+in.CreatedBy+"%")
	}
	//if in.ConnType != 0 {
	//	queryStr.WriteString(" AND conn_type = ? ")
	//	params = append(params, pb.OnecClusterConnType_name[int32(in.ConnType)])
	//}
	query := queryStr.String()
	if len(query) > 0 {
		query = query[:len(query)-5] // 去掉 " AND "
	}
	matchedClusters, total, err := l.svcCtx.ClusterModel.Search(
		l.ctx,
		in.OrderStr,
		in.IsAsc,
		in.Page, in.PageSize,
		query, params...)
	if err != nil {
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
