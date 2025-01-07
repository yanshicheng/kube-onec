package onecprojectapplicationservicelogic

import (
	"context"

	"github.com/yanshicheng/kube-onec/application/manager/rpc/internal/svc"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateOnecProjectApplicationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateOnecProjectApplicationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateOnecProjectApplicationLogic {
	return &UpdateOnecProjectApplicationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateOnecProjectApplicationLogic) UpdateOnecProjectApplication(in *pb.UpdateOnecProjectApplicationReq) (*pb.UpdateOnecProjectApplicationResp, error) {
	// todo: add your logic here and delete this line

	return &pb.UpdateOnecProjectApplicationResp{}, nil
}
