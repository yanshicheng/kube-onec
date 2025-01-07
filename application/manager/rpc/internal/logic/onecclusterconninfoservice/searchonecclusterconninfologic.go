package onecclusterconninfoservicelogic

import (
	"context"

	"github.com/yanshicheng/kube-onec/application/manager/rpc/internal/svc"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchOnecClusterConnInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchOnecClusterConnInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchOnecClusterConnInfoLogic {
	return &SearchOnecClusterConnInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchOnecClusterConnInfoLogic) SearchOnecClusterConnInfo(in *pb.SearchOnecClusterConnInfoReq) (*pb.SearchOnecClusterConnInfoResp, error) {
	// todo: add your logic here and delete this line

	return &pb.SearchOnecClusterConnInfoResp{}, nil
}
