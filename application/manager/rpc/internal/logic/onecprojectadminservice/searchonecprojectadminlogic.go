package onecprojectadminservicelogic

import (
	"context"

	"github.com/yanshicheng/kube-onec/application/manager/rpc/internal/svc"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchOnecProjectAdminLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchOnecProjectAdminLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchOnecProjectAdminLogic {
	return &SearchOnecProjectAdminLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchOnecProjectAdminLogic) SearchOnecProjectAdmin(in *pb.SearchOnecProjectAdminReq) (*pb.SearchOnecProjectAdminResp, error) {
	// todo: add your logic here and delete this line

	return &pb.SearchOnecProjectAdminResp{}, nil
}
