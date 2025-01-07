package onecprojectapplicationservicelogic

import (
	"context"

	"github.com/yanshicheng/kube-onec/application/manager/rpc/internal/svc"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddOnecProjectApplicationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddOnecProjectApplicationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddOnecProjectApplicationLogic {
	return &AddOnecProjectApplicationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// -----------------------应用表，-----------------------
func (l *AddOnecProjectApplicationLogic) AddOnecProjectApplication(in *pb.AddOnecProjectApplicationReq) (*pb.AddOnecProjectApplicationResp, error) {
	// todo: add your logic here and delete this line

	return &pb.AddOnecProjectApplicationResp{}, nil
}
