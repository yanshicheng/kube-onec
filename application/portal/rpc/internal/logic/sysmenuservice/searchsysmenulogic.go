package sysmenuservicelogic

import (
	"context"

	"github.com/yanshicheng/kube-onec/application/portal/rpc/internal/svc"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchSysMenuLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchSysMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchSysMenuLogic {
	return &SearchSysMenuLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchSysMenuLogic) SearchSysMenu(in *pb.SearchSysMenuReq) (*pb.SearchSysMenuResp, error) {
	// todo: add your logic here and delete this line

	return &pb.SearchSysMenuResp{}, nil
}
