package cluster

import (
	"context"

	"github.com/yanshicheng/kube-onec/application/manager/api/internal/svc"
	"github.com/yanshicheng/kube-onec/application/manager/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SyncOnecClusterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSyncOnecClusterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SyncOnecClusterLogic {
	return &SyncOnecClusterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SyncOnecClusterLogic) SyncOnecCluster(req *types.SyncOnecClusterRequest) (resp string, err error) {
	// todo: add your logic here and delete this line

	return
}
