package dict

import (
	"context"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/pb"
	"github.com/yanshicheng/kube-onec/common/handler/errorx"

	"github.com/yanshicheng/kube-onec/application/portal/api/internal/svc"
	"github.com/yanshicheng/kube-onec/application/portal/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSysDictByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetSysDictByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSysDictByIdLogic {
	return &GetSysDictByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSysDictByIdLogic) GetSysDictById(req *types.DefaultIdRequest) (resp *types.SysDict, err error) {
	res, err := l.svcCtx.SysDictRpc.GetSysDictById(l.ctx, &pb.GetSysDictByIdReq{
		Id: req.Id,
	})
	if err != nil {
		l.Logger.Errorf("数据字典查询失败: %v", err)
		return nil, errorx.DatabaseQueryErr
	}
	resp = &types.SysDict{
		CreateBy:    res.Data.CreateBy,
		CreateTime:  res.Data.CreateTime,
		Description: res.Data.Description,
		DictCode:    res.Data.DictCode,
		DictName:    res.Data.DictName,
		Id:          res.Data.Id,
		UpdateBy:    res.Data.UpdateBy,
		UpdateTime:  res.Data.UpdateTime,
	}
	return
}
