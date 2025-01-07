package onecnodeservicelogic

import (
	"context"
	"github.com/yanshicheng/kube-onec/common/handler/errorx"
	"sync"

	"github.com/yanshicheng/kube-onec/application/manager/rpc/internal/code"
	onecclusterservicelogic "github.com/yanshicheng/kube-onec/application/manager/rpc/internal/logic/onecclusterservice"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/internal/model"
	"github.com/yanshicheng/kube-onec/common/k8swrapper/core"
	"github.com/yanshicheng/kube-onec/utils"

	"github.com/yanshicheng/kube-onec/application/manager/rpc/internal/svc"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SyncOnecNodeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSyncOnecNodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SyncOnecNodeLogic {
	return &SyncOnecNodeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 同步节点信息
func (l *SyncOnecNodeLogic) SyncOnecNode(in *pb.SyncOnecNodeReq) (*pb.SyncOnecNodeResp, error) {
	// 获取客户端
	client, err := l.svcCtx.OnecClient.GetOrCreateOnecK8sClient(l.ctx, in.ClusterUuid, nil)
	if err != nil {
		l.Logger.Infof("获取集群客户端失败: %v", err)
		cluster, err := l.svcCtx.ClusterModel.FindOneByUuid(l.ctx, in.ClusterUuid)
		if err != nil {
			l.Logger.Errorf("获取集群信息失败: %v", err)
			return nil, code.GetClusterInfoErr
		}
		client, err = l.svcCtx.OnecClient.GetOrCreateOnecK8sClient(l.ctx, in.ClusterUuid, utils.NewRestConfig(cluster.Host, cluster.Token, utils.IntToBool(cluster.SkipInsecure)))
		if err != nil {
			l.Logger.Infof("获取集群客户端失败: %v", err)
			return nil, code.GetClusterClientErr
		}
	}

	if err := client.Ping(); err != nil {
		l.Logger.Infof("集群: %v, 连接失败: %v", in.ClusterUuid, err)
		return nil, code.ClusterConnectErr
	}

	// 查询节点数据
	node, err := l.svcCtx.NodeModel.FindOne(l.ctx, in.Id)
	if err != nil {
		l.Logger.Errorf("获取节点信息失败: %v, 集群: %v nodes: %v", err, in.ClusterUuid, in.Id)
		return nil, code.GetNodeInfoErr
	}
	l.Logger.Infof("集群Id: %v, nodes: %v 正在更新", in.ClusterUuid, node.NodeName)

	// 获取节点信息
	nodeInfo, err := client.GetNodes().GetNodeInfo(node.NodeName)
	if err != nil {
		l.Logger.Errorf("获取节点信息失败: %v, 集群: %v nodes: %v", err, in.ClusterUuid, node.NodeName)
		return nil, code.GetNodeInfoErr
	}

	node, ok := onecclusterservicelogic.CompareNodes(node, nodeInfo)
	if ok {
		node.UpdatedBy = in.UpdatedBy
		if err := l.svcCtx.NodeModel.Update(l.ctx, node); err != nil {
			l.Logger.Errorf("集群: %v , 更新节点: %v, 信息失败: %v", in.ClusterUuid, node.NodeName, err)
			return nil, code.SyncNodeInfoErr
		}
	}

	// 并行处理标签、注解和污点
	var wg sync.WaitGroup
	errChan := make(chan error, 3)

	// 处理标签
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := l.updateNodeLabels(nodeInfo.Labels, node); err != nil {
			l.Logger.Errorf("更新节点标签失败: %v, 集群: %v nodes: %v", err, in.ClusterUuid, node.NodeName)
			errChan <- code.SyncNodeLabelErr
		}
	}()

	// 处理注解
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := l.updateNodeAnnotations(nodeInfo.Annotations, node); err != nil {
			l.Logger.Errorf("更新节点注解失败: %v, 集群: %v nodes: %v", err, in.ClusterUuid, node.NodeName)
			errChan <- code.SyncNodeAnnotationsErr
		}
	}()

	// 处理污点
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := l.updateNodeTaints(nodeInfo.Taints, node); err != nil {
			l.Logger.Errorf("更新节点污点失败: %v, 集群: %v nodes: %v", err, in.ClusterUuid, node.NodeName)
			errChan <- code.SyncNodeTaintErr
		}
	}()

	// 等待所有并行任务完成
	wg.Wait()
	close(errChan)
	cluster, err := l.svcCtx.ClusterModel.FindOneByUuid(l.ctx, in.ClusterUuid)

	// 检查错误
	if len(errChan) > 0 {
		l.Logger.Errorf("同步部分节点失败: %v", err)
		cluster.Status = 0
		cluster.UpdatedBy = in.UpdatedBy
		if err := l.svcCtx.ClusterModel.Update(l.ctx, cluster); err != nil {
			l.Logger.Errorf("更新集群信息失败: %v", err)
			return nil, err
		}
		return nil, <-errChan // 返回第一个错误
	}

	// 同步完成后同步集群数据
	// 所有数据同步完成后进行回填数据
	if err != nil {
		l.Logger.Errorf("获取集群信息失败: %v", err)
		return nil, code.GetClusterInfoErr
	}
	totalInfo, err := l.svcCtx.NodeModel.FindOneClusterTotalInfo(l.ctx, cluster.Uuid)
	if err != nil {
		l.Logger.Errorf("获取集群节点信息失败: %v", err)
		return nil, errorx.DatabaseFindErr
	}
	cluster.UpdatedBy = in.UpdatedBy
	cluster.NodeCount = totalInfo.TotalNode
	cluster.CpuTotal = totalInfo.TotalCpu
	cluster.MemoryTotal = totalInfo.TotalMemory
	cluster.PodTotal = totalInfo.TotalPods

	// 更新集群
	if err := l.svcCtx.ClusterModel.Update(l.ctx, cluster); err != nil {
		l.Logger.Errorf("更新集群信息失败: %v", err)
		return nil, err
	}
	return &pb.SyncOnecNodeResp{}, nil
}

// 处理节点 标签信息
func (l *SyncOnecNodeLogic) updateNodeLabels(labels map[string]string, node *model.OnecNode) error {
	// 查询当前节点所有标签信息
	// 查询当前节点所有标签信息
	nodeLabels, err := l.svcCtx.LabelsResourceModel.SearchNoPage(
		l.ctx,
		"id",
		false,
		"resource_id = ? AND resource_type = ?",
		node.Id,
		"node",
	)
	if err != nil {
		l.Logger.Errorf("查询节点标签信息失败: %v", err)
		return err
	}
	// 构造数据库中标签的 map（用于对比）
	// 构造数据库中标签的 map（用于对比），key 为标签 Key，value 为 (ID, Value)
	dbLabelsMap := make(map[string]struct {
		ID    uint64
		Value string
	})
	for _, label := range nodeLabels {
		dbLabelsMap[label.Key] = struct {
			ID    uint64
			Value string
		}{
			ID:    label.Id,
			Value: label.Value,
		}
	}
	// 检查需要添加或更新的标签

	// 检查需要添加或更新的标签
	for key, value := range labels {
		dbLabel, exists := dbLabelsMap[key]
		if exists {
			// 如果标签已存在且值相同，跳过
			if dbLabel.Value == value {
				continue
			}

			// 如果值不同，执行更新操作
			err := l.svcCtx.LabelsResourceModel.Update(l.ctx, &model.OnecResourceLabels{
				Id:    dbLabel.ID,
				Key:   key,
				Value: value,
			})
			if err != nil {
				l.Logger.Errorf("更新节点标签失败, Node: %v, Key: %v, Value: %v, Error: %v", node.NodeUid, key, value, err)
				return err
			}
			l.Logger.Infof("更新节点标签成功, Node: %v, Key: %v, Value: %v", node.NodeUid, key, value)
		} else {
			// 如果标签不存在，执行插入操作
			_, err := l.svcCtx.LabelsResourceModel.Insert(l.ctx, &model.OnecResourceLabels{
				ResourceType: "node",
				ResourceId:   node.Id,
				Key:          key,
				Value:        value,
			})
			if err != nil {
				l.Logger.Errorf("添加节点标签失败, Node: %v, Key: %v, Value: %v, Error: %v", node.NodeUid, key, value, err)
				return err
			}
			l.Logger.Infof("添加节点标签成功, Node: %v, Key: %v, Value: %v", node.NodeUid, key, value)
		}
	}

	// 检查需要删除的标签（只能通过 ID 删除）
	for dbKey, dbLabel := range dbLabelsMap {
		if _, exists := labels[dbKey]; !exists {
			// 如果当前标签不存在于最新标签中，执行删除操作
			err := l.svcCtx.LabelsResourceModel.Delete(l.ctx, dbLabel.ID)
			if err != nil {
				l.Logger.Errorf("删除节点标签失败, Node: %v, Key: %v, ID: %v, Error: %v", node.NodeUid, dbKey, dbLabel.ID, err)
				return err
			}
			l.Logger.Infof("删除节点标签成功, Node: %v, Key: %v, ID: %v", node.NodeUid, dbKey, dbLabel.ID)
		}
	}

	return nil
}

func (l *SyncOnecNodeLogic) updateNodeAnnotations(annotations map[string]string, node *model.OnecNode) error {
	// 查询当前节点所有注解信息
	nodeAnnotations, err := l.svcCtx.AnnotationsResourceModel.SearchNoPage(
		l.ctx,
		"id",
		false,
		"resource_id = ? AND resource_type = ?",
		node.Id,
		"node",
	)
	if err != nil {
		l.Logger.Errorf("查询节点注解信息失败: %v", err)
		return err
	}

	// 构造数据库中注解的 map
	dbAnnotationsMap := make(map[string]struct {
		ID    uint64
		Value string
	})
	for _, annotation := range nodeAnnotations {
		dbAnnotationsMap[annotation.Key] = struct {
			ID    uint64
			Value string
		}{
			ID:    annotation.Id,
			Value: annotation.Value,
		}
	}

	// 检查需要添加或更新的注解
	for key, value := range annotations {
		dbAnnotation, exists := dbAnnotationsMap[key]
		if exists {
			// 如果注解已存在且值相同，跳过
			if dbAnnotation.Value == value {
				continue
			}

			// 如果值不同，执行更新操作
			err := l.svcCtx.AnnotationsResourceModel.Update(l.ctx, &model.OnecResourceAnnotations{
				Id:    dbAnnotation.ID,
				Key:   key,
				Value: value,
			})
			if err != nil {
				l.Logger.Errorf("更新节点注解失败, Node: %v, Key: %v, Value: %v, Error: %v", node.NodeUid, key, value, err)
				return err
			}
			l.Logger.Infof("更新节点注解成功, Node: %v, Key: %v, Value: %v", node.NodeUid, key, value)
		} else {
			// 如果注解不存在，执行插入操作
			_, err := l.svcCtx.AnnotationsResourceModel.Insert(l.ctx, &model.OnecResourceAnnotations{
				ResourceType: "node",
				ResourceId:   node.Id,
				Key:          key,
				Value:        value,
			})
			if err != nil {
				l.Logger.Errorf("添加节点注解失败, Node: %v, Key: %v, Value: %v, Error: %v", node.NodeUid, key, value, err)
				return err
			}
			l.Logger.Infof("添加节点注解成功, Node: %v, Key: %v, Value: %v", node.NodeUid, key, value)
		}
	}

	// 检查需要删除的注解（只能通过 ID 删除）
	for dbKey, dbAnnotation := range dbAnnotationsMap {
		if _, exists := annotations[dbKey]; !exists {
			// 如果当前注解不存在于最新注解中，执行删除操作
			err := l.svcCtx.AnnotationsResourceModel.Delete(l.ctx, dbAnnotation.ID)
			if err != nil {
				l.Logger.Errorf("删除节点注解失败, Node: %v, Key: %v, ID: %v, Error: %v", node.NodeUid, dbKey, dbAnnotation.ID, err)
				return err
			}
			l.Logger.Infof("删除节点注解成功, Node: %v, Key: %v, ID: %v", node.NodeUid, dbKey, dbAnnotation.ID)
		}
	}

	return nil
}

func (l *SyncOnecNodeLogic) updateNodeTaints(taints []core.Taint, node *model.OnecNode) error {
	// 查询当前节点所有污点信息
	nodeTaints, err := l.svcCtx.TaintsResourceModel.SearchNoPage(
		l.ctx,
		"id",
		false,
		"node_id = ?",
		node.Id,
	)
	if err != nil {
		l.Logger.Errorf("查询节点污点信息失败: %v", err)
		return err
	}

	// 构造数据库中污点的 map
	dbTaintsMap := make(map[string]struct {
		ID     uint64
		Value  string
		Effect string
	})
	for _, taint := range nodeTaints {
		dbTaintsMap[taint.Key] = struct {
			ID     uint64
			Value  string
			Effect string
		}{
			ID:     taint.Id,
			Value:  taint.Value,
			Effect: taint.EffectCode,
		}
	}

	// 检查需要添加或更新的污点
	for _, taint := range taints {
		dbTaint, exists := dbTaintsMap[taint.Key]
		if exists {
			// 如果污点已存在且值和效果相同，跳过
			if dbTaint.Value == taint.Value && dbTaint.Effect == string(taint.Effect) {
				continue
			}

			// 如果值或效果不同，执行更新操作
			err := l.svcCtx.TaintsResourceModel.Update(l.ctx, &model.OnecResourceTaints{
				Id:         dbTaint.ID,
				Key:        taint.Key,
				Value:      taint.Value,
				EffectCode: string(taint.Effect),
			})
			if err != nil {
				l.Logger.Errorf("更新节点污点失败, Node: %v, Key: %v, Value: %v, Effect: %v, Error: %v", node.NodeUid, taint.Key, taint.Value, taint.Effect, err)
				return err
			}
			l.Logger.Infof("更新节点污点成功, Node: %v, Key: %v, Value: %v, Effect: %v", node.NodeUid, taint.Key, taint.Value, taint.Effect)
		} else {
			// 如果污点不存在，执行插入操作
			_, err := l.svcCtx.TaintsResourceModel.Insert(l.ctx, &model.OnecResourceTaints{
				NodeId:     node.Id,
				Key:        taint.Key,
				Value:      taint.Value,
				EffectCode: string(taint.Effect),
			})
			if err != nil {
				l.Logger.Errorf("添加节点污点失败, Node: %v, Key: %v, Value: %v, Effect: %v, Error: %v", node.NodeUid, taint.Key, taint.Value, taint.Effect, err)
				return err
			}
			l.Logger.Infof("添加节点污点成功, Node: %v, Key: %v, Value: %v, Effect: %v", node.NodeUid, taint.Key, taint.Value, taint.Effect)
		}
	}

	// 检查需要删除的污点（只能通过 ID 删除）
	for dbKey, dbTaint := range dbTaintsMap {
		found := false
		for _, taint := range taints {
			if taint.Key == dbKey {
				found = true
				break
			}
		}
		if !found {
			// 如果当前污点不存在于最新污点中，执行删除操作
			err := l.svcCtx.TaintsResourceModel.Delete(l.ctx, dbTaint.ID)
			if err != nil {
				l.Logger.Errorf("删除节点污点失败, Node: %v, Key: %v, ID: %v, Error: %v", node.NodeUid, dbKey, dbTaint.ID, err)
				return err
			}
			l.Logger.Infof("删除节点污点成功, Node: %v, Key: %v, ID: %v", node.NodeUid, dbKey, dbTaint.ID)
		}
	}

	return nil
}
