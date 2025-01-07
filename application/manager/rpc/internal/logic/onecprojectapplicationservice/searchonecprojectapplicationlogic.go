package onecprojectapplicationservicelogic

import (
	"context"

	"github.com/yanshicheng/kube-onec/application/manager/rpc/internal/svc"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchOnecProjectApplicationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchOnecProjectApplicationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchOnecProjectApplicationLogic {
	return &SearchOnecProjectApplicationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchOnecProjectApplicationLogic) SearchOnecProjectApplication(in *pb.SearchOnecProjectApplicationReq) (*pb.SearchOnecProjectApplicationResp, error) {
	// todo: add your logic here and delete this line

	return &pb.SearchOnecProjectApplicationResp{}, nil
}
