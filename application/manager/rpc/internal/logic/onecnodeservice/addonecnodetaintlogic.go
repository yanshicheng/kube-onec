package onecnodeservicelogic

import (
	"context"

	"github.com/yanshicheng/kube-onec/application/manager/rpc/internal/svc"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddOnecNodeTaintLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddOnecNodeTaintLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddOnecNodeTaintLogic {
	return &AddOnecNodeTaintLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 添加污点
func (l *AddOnecNodeTaintLogic) AddOnecNodeTaint(in *pb.AddOnecNodeTaintReq) (*pb.AddOnecNodeTaintResp, error) {
	// todo: add your logic here and delete this line

	return &pb.AddOnecNodeTaintResp{}, nil
}
