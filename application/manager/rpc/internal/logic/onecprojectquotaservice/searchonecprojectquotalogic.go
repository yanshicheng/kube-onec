package onecprojectquotaservicelogic

import (
	"context"

	"github.com/yanshicheng/kube-onec/application/manager/rpc/internal/svc"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchOnecProjectQuotaLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchOnecProjectQuotaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchOnecProjectQuotaLogic {
	return &SearchOnecProjectQuotaLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchOnecProjectQuotaLogic) SearchOnecProjectQuota(in *pb.SearchOnecProjectQuotaReq) (*pb.SearchOnecProjectQuotaResp, error) {
	// todo: add your logic here and delete this line

	return &pb.SearchOnecProjectQuotaResp{}, nil
}
