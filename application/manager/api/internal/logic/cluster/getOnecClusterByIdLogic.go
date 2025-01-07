package cluster

import (
	"context"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/client/onecclusterservice"
	"github.com/yanshicheng/kube-onec/common/handler/errorx"

	"github.com/yanshicheng/kube-onec/application/manager/api/internal/svc"
	"github.com/yanshicheng/kube-onec/application/manager/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOnecClusterByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetOnecClusterByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOnecClusterByIdLogic {
	return &GetOnecClusterByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetOnecClusterByIdLogic) GetOnecClusterById(req *types.DefaultIdRequest) (resp *types.OnecCluster, err error) {
	res, err := l.svcCtx.ClusterRpc.GetOnecClusterById(l.ctx, &onecclusterservice.GetOnecClusterByIdReq{
		Id: req.Id,
	})
	if err != nil {
		l.Logger.Errorf("获取集群信息失败: %v", err)
		return nil, errorx.DatabaseFindErr
	}

	// 将返回的 RPC 数据映射到 API 响应结构中
	data := convertCluster(res.Data)
	return &data, nil
}
