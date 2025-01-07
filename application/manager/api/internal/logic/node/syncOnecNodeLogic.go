package node

import (
	"context"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/pb"

	"github.com/yanshicheng/kube-onec/application/manager/api/internal/svc"
	"github.com/yanshicheng/kube-onec/application/manager/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SyncOnecNodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSyncOnecNodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SyncOnecNodeLogic {
	return &SyncOnecNodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SyncOnecNodeLogic) SyncOnecNode(req *types.SyncOnecNodeRequest) (resp string, err error) {
	account, ok := l.ctx.Value("account").(string)
	if !ok {
		account = "system"
	}

	_, err = l.svcCtx.NodeRpc.SyncOnecNode(l.ctx, &pb.SyncOnecNodeReq{
		ClusterUuid: req.ClusterUuid,
		Id:          req.Id,
		UpdatedBy:    account,
	})
	if err != nil {
		l.Logger.Errorf("同步节点失败: %v", err)
		return
	}
	l.Logger.Infof("同步节点成功: %v", req)
	return "同步节点成功!", nil
}
