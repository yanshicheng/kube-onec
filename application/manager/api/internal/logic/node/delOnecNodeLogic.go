package node

import (
	"context"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/pb"

	"github.com/yanshicheng/kube-onec/application/manager/api/internal/svc"
	"github.com/yanshicheng/kube-onec/application/manager/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelOnecNodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDelOnecNodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelOnecNodeLogic {
	return &DelOnecNodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DelOnecNodeLogic) DelOnecNode(req *types.DelOnecNodeRequest) (resp string, err error) {
	_, err = l.svcCtx.NodeRpc.DelOnecNode(l.ctx, &pb.DelOnecNodeReq{
		Id:          req.Id,
		ClusterUuid: req.ClusterUuid,
	})
	if err != nil {
		l.Logger.Errorf("删除节点失败: %v", err)
		return
	}
	l.Logger.Infof("删除节点成功: %v", req)
	return "删除节点成功!", nil
}
