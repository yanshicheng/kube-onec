package onecclusterservicelogic

import (
	"context"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/internal/lib"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/internal/model"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/internal/svc"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/pb"
	"github.com/yanshicheng/kube-onec/common/handler/errorx"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOnecClusterByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetOnecClusterByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOnecClusterByIdLogic {
	return &GetOnecClusterByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetOnecClusterByIdLogic) GetOnecClusterById(in *pb.GetOnecClusterByIdReq) (*pb.GetOnecClusterByIdResp, error) {
	cluster, err := l.svcCtx.ClusterModel.FindOne(l.ctx, in.Id)
	if err != nil {
		l.Logger.Errorf("获取集群信息失败: %v", err)
		return nil, errorx.DatabaseFindErr
	}
	clusterInfo := ConvertModelToPBCluster(cluster)

	totalInfo, err := l.svcCtx.NodeModel.FindOneClusterTotalInfo(l.ctx, cluster.Id)
	if err != nil {
		l.Logger.Errorf("获取集群节点信息失败: %v", err)
		return nil, errorx.DatabaseFindErr
	}
	return &pb.GetOnecClusterByIdResp{
		Data: clusterInfo,
		OtherInfo: &pb.OtherInfo{
			NodeTotal:   totalInfo.TotalNode,
			CpuTotal:    totalInfo.TotalCpu,
			MemoryTotal: totalInfo.TotalMemory,
			PodTotal:    totalInfo.TotalPods,
		},
	}, nil
}

// ConvertModelToPBCluster 将 model.OnecCluster 转换为 pb.OnecCluster
func ConvertModelToPBCluster(cluster *model.OnecCluster) *pb.OnecCluster {
	// 创建 pb.OnecCluster 对象并填充字段
	clusterInfo := &pb.OnecCluster{
		Id:                cluster.Id,
		Name:              cluster.Name,
		SkipInsecure:      cluster.SkipInsecure,
		Host:              cluster.Host,
		ConnType:          pb.OnecClusterConnType(pb.OnecClusterConnType_value[cluster.ConnType]),
		EnvTag:            lib.EnvTagToString(cluster.EnvTag),
		Status:            lib.StatusTagName(cluster.Status),
		Version:           cluster.Version,
		Commit:            cluster.Commit,
		Platform:          cluster.Platform,
		VersionBuildTime:  cluster.VersionBuildTime.Unix(),
		ClusterCreateTime: cluster.ClusterCreateTime.Unix(),
		Description:       cluster.Description,
		CreateBy:          cluster.CreateBy,
		UpdateBy:          cluster.UpdateBy,
		CreateTime:        cluster.CreateTime.Unix(),
		UpdateTime:        cluster.UpdateTime.Unix(),
	}

	// 返回转换后的 pb.OnecCluster 对象
	return clusterInfo
}
