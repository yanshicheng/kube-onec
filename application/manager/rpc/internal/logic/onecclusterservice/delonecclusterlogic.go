package onecclusterservicelogic

import (
	"context"

	"github.com/yanshicheng/kube-onec/application/manager/rpc/internal/svc"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelOnecClusterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelOnecClusterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelOnecClusterLogic {
	return &DelOnecClusterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelOnecClusterLogic) DelOnecCluster(in *pb.DelOnecClusterReq) (*pb.DelOnecClusterResp, error) {
	// todo: add your logic here and delete this line

	return &pb.DelOnecClusterResp{}, nil
}
