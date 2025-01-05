package onecnodeservicelogic

import (
	"context"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/internal/code"
	onecclusterservicelogic "github.com/yanshicheng/kube-onec/application/manager/rpc/internal/logic/onecclusterservice"
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
	// 先获取 client
	client, err := l.svcCtx.OnecClient.GetOrCreateOnecK8sClient(l.ctx, in.ClusterId, nil)
	if err != nil {
		l.Logger.Infof("获取集群客户端失败: %v", err)
		cluster, err := l.svcCtx.ClusterModel.FindOne(l.ctx, in.ClusterId)
		if err != nil {
			l.Logger.Errorf("获取集群信息失败: %v", err)
			return nil, code.GetClusterInfoErr
		}
		client, err = l.svcCtx.OnecClient.GetOrCreateOnecK8sClient(l.ctx, in.ClusterId, utils.NewRestConfig(cluster.Host, cluster.Token, utils.IntToBool(cluster.SkipInsecure)))
		if err != nil {
			l.Logger.Infof("获取集群客户端失败: %v", err)
			return nil, code.GetClusterClientErr
		}
	}

	if err := client.Ping(); err != nil {
		l.Logger.Infof("集群: %v, 连接失败: %v", in.ClusterId, err)
		return nil, code.ClusterConnectErr
	}
	l.Logger.Infof("集群Id: %v, nodes: %v 正在更新", in.ClusterId, in.NodeName)
	// 获取单个数据
	nodeInfo, err := client.GetNodes().GetNodeInfo(in.NodeName)
	if err != nil {
		l.Logger.Errorf("获取节点信息失败: %v, 集群: %v nodes: %v", err, in.ClusterId, in.NodeName)
		return nil, code.GetNodeInfoErr
	}
	// 查询 node 数据
	node, err := l.svcCtx.NodeModel.FindOneByClusterIdNodeName(l.ctx, in.ClusterId, in.NodeName)
	if err != nil {
		l.Logger.Errorf("获取节点信息失败: %v, 集群: %v nodes: %v", err, in.ClusterId, in.NodeName)
		return nil, code.GetNodeInfoErr
	}
	node, ok := onecclusterservicelogic.CompareNodes(node, nodeInfo)
	if ok {
		node.UpdateBy = in.UpdateBy
		if err := l.svcCtx.NodeModel.Update(l.ctx, node); err != nil {
			l.Logger.Errorf("集群: %v , 更新节点: %v, 信息失败: %v", in.ClusterId, in.NodeName, err)
			return nil, code.SyncClusterInfoErr
		}
	}
	return &pb.SyncOnecNodeResp{}, nil
}
