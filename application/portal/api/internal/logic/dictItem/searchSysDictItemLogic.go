package dictItem

import (
	"context"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/client/sysdictitemservice"

	"github.com/yanshicheng/kube-onec/application/portal/api/internal/svc"
	"github.com/yanshicheng/kube-onec/application/portal/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchSysDictItemLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchSysDictItemLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchSysDictItemLogic {
	return &SearchSysDictItemLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchSysDictItemLogic) SearchSysDictItem(req *types.SearchSysDictItemRequest) (resp *types.SearchSysDictItemResponse, err error) {
	res, err := l.svcCtx.SysDictItemRpc.SearchSysDictItem(l.ctx, &sysdictitemservice.SearchSysDictItemReq{
		DictCode:    req.DictCode,
		ItemText:    req.ItemText,
		ItemCode:    req.ItemCode,
		Description: req.Description,
		CreatedBy:   req.CreatedBy,
		UpdatedBy:   req.UpdatedBy,
		Page:        req.Page,
		PageSize:    req.PageSize,
		OrderStr:    req.OrderStr,
		IsAsc:       req.IsAsc,
	})
	if err != nil {
		// 错误日志，明确上下文
		l.Logger.Errorf("查询字典信息失败: 请求=%+v, 错误=%v", req, err)
		return nil, err
	}

	data := make([]types.SysDictItem, len(res.Data))
	for i, v := range res.Data {
		data[i] = types.SysDictItem{
			Id:          v.Id,
			DictCode:    v.DictCode,
			ItemText:    v.ItemText,
			ItemCode:    v.ItemCode,
			Description: v.Description,
			CreatedBy:   v.CreatedBy,
			CreatedAt:   v.CreatedAt,
			UpdatedBy:   v.UpdatedBy,
			UpdatedAt:   v.UpdatedAt,
			SortOrder:   v.SortOrder,
		}
	}
	return &types.SearchSysDictItemResponse{
		Items: data,
		Total: res.Total,
	}, nil
}
