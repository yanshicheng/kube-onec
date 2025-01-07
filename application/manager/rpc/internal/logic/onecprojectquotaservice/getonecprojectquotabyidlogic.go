package onecprojectquotaservicelogic

import (
	"context"

	"github.com/yanshicheng/kube-onec/application/manager/rpc/internal/svc"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOnecProjectQuotaByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetOnecProjectQuotaByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOnecProjectQuotaByIdLogic {
	return &GetOnecProjectQuotaByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetOnecProjectQuotaByIdLogic) GetOnecProjectQuotaById(in *pb.GetOnecProjectQuotaByIdReq) (*pb.GetOnecProjectQuotaByIdResp, error) {
	// todo: add your logic here and delete this line

	return &pb.GetOnecProjectQuotaByIdResp{}, nil
}
