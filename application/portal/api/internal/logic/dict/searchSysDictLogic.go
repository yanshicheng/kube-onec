package dict

import (
	"context"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/client/sysdictservice"

	"github.com/yanshicheng/kube-onec/application/portal/api/internal/svc"
	"github.com/yanshicheng/kube-onec/application/portal/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchSysDictLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchSysDictLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchSysDictLogic {
	return &SearchSysDictLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchSysDictLogic) SearchSysDict(req *types.SearchSysDictRequest) (resp *types.SearchSysDictResponse, err error) {
	res, err := l.svcCtx.SysDictRpc.SearchSysDict(l.ctx, &sysdictservice.SearchSysDictReq{
		DictCode:    req.DictCode,
		DictName:    req.DictName,
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
	data := make([]types.SysDict, len(res.Data))
	for i, v := range res.Data {
		data[i] = types.SysDict{
			Id:          v.Id,
			DictCode:    v.DictCode,
			DictName:    v.DictName,
			Description: v.Description,
			CreatedBy:   v.CreatedBy,
			CreatedAt:   v.CreatedAt,
			UpdatedBy:   v.UpdatedBy,
			UpdatedAt:   v.UpdatedAt,
		}
	}
	return &types.SearchSysDictResponse{
		Items: data,
		Total: res.Total,
	}, nil
}
