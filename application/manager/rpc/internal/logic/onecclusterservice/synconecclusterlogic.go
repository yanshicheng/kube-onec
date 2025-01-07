package onecclusterservicelogic

import (
	"context"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/internal/code"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/internal/model"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/internal/svc"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/pb"
	"github.com/yanshicheng/kube-onec/common/handler/errorx"
	"github.com/yanshicheng/kube-onec/common/k8swrapper/core"
	"github.com/yanshicheng/kube-onec/utils"
	"github.com/zeromicro/go-zero/core/logx"
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
	client, err := l.svcCtx.OnecClient.GetOrCreateOnecK8sClient(l.ctx, cluster.Uuid, utils.NewRestConfig(cluster.Host, cluster.Token, utils.IntToBool(cluster.SkipInsecure)))
	if err != nil {
		l.Logger.Infof("获取集群客户端失败: %v", err)
		return nil, code.GetClusterClientErr
	}

	if err := client.Ping(); err != nil {
		l.Logger.Infof("集群连接失败: %v", err)
		return nil, code.ClusterConnectErr
	}

	// 同步集群信息
	clusterInfo, err := client.GetCluster().GetClusterInfo()
	if err != nil {
		l.Logger.Errorf("获取集群信息失败: %v", err)
		return nil, code.GetClusterInfoErr
	}
	err = l.updateClusterIfChanged(cluster, clusterInfo, in.UpdatedBy)
	if err != nil {
		return nil, err
	}

	// 同步节点信息
	nodeList, err := client.GetNodes().GetAllNodesInfo()
	if err != nil {
		l.Logger.Errorf("获取节点信息失败: %v", err)
		return nil, code.GetNodeInfoErr
	}
	nodeModeList, err := l.svcCtx.NodeModel.SearchNoPage(l.ctx, "create_time", false, "cluster_uuid = ?", cluster.Uuid)
	if err != nil {
		l.Logger.Errorf("获取节点信息失败: %v", err)
		return nil, code.GetNodeInfoErr
	}
	dbNodeMap := make(map[string]*model.OnecNode)
	for _, node := range nodeModeList {
		dbNodeMap[node.NodeUid] = node
	}

	currentNodeMap := make(map[string]*core.NodeInfo)
	for _, node := range nodeList {
		currentNodeMap[node.NodeUid] = node
	}

	// 如果不存在则修改状态位 Unknown
	for _, dbNode := range nodeModeList {
		if _, exists := currentNodeMap[dbNode.NodeUid]; !exists {
			dbNode.Status = "Unknown"
			if err := l.svcCtx.NodeModel.Update(l.ctx, dbNode); err != nil {
				l.Logger.Errorf("集群: %v , 删除更新节点: %v, 信息失败: %v", cluster.Name, dbNode.NodeName, err)
				l.ChangeCLusterStatusFalse(cluster, in.UpdatedBy)
				return nil, code.SyncClusterInfoErr
			}
		}
	}

	// 处理添加节点和更新节点
	for _, node := range currentNodeMap {
		existingNode, exists := dbNodeMap[node.NodeUid]
		if exists {
			// 更新节点
			newModel, ok := CompareNodes(existingNode, node)
			if ok {
				newModel.UpdatedBy = in.UpdatedBy
				l.Logger.Infof("集群: %v 更新节点信息: %v", cluster.Name, newModel.NodeName)
				if err := l.svcCtx.NodeModel.Update(l.ctx, newModel); err != nil {
					l.ChangeCLusterStatusFalse(cluster, in.UpdatedBy)
					l.Logger.Errorf("集群: %v ,更新节点: %v, 信息失败: %v", cluster.Name, newModel.NodeName, err)
					return nil, code.SyncClusterInfoErr
				}
			}
		} else {
			// 添加节点
			nodeModel, err := GenerateNewNode(node)
			if err != nil {
				l.ChangeCLusterStatusFalse(cluster, in.UpdatedBy)
				l.Logger.Errorf("集群: %v , 添加节点: %v, 信息失败: %v", cluster.Name, node.NodeName, err)
				return nil, code.SyncClusterInfoErr
			}
			nodeModel.ClusterUuid = cluster.Uuid
			nodeModel.CreatedBy = in.UpdatedBy
			nodeModel.UpdatedBy = in.UpdatedBy
			if _, err := l.svcCtx.NodeModel.Insert(l.ctx, nodeModel); err != nil {
				l.ChangeCLusterStatusFalse(cluster, in.UpdatedBy)
				l.Logger.Errorf("集群: %v , 添加节点: %v, 信息失败: %v", cluster.Name, node.NodeName, err)
				return nil, code.SyncClusterInfoErr
			}

		}
	}
	totalInfo, err := l.svcCtx.NodeModel.FindOneClusterTotalInfo(l.ctx, cluster.Uuid)
	if err != nil {
		l.Logger.Errorf("获取集群节点信息失败: %v", err)
		return nil, errorx.DatabaseFindErr
	}
	cluster.NodeCount = totalInfo.TotalNode
	cluster.CpuTotal = totalInfo.TotalCpu
	cluster.MemoryTotal = totalInfo.TotalMemory
	cluster.PodTotal = totalInfo.TotalPods

	// 更新集群
	if err := l.svcCtx.ClusterModel.Update(l.ctx, cluster); err != nil {
		l.Logger.Errorf("更新集群信息失败: %v", err)
		return nil, err
	}
	return &pb.SyncOnecClusterResp{}, nil
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
	nodeUnschedulableInt := utils.BoolToInt(node.Unschedulable)
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
