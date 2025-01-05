package syspositionservicelogic

import (
	"context"

	"github.com/yanshicheng/kube-onec/application/portal/rpc/internal/svc"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchSysPositionByUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchSysPositionByUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchSysPositionByUserLogic {
	return &SearchSysPositionByUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchSysPositionByUserLogic) SearchSysPositionByUser(in *pb.SearchSysPositionByUserReq) (*pb.SearchSysPositionByUserResp, error) {
	// todo: add your logic here and delete this line

	return &pb.SearchSysPositionByUserResp{}, nil
}
