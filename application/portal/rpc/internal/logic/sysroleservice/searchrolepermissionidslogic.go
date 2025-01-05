package sysroleservicelogic

import (
	"context"

	"github.com/yanshicheng/kube-onec/application/portal/rpc/internal/svc"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchRolePermissionIdsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchRolePermissionIdsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchRolePermissionIdsLogic {
	return &SearchRolePermissionIdsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchRolePermissionIdsLogic) SearchRolePermissionIds(in *pb.SearchRolePermissionIdsReq) (*pb.SearchRolePermissionIdsResp, error) {
	// todo: add your logic here and delete this line

	return &pb.SearchRolePermissionIdsResp{}, nil
}
