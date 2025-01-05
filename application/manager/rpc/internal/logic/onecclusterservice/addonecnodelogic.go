package onecclusterservicelogic

import (
	"context"

	"github.com/yanshicheng/kube-onec/application/manager/rpc/internal/svc"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddOnecNodeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddOnecNodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddOnecNodeLogic {
	return &AddOnecNodeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddOnecNodeLogic) AddOnecNode(in *pb.AddOnecNodeReq) (*pb.AddOnecNodeResp, error) {
	// todo: add your logic here and delete this line

	return &pb.AddOnecNodeResp{}, nil
}
