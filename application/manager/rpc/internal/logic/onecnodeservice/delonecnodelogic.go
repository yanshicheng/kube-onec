package onecnodeservicelogic

import (
	"context"

	"github.com/yanshicheng/kube-onec/application/manager/rpc/internal/svc"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelOnecNodeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelOnecNodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelOnecNodeLogic {
	return &DelOnecNodeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// -----------------------节点表，用于管理各集群中的节点信息-----------------------
func (l *DelOnecNodeLogic) DelOnecNode(in *pb.DelOnecNodeReq) (*pb.DelOnecNodeResp, error) {
	// todo: add your logic here and delete this line

	return &pb.DelOnecNodeResp{}, nil
}
