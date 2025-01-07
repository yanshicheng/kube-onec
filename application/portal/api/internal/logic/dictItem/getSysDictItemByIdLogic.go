package dictItem

import (
	"context"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/pb"
	"github.com/yanshicheng/kube-onec/common/handler/errorx"

	"github.com/yanshicheng/kube-onec/application/portal/api/internal/svc"
	"github.com/yanshicheng/kube-onec/application/portal/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSysDictItemByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetSysDictItemByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSysDictItemByIdLogic {
	return &GetSysDictItemByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSysDictItemByIdLogic) GetSysDictItemById(req *types.DefaultIdRequest) (*types.SysDictItem, error) {
	res, err := l.svcCtx.SysDictItemRpc.GetSysDictItemById(l.ctx, &pb.GetSysDictItemByIdReq{
		Id: req.Id,
	})

	if err != nil {
		l.Logger.Errorf("根据ID: %d查询数据失败: %v", req.Id, err)
		return nil, errorx.DatabaseQueryErr
	}
	data := &types.SysDictItem{
		Id:          res.Data.Id,
		DictCode:    res.Data.DictCode,
		Description: res.Data.Description,
		ItemText:    res.Data.ItemText,
		ItemCode:    res.Data.ItemCode,
		SortOrder:   res.Data.SortOrder,
		CreatedBy:   res.Data.CreatedBy,
		UpdatedBy:   res.Data.UpdatedBy,
		UpdatedAt:   res.Data.UpdatedAt,
		CreatedAt:   res.Data.CreatedAt,
	}
	return data, nil
}
