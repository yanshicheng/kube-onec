package cluster

import (
	"context"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/client/onecclusterservice"

	"github.com/yanshicheng/kube-onec/application/manager/api/internal/svc"
	"github.com/yanshicheng/kube-onec/application/manager/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelOnecClusterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDelOnecClusterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelOnecClusterLogic {
	return &DelOnecClusterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DelOnecClusterLogic) DelOnecCluster(req *types.DefaultIdRequest) (resp string, err error) {
	_, err = l.svcCtx.ClusterRpc.DelOnecCluster(l.ctx, &onecclusterservice.DelOnecClusterReq{
		Id: req.Id,
	})
	if err != nil {
		l.Logger.Errorf("删除集群失败，err: %v", err)
		return "", err
	}
	l.Logger.Infof("删除集群成功，req: %v", req)
	return "删除集群成功!", nil
}
