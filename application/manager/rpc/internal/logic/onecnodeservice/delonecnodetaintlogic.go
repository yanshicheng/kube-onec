package onecnodeservicelogic

import (
	"context"

	"github.com/yanshicheng/kube-onec/application/manager/rpc/internal/svc"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelOnecNodeTaintLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelOnecNodeTaintLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelOnecNodeTaintLogic {
	return &DelOnecNodeTaintLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除污点
func (l *DelOnecNodeTaintLogic) DelOnecNodeTaint(in *pb.DelOnecNodeTaintReq) (*pb.DelOnecNodeTaintResp, error) {
	// todo: add your logic here and delete this line

	return &pb.DelOnecNodeTaintResp{}, nil
}
