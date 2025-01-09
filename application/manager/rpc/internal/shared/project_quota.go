package shared

import (
	"github.com/yanshicheng/kube-onec/application/manager/rpc/internal/model"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/pb"
)

func ConvertToPBQuota(modeQuota model.OnecProjectQuota) *pb.OnecProjectQuota {
	return &pb.OnecProjectQuota{
		Id:                   modeQuota.Id,
		ClusterUuid:          modeQuota.ClusterUuid,
		ProjectId:            modeQuota.ProjectId,
		CpuQuota:             modeQuota.CpuQuota,
		CpuOvercommit:        modeQuota.CpuOvercommit,
		CpuLimit:             modeQuota.CpuLimit,
		CpuUsed:              modeQuota.CpuUsed,
		CpuLimitRemain:       modeQuota.CpuLimit - modeQuota.CpuUsed,
		MemoryQuota:          modeQuota.MemoryQuota,
		MemoryOvercommit:     modeQuota.MemoryOvercommit,
		MemoryLimit:          modeQuota.MemoryLimit,
		MemoryUsed:           modeQuota.MemoryUsed,
		MemoryLimitRemain:    modeQuota.MemoryLimit - modeQuota.MemoryUsed,
		StorageLimit:         modeQuota.StorageLimit,
		StorageUsed:          modeQuota.StorageUsed,
		StorageLimitRemain:   modeQuota.StorageLimit - modeQuota.StorageUsed,
		ConfigmapLimit:       modeQuota.ConfigmapLimit,
		ConfigmapUsed:        modeQuota.ConfigmapUsed,
		ConfigmapLimitRemain: modeQuota.ConfigmapLimit - modeQuota.ConfigmapUsed,
		SecretLimit:          modeQuota.SecretLimit,
		SecretUsed:           modeQuota.SecretUsed,
		SecretLimitRemain:    modeQuota.SecretLimit - modeQuota.SecretUsed,
		PvcLimit:             modeQuota.PvcLimit,
		PvcUsed:              modeQuota.PvcUsed,
		PvcLimitRemain:       modeQuota.PvcLimit - modeQuota.PvcUsed,
		PodLimit:             modeQuota.PodLimit,
		PodUsed:              modeQuota.PodUsed,
		PodLimitRemain:       modeQuota.PodLimit - modeQuota.PodUsed,
		NodeportLimit:        modeQuota.NodeportLimit,
		NodeportUsed:         modeQuota.NodeportUsed,
		NodeportLimitRemain:  modeQuota.NodeportLimit - modeQuota.NodeportUsed,
		Status:               modeQuota.Status,
		CreatedBy:            modeQuota.CreatedBy,
		UpdatedBy:            modeQuota.UpdatedBy,
		CreatedAt:            modeQuota.CreatedAt.Unix(),
		UpdatedAt:            modeQuota.UpdatedAt.Unix(),
	}
}
