package cluster

import (
	"context"
	"github.com/yanshicheng/kube-onec/application/manager/api/internal/svc"
	"github.com/yanshicheng/kube-onec/application/manager/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddOnecClusterConnInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddOnecClusterConnInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddOnecClusterConnInfoLogic {
	return &AddOnecClusterConnInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddOnecClusterConnInfoLogic) AddOnecClusterConnInfo(req *types.AddOnecClusterConnInfoRequest) (resp string, err error) {

	// 如果需要返回成功的响应，可以自行处理 res
	return "集群连接信息添加成功", nil
}
