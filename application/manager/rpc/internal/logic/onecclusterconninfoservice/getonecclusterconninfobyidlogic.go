package onecclusterconninfoservicelogic

import (
	"context"

	"github.com/yanshicheng/kube-onec/application/manager/rpc/internal/svc"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOnecClusterConnInfoByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetOnecClusterConnInfoByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOnecClusterConnInfoByIdLogic {
	return &GetOnecClusterConnInfoByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetOnecClusterConnInfoByIdLogic) GetOnecClusterConnInfoById(in *pb.GetOnecClusterConnInfoByIdReq) (*pb.GetOnecClusterConnInfoByIdResp, error) {
	// todo: add your logic here and delete this line

	return &pb.GetOnecClusterConnInfoByIdResp{}, nil
}
