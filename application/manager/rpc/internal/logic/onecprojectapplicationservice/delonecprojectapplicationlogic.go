package onecprojectapplicationservicelogic

import (
	"context"

	"github.com/yanshicheng/kube-onec/application/manager/rpc/internal/svc"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelOnecProjectApplicationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelOnecProjectApplicationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelOnecProjectApplicationLogic {
	return &DelOnecProjectApplicationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelOnecProjectApplicationLogic) DelOnecProjectApplication(in *pb.DelOnecProjectApplicationReq) (*pb.DelOnecProjectApplicationResp, error) {
	// todo: add your logic here and delete this line

	return &pb.DelOnecProjectApplicationResp{}, nil
}
