package onecclusterservicelogic

import (
	"context"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/internal/code"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/internal/model"
	"github.com/yanshicheng/kube-onec/common/handler/errorx"
	"github.com/yanshicheng/kube-onec/common/k8swrapper/core"
	"github.com/yanshicheng/kube-onec/utils"
	"sync"

	"github.com/yanshicheng/kube-onec/application/manager/rpc/internal/svc"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/pb"
	portalPb "github.com/yanshicheng/kube-onec/application/portal/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

const (
	DictEnvCode         = "cluster_env"
	DictTaintEffectCode = "cluster_taint_effect"
)

type AddOnecClusterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddOnecClusterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddOnecClusterLogic {
	return &AddOnecClusterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// -----------------------集群表，用于管理多个 Kubernetes 集群-----------------------
func (l *AddOnecClusterLogic) AddOnecCluster(in *pb.AddOnecClusterReq) (*pb.AddOnecClusterResp, error) {
	switch in.ConnCode {
	case "token":
		return l.addTokenCluster(in)
	case "agent":
		return l.addTokenCluster(in)

	case "kubeconfig":
		return &pb.AddOnecClusterResp{}, nil

	default:
		return nil, code.UnsupportedConnTypeErr
	}
}

func (l *AddOnecClusterLogic) addTokenCluster(in *pb.AddOnecClusterReq) (*pb.AddOnecClusterResp, error) {
	clusterUuid, err := utils.GenerateRandomID()
	if err != nil {
		l.Logger.Errorf("生成UUID失败: %v", err)
		return nil, errorx.UUIDGenerateErr
	}
	config := utils.NewRestConfig(in.Host, in.Token, utils.IntToBool(in.SkipInsecure))
	client, err := l.svcCtx.OnecClient.GetOrCreateOnecK8sClient(l.ctx, clusterUuid, config)
	if err != nil {
		l.Logger.Errorf("获取集群客户端失败: %v", err)
		return nil, code.GetClusterClientErr
	}
	errs := client.Ping()
	if errs != nil {
		l.Logger.Errorf("集群连接失败: %v", errs)
		return nil, code.ClusterConnectErr
	}
	l.Logger.Infof("集群连接成功: %v", in.Host)

	// 获取集群信息
	clusterInfo, err := client.GetCluster().GetClusterInfo()
	if err != nil {
		l.Logger.Errorf("获取集群信息失败: %v", err)
		return nil, code.GetClusterInfoErr
	}

	// 检查 env Code 是否存在
	if _, err := l.svcCtx.SysDictItemRpc.CheckDictItemCode(l.ctx, &portalPb.CheckDictItemCodeReq{
		DictCode: DictEnvCode,
		ItemCode: in.EnvCode,
	}); err != nil {
		l.Logger.Errorf("字典项不存在: %v", err)
		return nil, code.DictItemNotExistErr
	}

	newCluster := &model.OnecCluster{
		Name:             in.Name,
		SkipInsecure:     utils.BoolToInt(utils.IntToBool(in.SkipInsecure)),
		Uuid:             clusterUuid,
		Host:             in.Host,
		Token:            in.Token,
		ConnCode:         in.ConnCode,
		EnvCode:          in.EnvCode,
		Location:         in.Location,
		NodeLbIp:         in.NodeLbIp,
		Status:           1,
		Version:          clusterInfo.Version,
		Commit:           clusterInfo.Commit,
		Platform:         clusterInfo.Platform,
		VersionBuildAt:   clusterInfo.BuildTime,
		ClusterCreatedAt: clusterInfo.CreateTime,
		Description:      in.Description,
		CreatedBy:        in.CreatedBy,
		UpdatedBy:        in.UpdatedBy,
	}
	result, err := l.svcCtx.ClusterModel.Insert(l.ctx, newCluster)
	if err != nil {
		l.Logger.Errorf("添加集群失败: %v", err)
		return nil, code.AddClusterFailedErr
	}
	id, err := result.LastInsertId()
	if err != nil {
		l.Logger.Errorf("获取集群ID失败: %v", err)
		return nil, code.GetClusterIdErr
	}
	l.Logger.Infof("集群: %v, 添加成功, ID: %d", newCluster.Name, id)
	newCluster.Id = uint64(id)

	// 增加集群节点信息
	nodes, err := client.GetNodes().GetAllNodesInfo()
	if err != nil {
		l.Logger.Errorf("获取节点信息失败: %v", err)
		newCluster.Status = 0
		err := l.svcCtx.ClusterModel.Update(l.ctx, newCluster)
		if err != nil {
			l.Logger.Errorf("获取节点异常导致，更新集群状态失败: %v", err)
		}
		return nil, code.GetNodeInfoErr
	}
	err = l.syncClusterNode(nodes, clusterUuid, in.CreatedBy, in.UpdatedBy)
	if err != nil {
		l.Logger.Errorf("获取节点信息失败: %v", err)
		newCluster.Status = 0
		err := l.svcCtx.ClusterModel.Update(l.ctx, newCluster)
		if err != nil {
			l.Logger.Errorf("获取节点异常导致，更新集群状态失败: %v", err)
		}
		return nil, code.GetNodeInfoErr
	}

	l.Logger.Infof("集群: %v, 节点信息同步成功", newCluster.Name)
	// 回填集群节点信息

	totalInfo, err := l.svcCtx.NodeModel.FindOneClusterTotalInfo(l.ctx, clusterUuid)
	if err != nil {
		l.Logger.Errorf("获取集群节点信息失败: %v", err)
		return nil, errorx.DatabaseFindErr
	}
	newCluster.NodeCount = totalInfo.TotalNode
	newCluster.CpuTotal = totalInfo.TotalCpu
	newCluster.MemoryTotal = totalInfo.TotalMemory
	newCluster.PodTotal = totalInfo.TotalPods

	// 更新集群
	if err := l.svcCtx.ClusterModel.Update(l.ctx, newCluster); err != nil {
		l.Logger.Errorf("更新集群信息失败: %v", err)
		return nil, err
	}
	return &pb.AddOnecClusterResp{}, nil
}

func (l *AddOnecClusterLogic) syncClusterNode(nodes []*core.NodeInfo, clusterUuid, createdBy, updateBy string) error {
	var wg sync.WaitGroup
	var mu sync.Mutex
	var syncErr error

	for _, node := range nodes {
		wg.Add(1)
		go func(node *core.NodeInfo) {
			defer wg.Done()

			nodeInfo, err := GenerateNewNode(node)
			if err != nil {
				l.Logger.Errorf("节点信息转换失败: %v", err)
				mu.Lock()
				syncErr = err
				mu.Unlock()
				return
			}
			nodeInfo.ClusterUuid = clusterUuid
			nodeInfo.CreatedBy = createdBy
			nodeInfo.UpdatedBy = updateBy

			// 插入节点信息
			result, err := l.svcCtx.NodeModel.Insert(l.ctx, nodeInfo)
			if err != nil {
				l.Logger.Errorf("节点信息插入失败: %v", err)
				mu.Lock()
				syncErr = err
				mu.Unlock()
				return
			}

			nodeID, _ := result.LastInsertId()
			nodeInfo.Id = uint64(nodeID)

			// 处理 Labels、Annotations 和 Taints
			if err := l.processNodeAttributes(node, nodeInfo, clusterUuid); err != nil {
				mu.Lock()
				syncErr = err
				mu.Unlock()
				return
			}
			l.Logger.Infof("集群: %v, 节点: %v, 信息同步成功", clusterUuid, nodeInfo.NodeName)
		}(node)
	}

	wg.Wait()
	return syncErr
}

func (l *AddOnecClusterLogic) processNodeAttributes(node *core.NodeInfo, nodeInfo *model.OnecNode, clusterUuid string) error {
	// 处理 Labels
	for key, value := range node.Labels {
		_, err := l.svcCtx.LabelsResourceModel.Insert(l.ctx, &model.OnecResourceLabels{
			ResourceType: "node",
			ResourceId:   nodeInfo.Id,
			Key:          key,
			Value:        value,
		})
		if err != nil {
			l.Logger.Errorf("集群: %v, 节点: %v, 添加标签失败: %v", clusterUuid, nodeInfo.NodeName, key)
			return err
		}
	}

	// 处理 Annotations
	for key, value := range node.Annotations {
		_, err := l.svcCtx.AnnotationsResourceModel.Insert(l.ctx, &model.OnecResourceAnnotations{
			ResourceType: "node",
			ResourceId:   nodeInfo.Id,
			Key:          key,
			Value:        value,
		})
		if err != nil {
			l.Logger.Errorf("集群: %v, 节点: %v, 添加注解失败: %v", clusterUuid, nodeInfo.NodeName, key)
			return err
		}
	}

	// 处理 Taints
	for _, taint := range node.Taints {
		_, err := l.svcCtx.TaintsResourceModel.Insert(l.ctx, &model.OnecResourceTaints{
			NodeId:     nodeInfo.Id,
			Key:        taint.Key,
			Value:      taint.Value,
			EffectCode: taint.Effect,
		})
		if err != nil {
			l.Logger.Errorf("集群: %v, 节点: %v, 添加污点失败: %v", clusterUuid, nodeInfo.NodeName, taint.Key)
			return err
		}
	}

	return nil
}

func GenerateNewNode(node *core.NodeInfo) (*model.OnecNode, error) {
	//labelsStr, err := utils.MapStringToJSON(node.Labels)
	//if err != nil {
	//	return nil, code.GetNodeInfoErr
	//}
	//
	//annotationStr, err := utils.MapStringToJSON(node.Annotations)
	//if err != nil {
	//	return nil, code.GetNodeInfoErr
	//}
	//trainsStr, err := utils.TaintsToJSON(node.Taints)
	//if err != nil {
	//	return nil, code.GetNodeInfoErr
	//}
	nodeInfo := &model.OnecNode{
		NodeName:         node.NodeName,
		NodeUid:          node.NodeUid,
		Status:           node.Status,
		Roles:            node.Roles,
		JoinAt:           node.JoinTime,
		PodCidr:          node.PodCIDR,
		Unschedulable:    utils.BoolToInt(node.Unschedulable),
		NodeIp:           node.NodeIp,
		Os:               node.Os,
		MaxPods:          node.MaxPods,
		Cpu:              node.Cpu,
		Memory:           node.Memory,
		KernelVersion:    node.KernelVersion,
		ContainerRuntime: node.ContainerRuntime,
		KubeletVersion:   node.KubeletVersion,
		KubeletPort:      uint64(node.KubeletPort),
		OperatingSystem:  node.OperatingSystem,
		Architecture:     node.Architecture,
	}
	return nodeInfo, nil
}
