package sysmenuservicelogic

import (
	"context"

	"github.com/yanshicheng/kube-onec/application/portal/rpc/internal/svc"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddSysMenuLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddSysMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddSysMenuLogic {
	return &AddSysMenuLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// -----------------------菜单表-----------------------
func (l *AddSysMenuLogic) AddSysMenu(in *pb.AddSysMenuReq) (*pb.AddSysMenuResp, error) {
	// todo: add your logic here and delete this line

	return &pb.AddSysMenuResp{}, nil
}
