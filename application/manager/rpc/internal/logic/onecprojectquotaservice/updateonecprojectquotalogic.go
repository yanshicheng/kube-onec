package onecprojectquotaservicelogic

import (
	"context"

	"github.com/yanshicheng/kube-onec/application/manager/rpc/internal/svc"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateOnecProjectQuotaLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateOnecProjectQuotaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateOnecProjectQuotaLogic {
	return &UpdateOnecProjectQuotaLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateOnecProjectQuotaLogic) UpdateOnecProjectQuota(in *pb.UpdateOnecProjectQuotaReq) (*pb.UpdateOnecProjectQuotaResp, error) {
	// todo: add your logic here and delete this line

	return &pb.UpdateOnecProjectQuotaResp{}, nil
}
