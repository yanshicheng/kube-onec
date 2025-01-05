package onecnodeservicelogic

import (
	"context"

	"github.com/yanshicheng/kube-onec/application/manager/rpc/internal/svc"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CancelForbidOnecNodeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCancelForbidOnecNodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CancelForbidOnecNodeLogic {
	return &CancelForbidOnecNodeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 取消禁止调度
func (l *CancelForbidOnecNodeLogic) CancelForbidOnecNode(in *pb.CancelForbidOnecNodeReq) (*pb.CancelForbidOnecNodeResp, error) {
	// todo: add your logic here and delete this line

	return &pb.CancelForbidOnecNodeResp{}, nil
}
