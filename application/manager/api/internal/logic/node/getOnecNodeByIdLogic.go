package node

import (
	"context"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/client/onecnodeservice"

	"github.com/yanshicheng/kube-onec/application/manager/api/internal/svc"
	"github.com/yanshicheng/kube-onec/application/manager/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOnecNodeByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetOnecNodeByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOnecNodeByIdLogic {
	return &GetOnecNodeByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetOnecNodeByIdLogic) GetOnecNodeById(req *types.DefaultIdRequest) (resp *types.OnecNode, err error) {
	res, err := l.svcCtx.NodeRpc.GetOnecNodeById(l.ctx, &onecnodeservice.GetOnecNodeByIdReq{
		Id: req.Id,
	})
	if err != nil {
		l.Logger.Errorf("获取节点信息失败: %v", err)
		return nil, err
	}
	// 映射 RPC 返回的数据到 API 响应结构体
	l.Logger.Infof("获取节点信息成功: %v", res.Data.NodeName)

	return &types.OnecNode{
		Id:               res.Data.Id,
		ClusterUuid:      res.Data.ClusterUuid,
		NodeName:         res.Data.NodeName,
		Cpu:              res.Data.Cpu,
		Memory:           res.Data.Memory,
		MaxPods:          res.Data.MaxPods,
		IsGpu:            res.Data.IsGpu,
		NodeUid:          res.Data.NodeUid,
		Status:           res.Data.Status,
		Roles:            res.Data.Roles,
		JoinAt:           res.Data.JoinAt,
		PodCidr:          res.Data.PodCidr,
		Unschedulable:    res.Data.Unschedulable,
		NodeIp:           res.Data.NodeIp,
		Os:               res.Data.Os,
		KernelVersion:    res.Data.KernelVersion,
		ContainerRuntime: res.Data.ContainerRuntime,
		KubeletVersion:   res.Data.KubeletVersion,
		KubeletPort:      res.Data.KubeletPort,
		OperatingSystem:  res.Data.OperatingSystem,
		Architecture:     res.Data.Architecture,
		CreatedBy:        res.Data.CreatedBy,
		UpdatedBy:        res.Data.UpdatedBy,
		CreatedAt:        res.Data.CreatedAt,
		UpdatedAt:        res.Data.UpdatedAt,
	}, nil
}
