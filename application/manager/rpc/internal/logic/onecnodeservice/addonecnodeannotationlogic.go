package onecnodeservicelogic

import (
	"context"

	"github.com/yanshicheng/kube-onec/application/manager/rpc/internal/svc"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddOnecNodeAnnotationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddOnecNodeAnnotationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddOnecNodeAnnotationLogic {
	return &AddOnecNodeAnnotationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 节点添加注解
func (l *AddOnecNodeAnnotationLogic) AddOnecNodeAnnotation(in *pb.AddOnecNodeAnnotationReq) (*pb.AddOnecNodeAnnotationResp, error) {
	// todo: add your logic here and delete this line

	return &pb.AddOnecNodeAnnotationResp{}, nil
}
