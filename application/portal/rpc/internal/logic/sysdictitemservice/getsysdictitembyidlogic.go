package sysdictitemservicelogic

import (
	"context"
	"errors"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/internal/model"
	"github.com/yanshicheng/kube-onec/common/handler/errorx"

	"github.com/yanshicheng/kube-onec/application/portal/rpc/internal/svc"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSysDictItemByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetSysDictItemByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSysDictItemByIdLogic {
	return &GetSysDictItemByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetSysDictItemByIdLogic) GetSysDictItemById(in *pb.GetSysDictItemByIdReq) (*pb.GetSysDictItemByIdResp, error) {

	res, err := l.svcCtx.SysDictItem.FindOne(l.ctx, in.Id)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			l.Logger.Errorf("字典数据不存在: %v", err)
			return nil, errorx.DatabaseNotFound
		}
		l.Logger.Errorf("根据ID: %d查询数据失败: %v", in.Id, err)
		return nil, errorx.DatabaseQueryErr
	}
	data := &pb.SysDictItem{
		Id:          res.Id,
		DictCode:    res.DictCode,
		Description: res.Description,
		ItemText:    res.ItemText,
		ItemCode:    res.ItemCode,
		SortOrder:   res.SortOrder,
		CreateBy:    res.CreateBy,
		UpdateBy:    res.UpdateBy,
		UpdateTime:  res.UpdateTime.Unix(),
		CreateTime:  res.CreateTime.Unix(),
	}
	return &pb.GetSysDictItemByIdResp{
		Data: data,
	}, nil
}
