package onecprojectservicelogic

import (
	"context"

	"github.com/yanshicheng/kube-onec/application/manager/rpc/internal/svc"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelOnecProjectLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelOnecProjectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelOnecProjectLogic {
	return &DelOnecProjectLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelOnecProjectLogic) DelOnecProject(in *pb.DelOnecProjectReq) (*pb.DelOnecProjectResp, error) {
	// todo: add your logic here and delete this line

	return &pb.DelOnecProjectResp{}, nil
}
