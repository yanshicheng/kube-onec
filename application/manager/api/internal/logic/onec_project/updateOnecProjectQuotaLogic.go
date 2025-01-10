package onec_project

import (
	"context"
	"github.com/yanshicheng/kube-onec/application/manager/api/internal/svc"
	"github.com/yanshicheng/kube-onec/application/manager/api/internal/types"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/pb"
	"github.com/yanshicheng/kube-onec/pkg/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateOnecProjectQuotaLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateOnecProjectQuotaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateOnecProjectQuotaLogic {
	return &UpdateOnecProjectQuotaLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateOnecProjectQuotaLogic) UpdateOnecProjectQuota(req *types.UpdateOnecProjectQuotaRequest) (resp string, err error) {
	_, err = l.svcCtx.ProjectQuotaRpc.UpdateOnecProjectQuota(l.ctx, &pb.UpdateOnecProjectQuotaReq{
		Id:               req.Id,
		ClusterUuid:      req.ClusterUuid,
		CpuQuota:         req.CpuQuota,
		CpuOvercommit:    req.CpuOvercommit,
		MemoryQuota:      req.MemoryQuota,
		MemoryOvercommit: req.MemoryOvercommit,
		PodLimit:         req.PodLimit,
		PvcLimit:         req.PvcLimit,
		ConfigmapLimit:   req.ConfigmapLimit,
		SecretLimit:      req.SecretLimit,
		NodeportLimit:    req.NodeportLimit,
		StorageLimit:     req.StorageLimit,
		CpuLimit:         req.CpuLimit,
		MemoryLimit:      req.MemoryLimit,
		ProjectId:        req.ProjectId,
		UpdatedBy:        utils.GetAccount(l.ctx),
	})
	if err != nil {
		l.Logger.Errorf("更新项目配额失败: %v", err)
		return "", err
	}
	return "更新项目配额成功!", nil
}
