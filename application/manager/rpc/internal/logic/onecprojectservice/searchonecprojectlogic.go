package onecprojectservicelogic

import (
	"context"

	"github.com/yanshicheng/kube-onec/application/manager/rpc/internal/svc"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchOnecProjectLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchOnecProjectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchOnecProjectLogic {
	return &SearchOnecProjectLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchOnecProjectLogic) SearchOnecProject(in *pb.SearchOnecProjectReq) (*pb.SearchOnecProjectResp, error) {
	// todo: add your logic here and delete this line

	return &pb.SearchOnecProjectResp{}, nil
}
