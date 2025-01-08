package onecnodeservicelogic

import (
	"context"
	"errors"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/internal/code"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/internal/model"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/internal/shared"
	corev1 "k8s.io/api/core/v1"

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
	// 从数据库删除污点记录
	taintRecord, err := l.svcCtx.TaintsResourceModel.FindOne(l.ctx, in.TaintId)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			l.Logger.Infof("污点记录不存在: %v", err)
			return nil, code.NodeInfoNotExistErr
		}
		l.Logger.Errorf("查询污点数据库记录失败: %v", err)
		return nil, code.SearchNodeTaintDBErr
	}
	// 获取节点信息
	node, err := l.svcCtx.NodeModel.FindOne(l.ctx, taintRecord.NodeId)
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

	// 验证 effect 是否存在
	//if _, err := l.svcCtx.SysDictItemRpc.CheckDictItemCode(l.ctx, &portalPb.CheckDictItemCodeReq{DictCode: shared.DictTaintEffectCode, ItemCode: taintRecord.EffectCode}); err != nil {
	//	l.Logger.Errorf("字典项不存在: %v", err)
	//	return nil, code.DictItemNotExistErr
	//}

	// 构造 Taint 对象
	taint := corev1.Taint{
		Key:    taintRecord.Key,
		Value:  taintRecord.Value,
		Effect: corev1.TaintEffect(taintRecord.EffectCode),
	}

	// 删除污点
	err = client.GetNodeClient().RemoveTaint(node.NodeName, taint)
	if err != nil {
		l.Logger.Errorf("节点污点删除失败: %v", err)
		shared.ChangeNodeSyncStatus(l.ctx, l.svcCtx, *node, 0, in.UpdatedBy) // 更新节点同步状态为失败
		return nil, code.RemoveNodeTaintErr
	}

	if err := l.svcCtx.TaintsResourceModel.Delete(l.ctx, taintRecord.Id); err != nil {
		l.Logger.Errorf("删除污点数据库记录失败: %v", err)
		shared.ChangeNodeSyncStatus(l.ctx, l.svcCtx, *node, 0, in.UpdatedBy) // 更新节点同步状态为失败
		return nil, code.DeleteNodeTaintDBErr
	}

	// 更新同步状态为成功
	shared.ChangeNodeSyncStatus(l.ctx, l.svcCtx, *node, 1, in.UpdatedBy)
	l.Logger.Infof("节点污点删除成功: %v", node.NodeName)
	return &pb.DelOnecNodeTaintResp{}, nil
}
