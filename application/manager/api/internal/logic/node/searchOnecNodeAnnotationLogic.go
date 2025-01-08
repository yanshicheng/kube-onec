package node

import (
	"context"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/pb"

	"github.com/yanshicheng/kube-onec/application/manager/api/internal/svc"
	"github.com/yanshicheng/kube-onec/application/manager/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchOnecNodeAnnotationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchOnecNodeAnnotationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchOnecNodeAnnotationLogic {
	return &SearchOnecNodeAnnotationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchOnecNodeAnnotationLogic) SearchOnecNodeAnnotation(req *types.SearchNodeAnnotationsRequest) (resp *types.SearchNodeAnnotationsResponse, err error) {
	res, err := l.svcCtx.NodeRpc.SearchOnecNodeAnnotationList(l.ctx, &pb.SearchOnecNodeAnnotationListReq{
		NodeId:   req.NodeId,
		Key:      req.Key,
		OrderStr: req.OrderStr,
		IsAsc:    req.IsAsc,
		Page:     req.Page,
		PageSize: req.PageSize,
	})
	if err != nil {
		// 错误日志，明确上下文
		l.Logger.Errorf("查询节点注解信息失败: 请求=%+v, 错误=%v", req, err)
		return nil, err
	}

	data := make([]types.NodeAnnotation, len(res.Data))
	for i, v := range res.Data {
		data[i] = types.NodeAnnotation{
			Id:           v.Id,
			ResourceId:   v.ResourceId,
			ResourceType: v.ResourceType,
			Key:          v.Key,
			Value:        v.Value,
			CreatedAt:    v.CreatedAt,
			UpdatedAt:    v.UpdatedAt,
		}
	}
	return &types.SearchNodeAnnotationsResponse{
		Items: data,
		Total: res.Total,
	}, nil
}
