package onecclusterservicelogic

import (
	"context"
	"fmt"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/internal/code"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/internal/model"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/internal/svc"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/pb"
	"github.com/yanshicheng/kube-onec/common/handler/errorx"
	"github.com/yanshicheng/kube-onec/common/k8swrapper/core"
	utils2 "github.com/yanshicheng/kube-onec/pkg/utils"
	"github.com/zeromicro/go-zero/core/logx"
	"sync"
)

type SyncOnecClusterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSyncOnecClusterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SyncOnecClusterLogic {
	return &SyncOnecClusterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 同步集群信息
func (l *SyncOnecClusterLogic) SyncOnecCluster(in *pb.SyncOnecClusterReq) (*pb.SyncOnecClusterResp, error) {
	//
	cluster, err := l.svcCtx.ClusterModel.FindOne(l.ctx, in.Id)
	if err != nil {
		l.Logger.Errorf("获取集群信息失败: %v", err)
		return nil, code.GetClusterInfoErr
	}
	client, err := l.svcCtx.OnecClient.GetOrCreateOnecK8sClient(l.ctx, cluster.Uuid, utils2.NewRestConfig(cluster.Host, cluster.Token, utils2.IntToBool(cluster.SkipInsecure)))
	if err != nil {
		l.Logger.Infof("获取集群客户端失败: %v", err)
		return nil, code.GetClusterClientErr
	}

	if err := client.Ping(); err != nil {
		l.Logger.Infof("集群连接失败: %v", err)
		return nil, code.ClusterConnectErr
	}

	// 同步集群信息
	clusterInfo, err := client.GetClusterClient().GetClusterInfo()
	if err != nil {
		l.Logger.Errorf("获取集群信息失败: %v", err)
		return nil, code.GetClusterInfoErr
	}
	err = l.updateClusterIfChanged(cluster, clusterInfo, in.UpdatedBy)
	if err != nil {
		return nil, err
	}
	nodeList, err := client.GetNodeClient().GetAllNodesInfo()
	if err != nil {
		l.Logger.Errorf("获取节点信息失败: %v", err)
		return nil, code.GetNodeInfoErr
	}

	// 同步集群节点
	err = l.SyncAllNode(nodeList, cluster.Uuid, in.UpdatedBy)
	if err != nil {
		l.Logger.Errorf("节点同步失败: %v", err)
		l.ChangeCLusterStatusFalse(cluster, in.UpdatedBy)
		return nil, err
	}

	// 所有数据同步完成后进行回填数据
	totalInfo, err := l.svcCtx.NodeModel.FindOneClusterTotalInfo(l.ctx, cluster.Uuid)
	if err != nil {
		l.Logger.Errorf("获取集群节点信息失败: %v", err)
		return nil, errorx.DatabaseFindErr
	}
	// 获取已经分配的资源
	clusterUsed, err := l.svcCtx.ProjectQuotaModel.FindClusterQuotasByUuid(l.ctx, cluster.Uuid)
	if err == nil {
		cluster.CpuUsed = clusterUsed.CPUQuota
		cluster.MemoryUsed = clusterUsed.MemoryQuota
		cluster.PodUsed = clusterUsed.PodLimit
	}
	cluster.NodeCount = totalInfo.TotalNode
	cluster.CpuTotal = totalInfo.TotalCpu
	cluster.MemoryTotal = totalInfo.TotalMemory
	cluster.PodTotal = totalInfo.TotalPods

	// 更新集群
	if err := l.svcCtx.ClusterModel.Update(l.ctx, cluster); err != nil {
		l.Logger.Errorf("更新集群信息失败: %v", err)
		return nil, code.UpdateClusterInfoErr
	}

	return &pb.SyncOnecClusterResp{}, nil
}

//
//func (l *SyncOnecClusterLogic) SyncAllNode(nodeList []*core.NodeInfo, clusterUuid, updatedBy string) error {
//	flage := false // 设置标志 如果有任意一个node有问题 都要返回 err 外面在设置集群状态异常
//	// 同步节点信息
//	nodeModeList, err := l.svcCtx.NodeModel.SearchNoPage(l.ctx, "created_at", false, "cluster_uuid = ?", clusterUuid)
//	if err != nil {
//		l.Logger.Errorf("获取节点信息失败: %v", err)
//		return err
//	}
//	dbNodeMap := make(map[string]*model.OnecNode)
//	for _, node := range nodeModeList {
//		dbNodeMap[node.NodeUid] = node
//	}
//
//	currentNodeMap := make(map[string]*core.NodeInfo)
//	for _, node := range nodeList {
//		currentNodeMap[node.NodeUid] = node
//	}
//
//	// 如果不存在则修改状态位 Unknown
//	for _, dbNode := range nodeModeList {
//		if _, exists := currentNodeMap[dbNode.NodeUid]; !exists {
//			dbNode.Status = "Unknown"
//			if err := l.svcCtx.NodeModel.Update(l.ctx, dbNode); err != nil {
//				l.Logger.Errorf("集群: %v , 删除更新节点: %v, 信息失败: %v", clusterUuid, dbNode.NodeName, err)
//				return err
//			}
//		}
//	}
//
//	// 处理添加节点和更新节点
//	for _, node := range currentNodeMap {
//		existingNode, exists := dbNodeMap[node.NodeUid]
//		if exists {
//			// 更新节点
//			newModel, ok := CompareNodes(existingNode, node)
//			if ok {
//				newModel.UpdatedBy = updatedBy
//				l.Logger.Infof("集群: %v 更新节点信息: %v", clusterUuid, newModel.NodeName)
//				if err := l.svcCtx.NodeModel.Update(l.ctx, newModel); err != nil {
//					l.Logger.Errorf("集群: %v ,更新节点: %v, 信息失败: %v", clusterUuid, newModel.NodeName, err)
//					flage = true
//					continue
//				}
//
//				if err := l.updateNodeLabels(node.Labels, newModel); err != nil {
//					l.ChangeNodeStatusFalse(newModel, updatedBy)
//					flage = true
//				}
//				if err := l.updateNodeAnnotations(node.Annotations, newModel); err != nil {
//					l.ChangeNodeStatusFalse(newModel, updatedBy)
//
//					flage = true
//				}
//				if err := l.updateNodeTaints(node.Taints, newModel); err != nil {
//					l.ChangeNodeStatusFalse(newModel, updatedBy)
//					flage = true
//				}
//			}
//		} else {
//			// 添加节点
//			nodeModel, err := GenerateNewNode(node)
//			if err != nil {
//				l.Logger.Errorf("集群: %v , 添加节点: %v, 信息失败: %v", clusterUuid, node.NodeName, err)
//				flage = true
//				continue
//			}
//			nodeModel.ClusterUuid = clusterUuid
//			nodeModel.CreatedBy = updatedBy
//			nodeModel.UpdatedBy = updatedBy
//			if _, err := l.svcCtx.NodeModel.Insert(l.ctx, nodeModel); err != nil {
//				flage = true
//				continue
//			}
//
//			if err := l.updateNodeLabels(node.Labels, nodeModel); err != nil {
//				l.ChangeNodeStatusFalse(nodeModel, updatedBy)
//				flage = true
//			}
//			if err := l.updateNodeAnnotations(node.Annotations, nodeModel); err != nil {
//				l.ChangeNodeStatusFalse(nodeModel, updatedBy)
//				flage = true
//			}
//			if err := l.updateNodeTaints(node.Taints, nodeModel); err != nil {
//				l.ChangeNodeStatusFalse(nodeModel, updatedBy)
//				flage = true
//			}
//		}
//	}
//	if flage {
//		return fmt.Errorf("同步节点信息失败")
//	}
//	return nil
//}

func (l *SyncOnecClusterLogic) SyncAllNode(nodeList []*core.NodeInfo, clusterUuid, updatedBy string) error {
	// 并发控制
	maxConcurrency := 10 // 设置最大并发数
	semaphore := make(chan struct{}, maxConcurrency)
	var wg sync.WaitGroup
	var mu sync.Mutex

	var errList []error // 用于收集所有错误

	// 创建数据库中的节点映射
	nodeModeList, err := l.svcCtx.NodeModel.SearchNoPage(l.ctx, "created_at", false, "cluster_uuid = ?", clusterUuid)
	if err != nil {
		l.Logger.Errorf("获取数据库中的节点信息失败: %v", err)
		return err
	}

	// 构造数据库中节点的 Map
	dbNodeMap := make(map[string]*model.OnecNode)
	for _, node := range nodeModeList {
		dbNodeMap[node.NodeUid] = node
	}

	// 构造当前集群节点的 Map
	currentNodeMap := make(map[string]*core.NodeInfo)
	for _, node := range nodeList {
		currentNodeMap[node.NodeUid] = node
	}

	// 对每个节点进行同步
	for _, node := range nodeList {
		// 控制并发
		semaphore <- struct{}{}
		wg.Add(1)

		go func(node *core.NodeInfo) {
			defer wg.Done()
			defer func() { <-semaphore }() // 释放信号量

			// 节点同步逻辑
			if err := l.processSingleNode(node, clusterUuid, updatedBy, dbNodeMap); err != nil {
				mu.Lock()
				errList = append(errList, err)
				mu.Unlock()
			}
		}(node)
	}

	// 等待所有 Goroutine 完成
	wg.Wait()

	// 如果有错误，返回错误信息
	if len(errList) > 0 {
		l.Logger.Errorf("部分节点同步失败: %v", errList)
		return fmt.Errorf("同步部分节点失败，共计 %d 个错误", len(errList))
	}

	return nil
}

// 单个节点的同步逻辑
func (l *SyncOnecClusterLogic) processSingleNode(node *core.NodeInfo, clusterUuid, updatedBy string, dbNodeMap map[string]*model.OnecNode) error {
	// 检查节点是否已存在
	existingNode, exists := dbNodeMap[node.NodeUid]
	if exists {
		// 更新节点
		newModel, ok := CompareNodes(existingNode, node)
		if ok {
			newModel.UpdatedBy = updatedBy
			l.Logger.Infof("集群: %v 更新节点信息: %v", clusterUuid, newModel.NodeName)
			if err := l.svcCtx.NodeModel.Update(l.ctx, newModel); err != nil {
				l.Logger.Errorf("更新节点失败: %v", err)
				return fmt.Errorf("更新节点 %s 失败: %v", newModel.NodeName, err)
			}

			// 同步节点的标签、注解、污点
			if err := l.updateNodeLabels(node.Labels, newModel); err != nil {
				l.Logger.Errorf("更新节点标签失败: %v", err)
				return fmt.Errorf("更新节点 %s 标签失败: %v", newModel.NodeName, err)
			}

			if err := l.updateNodeAnnotations(node.Annotations, newModel); err != nil {
				l.Logger.Errorf("更新节点注解失败: %v", err)
				return fmt.Errorf("更新节点 %s 注解失败: %v", newModel.NodeName, err)
			}

			if err := l.updateNodeTaints(node.Taints, newModel); err != nil {
				l.Logger.Errorf("更新节点污点失败: %v", err)
				return fmt.Errorf("更新节点 %s 污点失败: %v", newModel.NodeName, err)
			}
		}
	} else {
		// 添加节点
		nodeModel, err := GenerateNewNode(node)
		if err != nil {
			l.Logger.Errorf("生成节点模型失败: %v", err)
			return fmt.Errorf("生成节点模型失败: %v", err)
		}

		nodeModel.ClusterUuid = clusterUuid
		nodeModel.CreatedBy = updatedBy
		nodeModel.UpdatedBy = updatedBy

		if _, err := l.svcCtx.NodeModel.Insert(l.ctx, nodeModel); err != nil {
			l.Logger.Errorf("插入节点失败: %v", err)
			return fmt.Errorf("插入节点失败: %v", err)
		}

		// 同步节点的标签、注解、污点
		if err := l.updateNodeLabels(node.Labels, nodeModel); err != nil {
			l.Logger.Errorf("添加节点标签失败: %v", err)
			return fmt.Errorf("添加节点 %s 标签失败: %v", nodeModel.NodeName, err)
		}

		if err := l.updateNodeAnnotations(node.Annotations, nodeModel); err != nil {
			l.Logger.Errorf("添加节点注解失败: %v", err)
			return fmt.Errorf("添加节点 %s 注解失败: %v", nodeModel.NodeName, err)
		}

		if err := l.updateNodeTaints(node.Taints, nodeModel); err != nil {
			l.Logger.Errorf("添加节点污点失败: %v", err)
			return fmt.Errorf("添加节点 %s 污点失败: %v", nodeModel.NodeName, err)
		}
	}

	return nil
}

// 设置node状态异常
func (l *SyncOnecClusterLogic) ChangeNodeStatusFalse(node *model.OnecNode, updatedBy string) {
	node.Status = "SyncErr"
	node.UpdatedBy = updatedBy
	if err := l.svcCtx.NodeModel.Update(l.ctx, node); err != nil {
		l.Logger.Errorf("更新集群状态失败: %v", err)
	}
}

// 处理节点 标签信息
func (l *SyncOnecClusterLogic) updateNodeLabels(labels map[string]string, node *model.OnecNode) error {
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

func (l *SyncOnecClusterLogic) updateNodeAnnotations(annotations map[string]string, node *model.OnecNode) error {
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

func (l *SyncOnecClusterLogic) updateNodeTaints(taints []core.Taint, node *model.OnecNode) error {
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

// updateClusterIfChanged 比较现有集群字段和传入的 clusterInfo，只有在发生变化时才更新
func (l *SyncOnecClusterLogic) updateClusterIfChanged(cluster *model.OnecCluster, clusterInfo *core.ClusterInfo, updatedBy string) error {
	// 判断集群版本信息是否发生变化
	if cluster.Version != clusterInfo.Version {
		cluster.Version = clusterInfo.Version
	}
	// 判断集群提交信息是否发生变化
	if cluster.Commit != clusterInfo.Commit {
		cluster.Commit = clusterInfo.Commit
	}
	// 判断平台信息是否发生变化
	if cluster.Platform != clusterInfo.Platform {
		cluster.Platform = clusterInfo.Platform
	}
	// 判断构建时间是否发生变化
	if cluster.VersionBuildAt != clusterInfo.BuildTime {
		cluster.VersionBuildAt = clusterInfo.BuildTime
	}
	// 判断集群创建时间是否发生变化
	if !cluster.ClusterCreatedAt.Equal(clusterInfo.CreateTime) {
		cluster.ClusterCreatedAt = clusterInfo.CreateTime
	}

	// 更新修改人
	if cluster.UpdatedBy != updatedBy {
		cluster.UpdatedBy = updatedBy
	}

	// 如果有任何字段被修改，执行更新操作
	if err := l.svcCtx.ClusterModel.Update(l.ctx, cluster); err != nil {
		l.Logger.Errorf("更新集群信息失败: %v", err)
		return code.SyncClusterInfoErr
	}

	return nil
}

// 修改集群状态为false
func (l *SyncOnecClusterLogic) ChangeCLusterStatusFalse(cluster *model.OnecCluster, updatedBy string) {

	cluster.Status = 0
	cluster.UpdatedBy = updatedBy
	if err := l.svcCtx.ClusterModel.Update(l.ctx, cluster); err != nil {
		l.Logger.Errorf("更新集群信息失败: %v", err)
	}
}

// CompareNodes 比较 existingNode 和 node 的所有字段，返回更新后的 existingNode 和是否需要更新的标志
func CompareNodes(existingNode *model.OnecNode, node *core.NodeInfo) (*model.OnecNode, bool) {
	flag := false

	// 比较并更新 Status
	if existingNode.Status != node.Status {
		existingNode.Status = node.Status
		flag = true
	}

	// 比较并更新 NodeIp
	if existingNode.NodeIp != node.NodeIp {
		existingNode.NodeIp = node.NodeIp
		flag = true
	}

	// 比较并更新 NodeName
	if existingNode.NodeName != node.NodeName {
		existingNode.NodeName = node.NodeName
		flag = true
	}

	// 比较并更新 Roles
	if existingNode.Roles != node.Roles {
		existingNode.Roles = node.Roles
		flag = true
	}

	// 比较并更新 JoinTime
	if !existingNode.JoinAt.Equal(node.JoinTime) {
		existingNode.JoinAt = node.JoinTime
		flag = true
	}

	// 比较并更新 PodCIDR
	if existingNode.PodCidr != node.PodCIDR {
		existingNode.PodCidr = node.PodCIDR
		flag = true
	}

	// 比较并更新 Unschedulable
	nodeUnschedulableInt := utils2.BoolToInt(node.Unschedulable)
	if existingNode.Unschedulable != nodeUnschedulableInt {
		existingNode.Unschedulable = nodeUnschedulableInt
		flag = true
	}

	// 比较并更新 OS
	if existingNode.Os != node.Os {
		existingNode.Os = node.Os
		flag = true
	}

	// 比较并更新 CPU
	if existingNode.Cpu != node.Cpu {
		existingNode.Cpu = node.Cpu
		flag = true
	}

	// 比较并更新 Memory
	if existingNode.Memory != node.Memory {
		existingNode.Memory = node.Memory
		flag = true
	}

	// 比较并更新 MaxPods
	if existingNode.MaxPods != node.MaxPods {
		existingNode.MaxPods = node.MaxPods
		flag = true
	}

	// 比较并更新 KernelVersion
	if existingNode.KernelVersion != node.KernelVersion {
		existingNode.KernelVersion = node.KernelVersion
		flag = true
	}

	// 比较并更新 ContainerRuntime
	if existingNode.ContainerRuntime != node.ContainerRuntime {
		existingNode.ContainerRuntime = node.ContainerRuntime
		flag = true
	}

	// 比较并更新 KubeletVersion
	if existingNode.KubeletVersion != node.KubeletVersion {
		existingNode.KubeletVersion = node.KubeletVersion
		flag = true
	}

	// 比较并更新 KubeletPort
	nodeKubeletPort := uint64(node.KubeletPort)
	if existingNode.KubeletPort != nodeKubeletPort {
		existingNode.KubeletPort = nodeKubeletPort
		flag = true
	}

	// 比较并更新 OperatingSystem
	if existingNode.OperatingSystem != node.OperatingSystem {
		existingNode.OperatingSystem = node.OperatingSystem
		flag = true
	}

	// 比较并更新 Architecture
	if existingNode.Architecture != node.Architecture {
		existingNode.Architecture = node.Architecture
		flag = true
	}
	// 比较并更新 Labels
	//existingNodeLabels, err := utils.JSONToMapString(existingNode.Labels)
	//if err != nil {
	//	existingNode.Status = "Unknown"
	//	flag = true
	//}
	//if !reflect.DeepEqual(existingNodeLabels, node.Labels) {
	//	nodeLabelsJSON, err := utils.MapStringToJSON(node.Labels)
	//	if err != nil {
	//		existingNode.Status = "Unknown"
	//		flag = true
	//	}
	//	existingNode.Labels = nodeLabelsJSON
	//	flag = true
	//}

	// 比较并更新 Annotations
	//existingNodeAnnotations, err := utils.JSONToMapString(existingNode.Annotations)
	//if !reflect.DeepEqual(existingNodeAnnotations, node.Annotations) {
	//	nodeAnnotationsJSON, err := utils.MapStringToJSON(node.Annotations)
	//	if err != nil {
	//		existingNode.Status = "Unknown"
	//		flag = true
	//	}
	//	existingNode.Annotations = nodeAnnotationsJSON
	//	flag = true
	//}

	// 比较并更新 Taints

	//existingNodeTaints, err := utils.JSONToTaints(existingNode.Taints)
	//if err != nil {
	//	existingNode.Status = "Unknown"
	//	flag = true
	//} else if !reflect.DeepEqual(existingNodeTaints, node.Taints) {
	//	nodeTaintsJSON, err := utils.TaintsToJSON(node.Taints)
	//	if err != nil {
	//		existingNode.Status = "Unknown"
	//		flag = true
	//	}
	//	existingNode.Taints = nodeTaintsJSON
	//	flag = true
	//}
	return existingNode, flag
}
