package position

import (
	"context"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/client/syspositionservice"

	"github.com/yanshicheng/kube-onec/application/portal/api/internal/svc"
	"github.com/yanshicheng/kube-onec/application/portal/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchSysPositionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchSysPositionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchSysPositionLogic {
	return &SearchSysPositionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchSysPositionLogic) SearchSysPosition(req *types.SearchSysPositionRequest) (resp *types.SearchSysPositionResponse, err error) {
	// todo: add your logic here and delete this line
	res, err := l.svcCtx.SysPositionRpc.SearchSysPosition(l.ctx, &syspositionservice.SearchSysPositionReq{
		Page:     req.Page,
		PageSize: req.PageSize,
		IsAsc:    req.IsAsc,
		OrderStr: req.OrderStr,
		Name:     req.Name,
	})
	if err != nil {
		// 错误日志，明确上下文
		logx.WithContext(l.ctx).Errorf("查询职位信息失败: 请求=%+v, 错误=%v", req, err)
		return nil, err
	}
	resp = &types.SearchSysPositionResponse{
		Items: make([]types.SysPosition, 0, len(res.Data)),
		Total: res.Total,
	}
	for _, v := range res.Data {
		resp.Items = append(resp.Items, types.SysPosition{
			Id:         v.Id,
			Name:       v.Name,
			CreateTime: v.CreateTime,
			UpdateTime: v.UpdateTime,
		})
	}
	resp.Total = res.Total
	return
}
