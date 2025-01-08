package onecnodeservicelogic

import (
	"context"
	"errors"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/internal/code"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/internal/model"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/internal/shared"

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
	// 从数据库删除注解记录
	annotation, err := l.svcCtx.AnnotationsResourceModel.FindOne(l.ctx, in.AnnotationId)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			l.Logger.Infof("污点记录不存在: %v", err)
			return nil, code.NodeInfoNotExistErr
		}
		l.Logger.Errorf("查询节点注解数据库记录失败: %v", err)
		return nil, code.SearchNodeAnnotationDBErr
	}
	// 获取节点信息
	node, err := l.svcCtx.NodeModel.FindOne(l.ctx, annotation.ResourceId)
	if err != nil {
		l.Logger.Errorf("获取节点信息失败: %v", err)
		return nil, code.GetNodeInfoErr
	}

	// 获取 Kubernetes 客户端
	client, err := shared.GetK8sClient(l.ctx, l.svcCtx, node.ClusterUuid)
	if err != nil {
		l.Logger.Errorf("获取集群客户端异常: %v", err)
		return nil, code.GetClusterClientErr
	}

	// 删除节点注解
	err = client.GetNodeClient().RemoveAnnotation(node.NodeName, annotation.Key)
	if err != nil {
		l.Logger.Errorf("节点注解删除失败: %v", err)
		shared.ChangeNodeSyncStatus(l.ctx, l.svcCtx, *node, 0, in.UpdatedBy) // 更新节点同步状态为失败
		return nil, code.RemoveNodeAnnotationErr
	}

	if err := l.svcCtx.AnnotationsResourceModel.Delete(l.ctx, annotation.Id); err != nil {
		l.Logger.Errorf("删除节点注解数据库记录失败: %v", err)
		shared.ChangeNodeSyncStatus(l.ctx, l.svcCtx, *node, 0, in.UpdatedBy) // 更新节点同步状态为失败
		return nil, code.DeleteNodeAnnotationDBErr
	}

	// 更新同步状态为成功
	shared.ChangeNodeSyncStatus(l.ctx, l.svcCtx, *node, 1, in.UpdatedBy)
	l.Logger.Infof("节点注解删除成功: %v", node.NodeName)
	return &pb.DelOnecNodeAnnotationResp{}, nil
}
