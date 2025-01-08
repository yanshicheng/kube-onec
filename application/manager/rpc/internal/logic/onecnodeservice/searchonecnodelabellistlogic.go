package onecnodeservicelogic

import (
	"context"
	"errors"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/internal/code"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/internal/model"
	"github.com/yanshicheng/kube-onec/utils"
	"strings"

	"github.com/yanshicheng/kube-onec/application/manager/rpc/internal/svc"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchOnecNodeLabelListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchOnecNodeLabelListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchOnecNodeLabelListLogic {
	return &SearchOnecNodeLabelListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchOnecNodeLabelListLogic) SearchOnecNodeLabelList(in *pb.SearchOnecNodeLabelListReq) (*pb.SearchOnecNodeLabelListResp, error) {
	var queryStr strings.Builder
	var params []interface{}
	if in.NodeId == 0 {
		l.Logger.Errorf("`nodeId` is empty")
		return nil, code.NodeIdEmptyErr
	}
	queryStr.WriteString("`resource_id` = ? AND ")
	params = append(params, in.NodeId)
	if in.Key != "" {
		queryStr.WriteString("`key` LIKE ? AND ")
		params = append(params, "%"+in.Key+"%")
	}
	query := utils.RemoveQueryADN(queryStr)
	res, total, err := l.svcCtx.LabelsResourceModel.Search(l.ctx,
		in.OrderStr,
		in.IsAsc,
		in.Page, in.PageSize,
		query, params...)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return &pb.SearchOnecNodeLabelListResp{
				Data:  make([]*pb.NodeLabels, 0),
				Total: 0,
			}, nil
		}
		l.Logger.Errorf("SearchNodeLabelDBErr err: %v", err)
		return nil, code.SearchNodeLabelDBErr
	}

	data := make([]*pb.NodeLabels, len(res))
	for i, v := range res {
		data[i] = &pb.NodeLabels{
			Id:           v.Id,
			ResourceId:   v.ResourceId,
			Key:          v.Key,
			Value:        v.Value,
			ResourceType: v.ResourceType,
			CreatedAt:    v.CreatedAt.Unix(),
			UpdatedAt:    v.UpdatedAt.Unix(),
		}
	}
	return &pb.SearchOnecNodeLabelListResp{
		Data:  data,
		Total: total,
	}, nil
}
