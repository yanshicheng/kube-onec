package onecprojectservicelogic

import (
	"context"

	"github.com/yanshicheng/kube-onec/application/manager/rpc/internal/svc"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateOnecProjectLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateOnecProjectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateOnecProjectLogic {
	return &UpdateOnecProjectLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateOnecProjectLogic) UpdateOnecProject(in *pb.UpdateOnecProjectReq) (*pb.UpdateOnecProjectResp, error) {
	// todo: add your logic here and delete this line

	return &pb.UpdateOnecProjectResp{}, nil
}
