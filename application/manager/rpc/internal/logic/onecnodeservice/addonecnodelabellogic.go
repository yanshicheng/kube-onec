package onecnodeservicelogic

import (
	"context"

	"github.com/yanshicheng/kube-onec/application/manager/rpc/internal/svc"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddOnecNodeLabelLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddOnecNodeLabelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddOnecNodeLabelLogic {
	return &AddOnecNodeLabelLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 节点添加标签
func (l *AddOnecNodeLabelLogic) AddOnecNodeLabel(in *pb.AddOnecNodeLabelReq) (*pb.AddOnecNodeLabelResp, error) {
	// todo: add your logic here and delete this line

	return &pb.AddOnecNodeLabelResp{}, nil
}
