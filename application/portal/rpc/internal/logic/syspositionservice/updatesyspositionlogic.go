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

type UpdateSysPositionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateSysPositionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateSysPositionLogic {
	return &UpdateSysPositionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateSysPositionLogic) UpdateSysPosition(in *pb.UpdateSysPositionReq) (*pb.UpdateSysPositionResp, error) {
	position, err := l.svcCtx.SysPosition.FindOne(l.ctx, in.Id)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return nil, errorx.DatabaseNotFound
		}
		l.Logger.Errorf("查询岗位失败: %v", err)
		return nil, errorx.DatabaseQueryErr
	}
	if in.Name != "" {
		position.Name = in.Name
	}
	err = l.svcCtx.SysPosition.Update(l.ctx, position)
	if err != nil {
		l.Logger.Errorf("更新岗位失败: %v", err)
		return nil, errorx.DatabaseUpdateErr
	}
	return &pb.UpdateSysPositionResp{}, nil
}
