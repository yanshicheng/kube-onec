package syspositionservicelogic

import (
	"context"
	"errors"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/internal/model"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/internal/svc"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/pb"
	"github.com/yanshicheng/kube-onec/common/handler/errorx"
	"github.com/zeromicro/go-zero/core/logx"
)

type SearchSysPositionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchSysPositionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchSysPositionLogic {
	return &SearchSysPositionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchSysPositionLogic) SearchSysPosition(in *pb.SearchSysPositionReq) (*pb.SearchSysPositionResp, error) {
	queryStr := "name like ?"
	resp, total, err := l.svcCtx.SysPosition.Search(l.ctx, in.OrderStr, in.IsAsc, in.Page, in.PageSize, queryStr, "%"+in.Name+"%")
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return nil, errorx.DatabaseNotFound
		}
		l.Logger.Errorf("查询岗位失败: %v", err)
		return nil, errorx.DatabaseQueryErr
	}
	// 时间修改为时间戳
	var data []*pb.SysPosition

	for _, item := range resp {
		data = append(data, &pb.SysPosition{
			Id:         item.Id,
			Name:       item.Name,
			CreateTime: item.CreateTime.Unix(),
			UpdateTime: item.UpdateTime.Unix(),
		})
	}
	return &pb.SearchSysPositionResp{
		Data:  data,
		Total: total,
	}, nil
}
