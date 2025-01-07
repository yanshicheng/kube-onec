package onecprojectservicelogic

import (
	"context"

	"github.com/yanshicheng/kube-onec/application/manager/rpc/internal/svc"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOnecProjectByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetOnecProjectByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOnecProjectByIdLogic {
	return &GetOnecProjectByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetOnecProjectByIdLogic) GetOnecProjectById(in *pb.GetOnecProjectByIdReq) (*pb.GetOnecProjectByIdResp, error) {
	// todo: add your logic here and delete this line

	return &pb.GetOnecProjectByIdResp{}, nil
}
