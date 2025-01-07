package onecclusterconninfoservicelogic

import (
	"context"

	"github.com/yanshicheng/kube-onec/application/manager/rpc/internal/svc"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateOnecClusterConnInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateOnecClusterConnInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateOnecClusterConnInfoLogic {
	return &UpdateOnecClusterConnInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateOnecClusterConnInfoLogic) UpdateOnecClusterConnInfo(in *pb.UpdateOnecClusterConnInfoReq) (*pb.UpdateOnecClusterConnInfoResp, error) {
	// todo: add your logic here and delete this line

	return &pb.UpdateOnecClusterConnInfoResp{}, nil
}
