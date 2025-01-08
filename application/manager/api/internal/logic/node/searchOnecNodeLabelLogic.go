package node

import (
	"context"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/pb"

	"github.com/yanshicheng/kube-onec/application/manager/api/internal/svc"
	"github.com/yanshicheng/kube-onec/application/manager/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchOnecNodeLabelLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchOnecNodeLabelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchOnecNodeLabelLogic {
	return &SearchOnecNodeLabelLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchOnecNodeLabelLogic) SearchOnecNodeLabel(req *types.SearchNodeLabelsRequest) (resp *types.SearchNodeLabelsResponse, err error) {
	res, err := l.svcCtx.NodeRpc.SearchOnecNodeLabelList(l.ctx, &pb.SearchOnecNodeLabelListReq{
		NodeId:   req.NodeId,
		Key:      req.Key,
		OrderStr: req.OrderStr,
		Page:     req.Page,
		PageSize: req.PageSize,
		IsAsc:    req.IsAsc,
	})
	if err != nil {
		// 错误日志，明确上下文
		l.Logger.Errorf("查询节点标签信息失败: 请求=%+v, 错误=%v", req, err)
		return nil, err
	}
	data := make([]types.NodeLabel, len(res.Data))
	for i, v := range res.Data {
		data[i] = types.NodeLabel{
			Id:           v.Id,
			ResourceId:   v.ResourceId,
			ResourceType: v.ResourceType,
			Key:          v.Key,
			Value:        v.Value,
			CreatedAt:    v.CreatedAt,
			UpdatedAt:    v.UpdatedAt,
		}
	}
	return &types.SearchNodeLabelsResponse{
		Items: data,
		Total: res.Total,
	}, nil
}
