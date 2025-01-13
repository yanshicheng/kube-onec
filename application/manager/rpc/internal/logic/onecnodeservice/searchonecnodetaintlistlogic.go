package onecnodeservicelogic

import (
	"context"
	"errors"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/internal/code"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/internal/model"
	"github.com/yanshicheng/kube-onec/pkg/utils"
	"strings"

	"github.com/yanshicheng/kube-onec/application/manager/rpc/internal/svc"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchOnecNodeTaintListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchOnecNodeTaintListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchOnecNodeTaintListLogic {
	return &SearchOnecNodeTaintListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchOnecNodeTaintListLogic) SearchOnecNodeTaintList(in *pb.SearchOnecNodeTaintListReq) (*pb.SearchOnecNodeTaintListResp, error) {
	var queryStr strings.Builder
	var params []interface{}
	if in.NodeId == 0 {
		l.Logger.Errorf("`nodeId` is empty")
		return nil, code.NodeIdEmptyErr
	}
	queryStr.WriteString("`node_id` = ? AND ")
	params = append(params, in.NodeId)
	if in.Key != "" {
		queryStr.WriteString("`key` LIKE ? AND ")
		params = append(params, "%"+in.Key+"%")
	}
	query := utils.RemoveQueryADN(queryStr)

	res, total, err := l.svcCtx.TaintsResourceModel.Search(l.ctx,
		in.OrderStr,
		in.IsAsc,
		in.Page, in.PageSize,
		query, params...)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			l.Logger.Infof("查询节点污点为空:%v ,sql: %v", err, query)
			return &pb.SearchOnecNodeTaintListResp{
				Data:  make([]*pb.NodeTaints, 0),
				Total: 0,
			}, nil
		}
		l.Logger.Errorf("SearchOnecNodeTaintList err: %v", err)
		return nil, code.SearchNodeTaintDBErr
	}

	data := make([]*pb.NodeTaints, len(res))
	for i, v := range res {
		data[i] = &pb.NodeTaints{
			Id:         v.Id,
			NodeId:     v.NodeId,
			Key:        v.Key,
			Value:      v.Value,
			EffectCode: v.EffectCode,
			CreatedAt:  v.CreatedAt.Unix(),
			UpdatedAt:  v.UpdatedAt.Unix(),
		}
	}
	return &pb.SearchOnecNodeTaintListResp{
		Data:  data,
		Total: total,
	}, nil
}
