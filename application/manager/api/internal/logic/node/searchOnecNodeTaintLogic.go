package node

import (
	"context"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/pb"

	"github.com/yanshicheng/kube-onec/application/manager/api/internal/svc"
	"github.com/yanshicheng/kube-onec/application/manager/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchOnecNodeTaintLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchOnecNodeTaintLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchOnecNodeTaintLogic {
	return &SearchOnecNodeTaintLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchOnecNodeTaintLogic) SearchOnecNodeTaint(req *types.SearchNodeTaintsRequest) (resp *types.SearchNodeTaintsResponse, err error) {
	res, err := l.svcCtx.NodeRpc.SearchOnecNodeTaintList(l.ctx, &pb.SearchOnecNodeTaintListReq{
		IsAsc:    req.IsAsc,
		OrderStr: req.OrderStr,
		Page:     req.Page,
		PageSize: req.PageSize,
		NodeId:   req.NodeId,
		Key:      req.Key,
	})
	if err != nil {
		// 错误日志，明确上下文
		l.Logger.Errorf("查询节点污点信息失败: 请求=%+v, 错误=%v", req, err)
		return nil, err
	}
	data := make([]types.NodeTaint, len(res.Data))
	for i, v := range res.Data {
		data[i] = types.NodeTaint{
			Id:         v.Id,
			NodeId:     v.NodeId,
			EffectCode: v.EffectCode,
			Key:        v.Key,
			Value:      v.Value,
			CreatedAt:  v.CreatedAt,
			UpdatedAt:  v.UpdatedAt,
		}
	}
	return &types.SearchNodeTaintsResponse{
		Items: data,
		Total: res.Total,
	}, nil
}
