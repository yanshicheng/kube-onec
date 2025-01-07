package onecprojectapplicationservicelogic

import (
	"context"

	"github.com/yanshicheng/kube-onec/application/manager/rpc/internal/svc"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOnecProjectApplicationByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetOnecProjectApplicationByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOnecProjectApplicationByIdLogic {
	return &GetOnecProjectApplicationByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetOnecProjectApplicationByIdLogic) GetOnecProjectApplicationById(in *pb.GetOnecProjectApplicationByIdReq) (*pb.GetOnecProjectApplicationByIdResp, error) {
	// todo: add your logic here and delete this line

	return &pb.GetOnecProjectApplicationByIdResp{}, nil
}
