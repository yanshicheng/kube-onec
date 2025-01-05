package sysroleservicelogic

import (
	"context"

	"github.com/yanshicheng/kube-onec/application/portal/rpc/internal/svc"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchRoleMenuLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchRoleMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchRoleMenuLogic {
	return &SearchRoleMenuLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchRoleMenuLogic) SearchRoleMenu(in *pb.SearchRoleMenuReq) (*pb.SearchRoleMenuResp, error) {
	// todo: add your logic here and delete this line

	return &pb.SearchRoleMenuResp{}, nil
}
