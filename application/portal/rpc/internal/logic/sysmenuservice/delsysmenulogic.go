package sysmenuservicelogic

import (
	"context"

	"github.com/yanshicheng/kube-onec/application/portal/rpc/internal/svc"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelSysMenuLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelSysMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelSysMenuLogic {
	return &DelSysMenuLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelSysMenuLogic) DelSysMenu(in *pb.DelSysMenuReq) (*pb.DelSysMenuResp, error) {
	// todo: add your logic here and delete this line

	return &pb.DelSysMenuResp{}, nil
}
