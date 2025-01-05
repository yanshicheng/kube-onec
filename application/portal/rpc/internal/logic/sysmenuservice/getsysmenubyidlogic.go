package sysmenuservicelogic

import (
	"context"

	"github.com/yanshicheng/kube-onec/application/portal/rpc/internal/svc"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSysMenuByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetSysMenuByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSysMenuByIdLogic {
	return &GetSysMenuByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetSysMenuByIdLogic) GetSysMenuById(in *pb.GetSysMenuByIdReq) (*pb.GetSysMenuByIdResp, error) {
	// todo: add your logic here and delete this line

	return &pb.GetSysMenuByIdResp{}, nil
}
