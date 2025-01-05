package syspositionservicelogic

import (
	"context"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/internal/model"
	"github.com/yanshicheng/kube-onec/common/handler/errorx"

	"github.com/yanshicheng/kube-onec/application/portal/rpc/internal/svc"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddSysPositionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddSysPositionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddSysPositionLogic {
	return &AddSysPositionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// -----------------------职位表-----------------------
func (l *AddSysPositionLogic) AddSysPosition(in *pb.AddSysPositionReq) (*pb.AddSysPositionResp, error) {
	_, err := l.svcCtx.SysPosition.Insert(l.ctx, &model.SysPosition{
		Name: in.Name,
	})
	if err != nil {
		l.Logger.Errorf("添加职位失败: %v", err)
		return nil, errorx.DatabaseCreateErr
	}
	return &pb.AddSysPositionResp{}, nil
}
