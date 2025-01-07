package onecclusterservicelogic

import (
	"context"
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
	clusterInfo.Token = cluster.Token
	return &pb.GetOnecClusterByIdResp{
		Data: clusterInfo,
	}, nil
}

// ConvertModelToPBCluster 将 model.OnecCluster 转换为 pb.OnecCluster
func ConvertModelToPBCluster(cluster *model.OnecCluster) *pb.OnecCluster {
	// 创建 pb.OnecCluster 对象并填充字段
	//
	clusterInfo := &pb.OnecCluster{
		Id:               cluster.Id,
		Name:             cluster.Name,
		SkipInsecure:     cluster.SkipInsecure,
		Host:             cluster.Host,
		ConnCode:         cluster.ConnCode,
		EnvCode:          cluster.EnvCode,
		Status:           cluster.Status,
		Version:          cluster.Version,
		Commit:           cluster.Commit,
		Platform:         cluster.Platform,
		VersionBuildAt:   cluster.VersionBuildAt.Unix(),
		ClusterCreatedAt: cluster.ClusterCreatedAt.Unix(),
		Description:      cluster.Description,
		CreatedBy:        cluster.CreatedBy,
		UpdatedBy:        cluster.UpdatedBy,
		CreatedAt:        cluster.CreatedAt.Unix(),
		UpdatedAt:        cluster.UpdatedAt.Unix(),
		Uuid:             cluster.Uuid,
		NodeLbIp:         cluster.NodeLbIp,
		Location:         cluster.Location,
		NodeCount:        cluster.NodeCount,
		CpuTotal:         cluster.CpuTotal,
		CpuUsed:          cluster.CpuUsed,
		MemoryTotal:      cluster.MemoryTotal,
		MemoryUsed:       cluster.MemoryUsed,
		PodTotal:         cluster.PodTotal,
		PodUsed:          cluster.PodUsed,
	}

	// 返回转换后的 pb.OnecCluster 对象
	return clusterInfo
}
