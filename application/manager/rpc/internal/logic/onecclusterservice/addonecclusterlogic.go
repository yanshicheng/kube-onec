package onecclusterservicelogic

import (
	"context"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/internal/code"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/internal/model"
	"github.com/yanshicheng/kube-onec/common/handler/errorx"
	"github.com/yanshicheng/kube-onec/common/k8swrapper/core"
	"github.com/yanshicheng/kube-onec/utils"

	"github.com/yanshicheng/kube-onec/application/manager/rpc/internal/svc"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
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
	switch in.ConnType {
	case 0:
		return &pb.AddOnecClusterResp{}, nil
	case 1:
		clusterId, err := utils.GenerateRandomID()
		if err != nil {
			l.Logger.Errorf("生成UUID失败: %v", err)
			return nil, errorx.UUIDGenerateErr
		}
		config := utils.NewRestConfig(in.Host, in.Token, utils.IntToBool(in.SkipInsecure))
		client, err := l.svcCtx.OnecClient.GetOrCreateOnecK8sClient(l.ctx, clusterId, config)
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
		newCluster := &model.OnecCluster{
			Id:                clusterId,
			Name:              in.Name,
			SkipInsecure:      utils.BoolToInt(utils.IntToBool(in.SkipInsecure)),
			Host:              in.Host,
			Token:             in.Token,
			ConnType:          in.ConnType.String(),
			EnvTag:            in.EnvTag,
			Status:            1,
			Version:           clusterInfo.Version,
			Commit:            clusterInfo.Commit,
			Platform:          clusterInfo.Platform,
			VersionBuildTime:  clusterInfo.BuildTime,
			ClusterCreateTime: clusterInfo.CreateTime,
			Description:       in.Description,
			CreateBy:          in.CreateBy,
			UpdateBy:          in.UpdateBy,
		}
		_, err = l.svcCtx.ClusterModel.Insert(l.ctx, newCluster)

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
		for _, node := range nodes {
			nodeId, err := utils.GenerateRandomID()
			if err != nil {
				l.Logger.Errorf("生成UUID失败: %v", err)
				newCluster.Status = 0
				err := l.svcCtx.ClusterModel.Update(l.ctx, newCluster)
				if err != nil {
					l.Logger.Errorf("获取节点异常导致，更新集群状态失败: %v", err)
				}
				return nil, errorx.UUIDGenerateErr
			}

			nodeInfo, err := GenerateNewNode(node)
			if err != nil {
				l.Logger.Errorf("节点信息转换失败: %v", err)
				newCluster.Status = 0
				err := l.svcCtx.ClusterModel.Update(l.ctx, newCluster)
				if err != nil {
					l.Logger.Errorf("获取节点异常导致，更新集群状态失败: %v", err)
				}
				return nil, code.GetNodeInfoErr
			}
			nodeInfo.Id = nodeId
			nodeInfo.ClusterId = clusterId
			nodeInfo.CreateBy = in.CreateBy
			nodeInfo.UpdateBy = in.UpdateBy
			_, err = l.svcCtx.NodeModel.Insert(l.ctx, nodeInfo)
			if err != nil {
				l.Logger.Errorf("节点信息插入失败: %v", err)
				newCluster.Status = 0
				err := l.svcCtx.ClusterModel.Update(l.ctx, newCluster)
				if err != nil {
					l.Logger.Errorf("获取节点异常导致，更新集群: %v, 状态失败: %v", newCluster.Name, err)
				}
				return nil, code.AddNodeInfoErr
			}
			l.Logger.Infof("集群: %v, 节点信息插入成功: %v", newCluster.Name, nodeInfo.NodeName)
		}
		return &pb.AddOnecClusterResp{}, nil
	default:
		return nil, code.UnsupportedConnTypeErr
	}
}

func GenerateNewNode(node *core.NodeInfo) (*model.OnecNode, error) {
	labelsStr, err := utils.MapStringToJSON(node.Labels)
	if err != nil {
		return nil, code.GetNodeInfoErr
	}

	annotationStr, err := utils.MapStringToJSON(node.Annotations)
	if err != nil {
		return nil, code.GetNodeInfoErr
	}
	trainsStr, err := utils.TaintsToJSON(node.Taints)
	if err != nil {
		return nil, code.GetNodeInfoErr
	}
	nodeInfo := &model.OnecNode{
		NodeName:         node.NodeName,
		NodeUid:          node.NodeUid,
		Status:           node.Status,
		Roles:            node.Roles,
		JoinTime:         node.JoinTime,
		Labels:           labelsStr,
		Annotations:      annotationStr,
		PodCidr:          node.PodCIDR,
		Unschedulable:    utils.BoolToInt(node.Unschedulable),
		Taints:           trainsStr,
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
