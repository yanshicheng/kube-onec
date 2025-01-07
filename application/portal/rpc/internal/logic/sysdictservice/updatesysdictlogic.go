package sysdictservicelogic

import (
	"context"
	"errors"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/internal/model"
	"github.com/yanshicheng/kube-onec/common/handler/errorx"

	"github.com/yanshicheng/kube-onec/application/portal/rpc/internal/svc"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateSysDictLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateSysDictLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateSysDictLogic {
	return &UpdateSysDictLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateSysDictLogic) UpdateSysDict(in *pb.UpdateSysDictReq) (*pb.UpdateSysDictResp, error) {
	res, err := l.svcCtx.SysDict.FindOne(l.ctx, in.Id)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			l.Logger.Errorf("数据字典未查询到: %v", err)
			return nil, errorx.DatabaseNotFound
		}
		l.Logger.Errorf("数据字典查询失败: %v", err)
		return nil, errorx.DatabaseQueryErr
	}
	if in.UpdatedBy != "" {
		res.UpdatedBy = in.UpdatedBy
	}
	if in.DictName != "" {
		res.DictName = in.DictName
	}
	if in.Description != "" {
		res.Description = in.Description
	}
	if err := l.svcCtx.SysDict.Update(l.ctx, res); err != nil {
		l.Logger.Errorf("数据字典更新失败: %v", err)
		return nil, errorx.DatabaseUpdateErr
	}
	return &pb.UpdateSysDictResp{}, nil
}
