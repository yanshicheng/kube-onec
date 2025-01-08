package onecnodeservicelogic

import (
	"context"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/internal/code"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/internal/model"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/internal/shared"

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
// 节点添加注解
func (l *AddOnecNodeAnnotationLogic) AddOnecNodeAnnotation(in *pb.AddOnecNodeAnnotationReq) (*pb.AddOnecNodeAnnotationResp, error) {

	// 获取节点信息
	node, err := l.svcCtx.NodeModel.FindOne(l.ctx, in.NodeId)
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

	// 添加节点注解
	err = client.GetNodeClient().AddAnnotation(node.NodeName, in.Key, in.Value)
	if err != nil {
		l.Logger.Errorf("节点注解数据同步失败: %v", err)
		shared.ChangeNodeSyncStatus(l.ctx, l.svcCtx, *node, 0, in.UpdatedBy) // 更新节点同步状态为失败
		return nil, code.SyncNodeAnnotationErr
	}

	// 同步节点注解到数据库
	_, err = l.svcCtx.AnnotationsResourceModel.Insert(l.ctx, &model.OnecResourceAnnotations{
		Key:          in.Key,
		Value:        in.Value,
		ResourceType: "node",
		ResourceId:   node.Id,
	})
	if err != nil {
		l.Logger.Errorf("节点注解数据同步失败: %v", err)
		shared.ChangeNodeSyncStatus(l.ctx, l.svcCtx, *node, 0, in.UpdatedBy) // 更新节点同步状态为失败
		return nil, code.SyncNodeAnnotationErr
	}

	// 更新同步状态为成功
	shared.ChangeNodeSyncStatus(l.ctx, l.svcCtx, *node, 1, in.UpdatedBy)
	l.Logger.Infof("节点注解数据同步成功: %v", node.NodeName)
	return &pb.AddOnecNodeAnnotationResp{}, nil
}
