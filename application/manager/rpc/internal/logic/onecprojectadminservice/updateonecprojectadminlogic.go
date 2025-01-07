package onecprojectadminservicelogic

import (
	"context"

	"github.com/yanshicheng/kube-onec/application/manager/rpc/internal/svc"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateOnecProjectAdminLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateOnecProjectAdminLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateOnecProjectAdminLogic {
	return &UpdateOnecProjectAdminLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateOnecProjectAdminLogic) UpdateOnecProjectAdmin(in *pb.UpdateOnecProjectAdminReq) (*pb.UpdateOnecProjectAdminResp, error) {
	// todo: add your logic here and delete this line

	return &pb.UpdateOnecProjectAdminResp{}, nil
}
