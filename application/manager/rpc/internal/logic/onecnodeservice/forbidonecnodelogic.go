package onecnodeservicelogic

import (
	"context"

	"github.com/yanshicheng/kube-onec/application/manager/rpc/internal/svc"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ForbidOnecNodeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewForbidOnecNodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ForbidOnecNodeLogic {
	return &ForbidOnecNodeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 禁止调度
func (l *ForbidOnecNodeLogic) ForbidOnecNode(in *pb.ForbidOnecNodeReq) (*pb.ForbidOnecNodeResp, error) {
	// todo: add your logic here and delete this line

	return &pb.ForbidOnecNodeResp{}, nil
}
