package cluster

import (
	"github.com/yanshicheng/kube-onec/application/manager/api/internal/types"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/client/onecclusterservice"
)

func convertCluster(data *onecclusterservice.OnecCluster) types.OnecCluster {
	return types.OnecCluster{
		Id:               data.Id,
		Name:             data.Name,
		UUID:             data.Uuid,
		SkipInsecure:     data.SkipInsecure,
		Host:             data.Host,
		EnvName:          data.EnvName,
		ConnCode:         data.ConnCode,
		EnvCode:          data.EnvCode,
		Status:           data.Status,
		Version:          data.Version,
		Commit:           data.Commit,
		Platform:         data.Platform,
		VersionBuildAt:   data.VersionBuildAt,
		ClusterCreatedAt: data.ClusterCreatedAt,
		NodeCount:        data.NodeCount,
		CpuTotal:         data.CpuTotal,
		MemoryTotal:      data.MemoryTotal,
		PodTotal:         data.PodTotal,
		CpuUsed:          data.CpuUsed,
		MemoryUsed:       data.MemoryUsed,
		PodUsed:          data.PodUsed,
		Location:         data.Location,
		NodeLbIp:         data.NodeLbIp,
		Description:      data.Description,
		CreatedBy:        data.CreatedBy,
		UpdatedBy:        data.UpdatedBy,
		CreatedAt:        data.CreatedAt,
		UpdatedAt:        data.UpdatedAt,
		Token:            data.Token,
	}
}

// 辅助函数：将 RPC 的集群信息转换为 API 返回的集群信息
func convertClusters(data []*onecclusterservice.OnecCluster) []types.OnecCluster {
	var clusters []types.OnecCluster
	for _, c := range data {
		clusters = append(clusters, types.OnecCluster{
			Id:               c.Id,
			Name:             c.Name,
			UUID:             c.Uuid,
			SkipInsecure:     c.SkipInsecure,
			Host:             c.Host,
			EnvName:          c.EnvName,
			ConnCode:         c.ConnCode,
			EnvCode:          c.EnvCode,
			Status:           c.Status,
			Version:          c.Version,
			Commit:           c.Commit,
			Platform:         c.Platform,
			VersionBuildAt:   c.VersionBuildAt,
			ClusterCreatedAt: c.ClusterCreatedAt,
			NodeCount:        c.NodeCount,
			CpuTotal:         c.CpuTotal,
			MemoryTotal:      c.MemoryTotal,
			PodTotal:         c.PodTotal,
			CpuUsed:          c.CpuUsed,
			MemoryUsed:       c.MemoryUsed,
			PodUsed:          c.PodUsed,
			Location:         c.Location,
			NodeLbIp:         c.NodeLbIp,
			Description:      c.Description,
			CreatedBy:        c.CreatedBy,
			UpdatedBy:        c.UpdatedBy,
			CreatedAt:        c.CreatedAt,
			UpdatedAt:        c.UpdatedAt,
			Token:            c.Token,
		})
	}
	return clusters
}
