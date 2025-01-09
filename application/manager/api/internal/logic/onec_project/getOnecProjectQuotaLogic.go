package onec_project

import (
	"context"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/pb"

	"github.com/yanshicheng/kube-onec/application/manager/api/internal/svc"
	"github.com/yanshicheng/kube-onec/application/manager/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOnecProjectQuotaLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetOnecProjectQuotaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOnecProjectQuotaLogic {
	return &GetOnecProjectQuotaLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetOnecProjectQuotaLogic) GetOnecProjectQuota(req *types.GetOnecProjectQuotaRequest) (resp *types.OnecProjectQuota, err error) {
	res, err := l.svcCtx.ProjectQuotaRpc.GetOnecProjectQuota(l.ctx, &pb.GetOnecProjectQuotaReq{
		ClusterUuid: req.ClusterUuid,
		ProjectId:   req.ProjectId,
	})
	if err != nil {
		l.Logger.Errorf("获取项目配额失败: %v", err)
		return nil, err
	}

	return ConvertPBToTypesQuota(res.Data), nil
}

func ConvertPBToTypesQuota(pbQuota *pb.OnecProjectQuota) *types.OnecProjectQuota {
	return &types.OnecProjectQuota{
		Id:                   pbQuota.Id,
		ClusterUuid:          pbQuota.ClusterUuid,
		ProjectId:            pbQuota.ProjectId,
		CpuQuota:             pbQuota.CpuQuota,
		CpuOvercommit:        pbQuota.CpuOvercommit,
		CpuLimit:             pbQuota.CpuLimit,
		CpuUsed:              pbQuota.CpuUsed,
		CpuLimitRemain:       pbQuota.CpuLimitRemain,
		MemoryQuota:          pbQuota.MemoryQuota,
		MemoryOvercommit:     pbQuota.MemoryOvercommit,
		MemoryLimit:          pbQuota.MemoryLimit,
		MemoryUsed:           pbQuota.MemoryUsed,
		MemoryLimitRemain:    pbQuota.MemoryLimitRemain,
		StorageLimit:         pbQuota.StorageLimit,
		StorageUsed:          pbQuota.StorageUsed,
		StorageLimitRemain:   pbQuota.StorageLimitRemain,
		ConfigmapLimit:       pbQuota.ConfigmapLimit,
		ConfigMapUsed:        pbQuota.ConfigmapUsed,
		ConfigMapLimitRemain: pbQuota.ConfigmapLimitRemain,
		SecretLimit:          pbQuota.SecretLimit,
		SecretUsed:           pbQuota.SecretUsed,
		SecretLimitRemain:    pbQuota.SecretLimitRemain,
		PvcLimit:             pbQuota.PvcLimit,
		PvcUsed:              pbQuota.PvcUsed,
		PvcLimitRemain:       pbQuota.PvcLimitRemain,
		PodLimit:             pbQuota.PodLimit,
		PodUsed:              pbQuota.PodUsed,
		PodLimitRemain:       pbQuota.PodLimitRemain,
		NodeportLimit:        pbQuota.NodeportLimit,
		NodeportUsed:         pbQuota.NodeportUsed,
		NodeportLimitRemain:  pbQuota.NodeportLimitRemain,
		Status:               pbQuota.Status,
		CreatedBy:            pbQuota.CreatedBy,
		UpdatedBy:            pbQuota.UpdatedBy,
		CreatedAt:            pbQuota.CreatedAt,
		UpdatedAt:            pbQuota.UpdatedAt,
	}
}
