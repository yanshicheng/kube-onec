package node

import (
	"context"
	"github.com/yanshicheng/kube-onec/application/manager/api/internal/svc"
	"github.com/yanshicheng/kube-onec/application/manager/api/internal/types"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/pb"
	"github.com/yanshicheng/kube-onec/pkg/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelOnecNodeLabelLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDelOnecNodeLabelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelOnecNodeLabelLogic {
	return &DelOnecNodeLabelLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DelOnecNodeLabelLogic) DelOnecNodeLabel(req *types.DelOnecNodeLabelRequest) (resp string, err error) {

	_, err = l.svcCtx.NodeRpc.DelOnecNodeLabel(l.ctx, &pb.DelOnecNodeLabelReq{
		LabelId:   req.Id,
		UpdatedBy: utils.GetAccount(l.ctx),
	})
	if err != nil {
		l.Logger.Errorf("删除节点标签失败: %v", err)
		return
	}
	l.Logger.Infof("删除节点标签成功: %v", req)
	return "删除节点标签成功!", nil
}
