package onecprojectadminservicelogic

import (
	"context"

	"github.com/yanshicheng/kube-onec/application/manager/rpc/internal/svc"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelOnecProjectAdminLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelOnecProjectAdminLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelOnecProjectAdminLogic {
	return &DelOnecProjectAdminLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelOnecProjectAdminLogic) DelOnecProjectAdmin(in *pb.DelOnecProjectAdminReq) (*pb.DelOnecProjectAdminResp, error) {
	// todo: add your logic here and delete this line

	return &pb.DelOnecProjectAdminResp{}, nil
}
