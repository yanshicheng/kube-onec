package sysmenuservicelogic

import (
	"context"

	"github.com/yanshicheng/kube-onec/application/portal/rpc/internal/svc"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateSysMenuLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateSysMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateSysMenuLogic {
	return &UpdateSysMenuLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateSysMenuLogic) UpdateSysMenu(in *pb.UpdateSysMenuReq) (*pb.UpdateSysMenuResp, error) {
	// todo: add your logic here and delete this line

	return &pb.UpdateSysMenuResp{}, nil
}
