package onecnodeservicelogic

import (
	"context"

	"github.com/yanshicheng/kube-onec/application/manager/rpc/internal/svc"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelOnecNodeAnnotationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelOnecNodeAnnotationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelOnecNodeAnnotationLogic {
	return &DelOnecNodeAnnotationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 节点删除注解
func (l *DelOnecNodeAnnotationLogic) DelOnecNodeAnnotation(in *pb.DelOnecNodeAnnotationReq) (*pb.DelOnecNodeAnnotationResp, error) {
	// todo: add your logic here and delete this line

	return &pb.DelOnecNodeAnnotationResp{}, nil
}
