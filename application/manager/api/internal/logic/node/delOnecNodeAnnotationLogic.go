package node

import (
	"context"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/client/onecnodeservice"
	"github.com/yanshicheng/kube-onec/utils"

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

	_, err = l.svcCtx.NodeRpc.DelOnecNodeAnnotation(l.ctx, &onecnodeservice.DelOnecNodeAnnotationReq{
		UpdatedBy:    utils.GetAccount(l.ctx),
		AnnotationId: req.Id,
	})

	if err != nil {
		l.Logger.Errorf("删除节点注解失败: %v", err)
		return
	}
	return "删除节点注解成功!", nil
}
