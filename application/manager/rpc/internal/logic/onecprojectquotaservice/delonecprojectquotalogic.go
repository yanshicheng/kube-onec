package onecprojectquotaservicelogic

import (
	"context"

	"github.com/yanshicheng/kube-onec/application/manager/rpc/internal/svc"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelOnecProjectQuotaLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelOnecProjectQuotaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelOnecProjectQuotaLogic {
	return &DelOnecProjectQuotaLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelOnecProjectQuotaLogic) DelOnecProjectQuota(in *pb.DelOnecProjectQuotaReq) (*pb.DelOnecProjectQuotaResp, error) {
	// todo: add your logic here and delete this line

	return &pb.DelOnecProjectQuotaResp{}, nil
}
