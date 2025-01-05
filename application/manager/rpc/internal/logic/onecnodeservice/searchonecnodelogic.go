package onecnodeservicelogic

import (
	"context"

	"github.com/yanshicheng/kube-onec/application/manager/rpc/internal/svc"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchOnecNodeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchOnecNodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchOnecNodeLogic {
	return &SearchOnecNodeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchOnecNodeLogic) SearchOnecNode(in *pb.SearchOnecNodeReq) (*pb.SearchOnecNodeResp, error) {
	// todo: add your logic here and delete this line

	return &pb.SearchOnecNodeResp{}, nil
}
