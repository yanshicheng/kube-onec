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

type GetSysDictByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetSysDictByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSysDictByIdLogic {
	return &GetSysDictByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetSysDictByIdLogic) GetSysDictById(in *pb.GetSysDictByIdReq) (*pb.GetSysDictByIdResp, error) {
	res, err := l.svcCtx.SysDict.FindOne(l.ctx, in.Id)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			l.Logger.Errorf("数据字典未查询到: %v", err)
			return nil, errorx.DatabaseNotFound
		}
		l.Logger.Errorf("数据字典查询失败: %v", err)
		return nil, errorx.DatabaseQueryErr
	}

	return &pb.GetSysDictByIdResp{
		Data: &pb.SysDict{
			Id:          res.Id,
			DictName:    res.DictName,
			DictCode:    res.DictCode,
			Description: res.Description,
			UpdatedBy:   res.UpdatedBy,
			UpdatedAt:   res.UpdatedAt.Unix(),
			CreatedBy:   res.CreatedBy,
			CreatedAt:   res.CreatedAt.Unix(),
		},
	}, nil
}
