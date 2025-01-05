package sysroleservicelogic

import (
	"context"

	"github.com/yanshicheng/kube-onec/application/portal/rpc/internal/svc"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type BindRoleMenuLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBindRoleMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BindRoleMenuLogic {
	return &BindRoleMenuLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *BindRoleMenuLogic) BindRoleMenu(in *pb.BindRoleMenuReq) (*pb.BindRoleMenuResp, error) {
	// todo: add your logic here and delete this line

	return &pb.BindRoleMenuResp{}, nil
}
