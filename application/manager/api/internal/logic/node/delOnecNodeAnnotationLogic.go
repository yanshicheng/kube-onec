package node

import (
	"context"

	"github.com/yanshicheng/kube-onec/application/manager/api/internal/svc"
	"github.com/yanshicheng/kube-onec/application/manager/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelOnecNodeAnnotationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDelOnecNodeAnnotationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelOnecNodeAnnotationLogic {
	return &DelOnecNodeAnnotationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DelOnecNodeAnnotationLogic) DelOnecNodeAnnotation(req *types.DelOnecNodeAnnotationRequest) (resp string, err error) {
	// todo: add your logic here and delete this line

	return
}
