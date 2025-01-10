package shared

import (
	"github.com/yanshicheng/kube-onec/application/manager/rpc/internal/model"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/pb"
)

// ConvertDbModelToPbModel 将数据库模型转换为 Protobuf 模型
func ConvertDbModelToPbModel(dbNode *model.OnecNode) *pb.OnecNode {
	return &pb.OnecNode{
		Id:               dbNode.Id,
		ClusterUuid:      dbNode.ClusterUuid,
		NodeName:         dbNode.NodeName,
		Cpu:              dbNode.Cpu,
		Memory:           dbNode.Memory,
		MaxPods:          dbNode.MaxPods,
		IsGpu:            dbNode.IsGpu,
		NodeUid:          dbNode.NodeUid,
		Status:           dbNode.Status,
		Roles:            dbNode.Roles,
		JoinAt:           dbNode.JoinAt.Unix(), // 转换为 Unix 时间戳
		PodCidr:          dbNode.PodCidr,
		Unschedulable:    dbNode.Unschedulable,
		NodeIp:           dbNode.NodeIp,
		Os:               dbNode.Os,
		KernelVersion:    dbNode.KernelVersion,
		ContainerRuntime: dbNode.ContainerRuntime,
		KubeletVersion:   dbNode.KubeletVersion,
		KubeletPort:      int64(dbNode.KubeletPort), // uint64 转为 int64
		OperatingSystem:  dbNode.OperatingSystem,
		Architecture:     dbNode.Architecture,
		CreatedBy:        dbNode.CreatedBy,
		UpdatedBy:        dbNode.UpdatedBy,
		CreatedAt:        dbNode.CreatedAt.Unix(), // 转换为 Unix 时间戳
		UpdatedAt:        dbNode.UpdatedAt.Unix(), // 转换为 Unix 时间戳
	}
}

func ConvertModelToPbApplication(modelApp *model.OnecProjectApplication) *pb.OnecProjectApplication {
	return &pb.OnecProjectApplication{
		Id:             modelApp.Id,
		ProjectId:      modelApp.ProjectId,
		ClusterUuid:    modelApp.ClusterUuid,
		Name:           modelApp.Name,
		Identifier:     modelApp.Identifier,
		Uuid:           modelApp.Uuid,
		Description:    modelApp.Description,
		CpuLimit:       modelApp.CpuLimit,
		MemoryLimit:    modelApp.MemoryLimit,
		StorageLimit:   modelApp.StorageLimit,
		ConfigmapLimit: modelApp.ConfigmapLimit,
		PvcLimit:       modelApp.PvcLimit,
		PodLimit:       modelApp.PodLimit,
		NodeportLimit:  modelApp.NodeportLimit,
		SecretLimit:    modelApp.SecretLimit,
		ServiceLimit:   modelApp.ServiceLimit,
		Status:         modelApp.Status,
		AppCreatedAt:   modelApp.AppCreateTime.Unix(), // 转换为 Unix 时间戳
		CreatedBy:      modelApp.CreatedBy,
		UpdatedBy:      modelApp.UpdatedBy,
		CreatedAt:      modelApp.CreatedAt.Unix(), // 转换为 Unix 时间戳
		UpdatedAt:      modelApp.UpdatedAt.Unix(), // 转换为 Unix 时间戳
	}
}
