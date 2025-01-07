package sysdictitemservicelogic

import (
	"context"
	"errors"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/internal/model"

	"github.com/yanshicheng/kube-onec/application/portal/rpc/internal/svc"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateSysDictItemLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateSysDictItemLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateSysDictItemLogic {
	return &UpdateSysDictItemLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateSysDictItemLogic) UpdateSysDictItem(in *pb.UpdateSysDictItemReq) (*pb.UpdateSysDictItemResp, error) {
	res, err := l.svcCtx.SysDictItem.FindOne(l.ctx, in.Id)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			l.Logger.Errorf("字典数据不存在: %v", err)
		}
		l.Logger.Errorf("根据ID: %d查询数据失败: %v", in.Id, err)
		return nil, err
	}
	if in.UpdatedBy != "" {
		res.UpdatedBy = in.UpdatedBy
	}
	if in.Description != "" {
		res.Description = in.Description
	}
	if in.ItemText != "" {
		res.ItemText = in.ItemText
	}
	if in.SortOrder != 0 {
		res.SortOrder = in.SortOrder
	}
	if err := l.svcCtx.SysDictItem.Update(l.ctx, res); err != nil {
		l.Logger.Errorf("更新字典数据失败: %v", err)
		return nil, err
	}
	return &pb.UpdateSysDictItemResp{}, nil
}
