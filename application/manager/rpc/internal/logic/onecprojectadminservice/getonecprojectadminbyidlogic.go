package onecprojectadminservicelogic

import (
	"context"

	"github.com/yanshicheng/kube-onec/application/manager/rpc/internal/svc"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOnecProjectAdminByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetOnecProjectAdminByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOnecProjectAdminByIdLogic {
	return &GetOnecProjectAdminByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetOnecProjectAdminByIdLogic) GetOnecProjectAdminById(in *pb.GetOnecProjectAdminByIdReq) (*pb.GetOnecProjectAdminByIdResp, error) {
	// todo: add your logic here and delete this line

	return &pb.GetOnecProjectAdminByIdResp{}, nil
}
