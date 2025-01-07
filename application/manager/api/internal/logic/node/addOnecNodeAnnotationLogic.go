package node

import (
	"context"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/client/onecnodeservice"

	"github.com/yanshicheng/kube-onec/application/manager/api/internal/svc"
	"github.com/yanshicheng/kube-onec/application/manager/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddOnecNodeAnnotationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddOnecNodeAnnotationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddOnecNodeAnnotationLogic {
	return &AddOnecNodeAnnotationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddOnecNodeAnnotationLogic) AddOnecNodeAnnotation(req *types.AddOnecNodeAnnotationRequest) (resp string, err error) {
	_, err = l.svcCtx.NodeRpc.AddOnecNodeAnnotation(l.ctx, &onecnodeservice.AddOnecNodeAnnotationReq{
		ClusterUuid: req.ClusterUuid,
		Id:          req.Id,
		Key:         req.Key,
		Value:       req.Value,
	})
	if err != nil {
		l.Logger.Errorf("添加节点注解失败: %v", err)
		return
	}
	l.Logger.Infof("添加节点注解成功: %v", req)
	return "添加节点注解成功!", nil
}
