package onec_project

import (
	"context"
	"github.com/yanshicheng/kube-onec/application/manager/api/internal/svc"
	"github.com/yanshicheng/kube-onec/application/manager/api/internal/types"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/pb"
	"github.com/yanshicheng/kube-onec/pkg/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddOnecProjectQuotaLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddOnecProjectQuotaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddOnecProjectQuotaLogic {
	return &AddOnecProjectQuotaLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddOnecProjectQuotaLogic) AddOnecProjectQuota(req *types.AddOnecProjectQuotaRequest) (resp string, err error) {
	res, err := l.svcCtx.ProjectQuotaRpc.AddOnecProjectQuota(l.ctx, &pb.AddOnecProjectQuotaReq{
		ClusterUuid:      req.ClusterUuid,
		ProjectId:        req.ProjectId,
		CpuQuota:         req.CpuQuota,
		CpuOvercommit:    req.CpuOvercommit,
		CpuLimit:         req.CpuLimit,
		MemoryQuota:      req.MemoryQuota,
		MemoryOvercommit: req.MemoryOvercommit,
		MemoryLimit:      req.MemoryLimit,
		StorageLimit:     req.StorageLimit,
		ConfigmapLimit:   req.ConfigmapLimit,
		SecretLimit:      req.SecretLimit,
		PvcLimit:         req.PvcLimit,
		PodLimit:         req.PodLimit,
		NodeportLimit:    req.NodeportLimit,
		CreatedBy:        utils.GetAccount(l.ctx),
		UpdatedBy:        utils.GetAccount(l.ctx),
	})
	if err != nil {
		l.Logger.Errorf("添加项目配额失败: %v", err)
		return
	}
	l.Logger.Infof("添加项目配额成功: %v", res)
	return "添加项目配额成功!", nil
}
