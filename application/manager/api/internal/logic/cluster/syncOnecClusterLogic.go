package cluster

import (
	"context"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/client/onecclusterservice"

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

func (l *SyncOnecClusterLogic) SyncOnecCluster(req *types.DefaultIdRequest) (resp string, err error) {
	account, ok := l.ctx.Value("account").(string)
	if !ok || account == "" {
		account = "system"
	}
	_, err = l.svcCtx.ClusterRpc.SyncOnecCluster(l.ctx, &onecclusterservice.SyncOnecClusterReq{
		Id:        req.Id,
		UpdatedBy: account,
	})
	if err != nil {
		l.Logger.Errorf("同步集群失败，err: %v", err)
		return "", err
	}
	l.Logger.Infof("同步集群成功，req: %v", req)
	return "集群同步成功!", nil
}
