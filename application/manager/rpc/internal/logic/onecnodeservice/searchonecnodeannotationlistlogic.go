package onecnodeservicelogic

import (
	"context"
	"errors"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/internal/code"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/internal/model"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/internal/svc"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/pb"
	"github.com/yanshicheng/kube-onec/pkg/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchOnecNodeAnnotationListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchOnecNodeAnnotationListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchOnecNodeAnnotationListLogic {
	return &SearchOnecNodeAnnotationListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchOnecNodeAnnotationListLogic) SearchOnecNodeAnnotationList(in *pb.SearchOnecNodeAnnotationListReq) (*pb.SearchOnecNodeAnnotationListResp, error) {
	var queryParts []string
	var params []interface{}
	if in.NodeId == 0 {
		l.Logger.Errorf("`nodeId` is empty")
		return nil, code.NodeIdEmptyErr
	}
	queryParts = append(queryParts, "`resource_id` = ? AND")
	params = append(params, in.NodeId)
	if in.Key != "" {
		queryParts = append(queryParts, "`key` LIKE ? AND")
		params = append(params, "%"+in.Key+"%")
	}
	query := utils.RemoveQueryADN(queryParts)

	res, total, err := l.svcCtx.AnnotationsResourceModel.Search(l.ctx,
		in.OrderStr,
		in.IsAsc,
		in.Page, in.PageSize,
		query, params...)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			l.Logger.Infof("查询节点注解为空:%v ,sql: %v", err, query)
			return &pb.SearchOnecNodeAnnotationListResp{
				Data:  make([]*pb.NodeAnnotations, 0),
				Total: 0,
			}, nil
		}
		l.Logger.Errorf("AnnotationsResourceModel err: %v", err)
		return nil, code.SearchNodeAnnotationDBErr
	}

	data := make([]*pb.NodeAnnotations, len(res))
	for i, v := range res {
		data[i] = &pb.NodeAnnotations{
			Id:           v.Id,
			ResourceId:   v.ResourceId,
			Key:          v.Key,
			Value:        v.Value,
			ResourceType: v.ResourceType,
			CreatedAt:    v.CreatedAt.Unix(),
			UpdatedAt:    v.UpdatedAt.Unix(),
		}
	}
	return &pb.SearchOnecNodeAnnotationListResp{
		Data:  data,
		Total: total,
	}, nil
}
