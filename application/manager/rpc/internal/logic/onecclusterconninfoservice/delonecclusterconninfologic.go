package onecclusterconninfoservicelogic

import (
	"context"

	"github.com/yanshicheng/kube-onec/application/manager/rpc/internal/svc"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelOnecClusterConnInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelOnecClusterConnInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelOnecClusterConnInfoLogic {
	return &DelOnecClusterConnInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelOnecClusterConnInfoLogic) DelOnecClusterConnInfo(in *pb.DelOnecClusterConnInfoReq) (*pb.DelOnecClusterConnInfoResp, error) {
	// todo: add your logic here and delete this line

	return &pb.DelOnecClusterConnInfoResp{}, nil
}
