package onecnodeservicelogic

import (
	"context"

	"github.com/yanshicheng/kube-onec/application/manager/rpc/internal/svc"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelOnecNodeLabelLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelOnecNodeLabelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelOnecNodeLabelLogic {
	return &DelOnecNodeLabelLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 节点删除标签
func (l *DelOnecNodeLabelLogic) DelOnecNodeLabel(in *pb.DelOnecNodeLabelReq) (*pb.DelOnecNodeLabelResp, error) {
	// todo: add your logic here and delete this line

	return &pb.DelOnecNodeLabelResp{}, nil
}
