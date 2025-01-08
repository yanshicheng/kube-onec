package onecnodeservicelogic

import (
	"context"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/internal/code"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/internal/model"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/internal/shared"
	portalPb "github.com/yanshicheng/kube-onec/application/portal/rpc/pb"
	corev1 "k8s.io/api/core/v1"

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
// 添加污点
func (l *AddOnecNodeTaintLogic) AddOnecNodeTaint(in *pb.AddOnecNodeTaintReq) (*pb.AddOnecNodeTaintResp, error) {
	// 验证 effect 是否存在
	if _, err := l.svcCtx.SysDictItemRpc.CheckDictItemCode(l.ctx, &portalPb.CheckDictItemCodeReq{DictCode: shared.DictTaintEffectCode, ItemCode: in.Effect}); err != nil {
		l.Logger.Errorf("字典项不存在: %v", err)
		return nil, code.DictItemNotExistErr
	}

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

	// 构造 Taint 对象
	taint := corev1.Taint{
		Key:    in.Key,
		Value:  in.Value,
		Effect: corev1.TaintEffect(in.Effect), // TaintEffect 枚举值
	}

	// 添加污点
	err = client.GetNodeClient().AddTaint(node.NodeName, taint)
	if err != nil {
		l.Logger.Errorf("节点污点添加失败: %v", err)
		shared.ChangeNodeSyncStatus(l.ctx, l.svcCtx, *node, 0, in.UpdatedBy) // 更新节点同步状态为失败
		return nil, code.AddNodeTaintErr
	}

	// 同步污点到数据库
	_, err = l.svcCtx.TaintsResourceModel.Insert(l.ctx, &model.OnecResourceTaints{
		Key:        in.Key,
		Value:      in.Value,
		EffectCode: in.Effect,
		NodeId:     node.Id,
	})
	if err != nil {
		l.Logger.Errorf("污点数据同步到数据库失败: %v", err)
		shared.ChangeNodeSyncStatus(l.ctx, l.svcCtx, *node, 0, in.UpdatedBy) // 更新节点同步状态为失败
		return nil, code.SyncNodeTaintDBErr
	}

	// 更新同步状态为成功
	shared.ChangeNodeSyncStatus(l.ctx, l.svcCtx, *node, 1, in.UpdatedBy)
	l.Logger.Infof("节点污点添加成功: %v", node.NodeName)
	return &pb.AddOnecNodeTaintResp{}, nil
}
