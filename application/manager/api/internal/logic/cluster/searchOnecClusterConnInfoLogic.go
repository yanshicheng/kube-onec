package cluster

import (
	"context"

	"github.com/yanshicheng/kube-onec/application/manager/api/internal/svc"
	"github.com/yanshicheng/kube-onec/application/manager/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchOnecClusterConnInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchOnecClusterConnInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchOnecClusterConnInfoLogic {
	return &SearchOnecClusterConnInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchOnecClusterConnInfoLogic) SearchOnecClusterConnInfo(req *types.SearchOnecClusterConnInfoRequest) (resp *types.SearchOnecClusterConnInfoResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
