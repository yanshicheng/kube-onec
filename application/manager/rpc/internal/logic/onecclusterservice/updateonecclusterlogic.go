package onecclusterservicelogic

import (
	"context"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/internal/svc"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateOnecClusterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateOnecClusterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateOnecClusterLogic {
	return &UpdateOnecClusterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateOnecClusterLogic) UpdateOnecCluster(in *pb.UpdateOnecClusterReq) (*pb.UpdateOnecClusterResp, error) {

	return &pb.UpdateOnecClusterResp{}, nil
}
