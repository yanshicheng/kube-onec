package onecprojectquotaservicelogic

import (
	"context"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/internal/svc"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddOnecProjectQuotaLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddOnecProjectQuotaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddOnecProjectQuotaLogic {
	return &AddOnecProjectQuotaLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// -----------------------项目与集群的对应关系表，记录资源配额和使用情况-----------------------
func (l *AddOnecProjectQuotaLogic) AddOnecProjectQuota(in *pb.AddOnecProjectQuotaReq) (*pb.AddOnecProjectQuotaResp, error) {
	//cluster, err := l.svcCtx.ClusterModel.FindOneByUuid(l.ctx, in.ClusterUuid)
	//if err != nil {
	//	l.Logger.Errorf("获取集群信息失败: %v", err)
	//	return nil, code.GetClusterInfoErr
	//}
	//
	//project, err := l.svcCtx.ProjectModel.F(l.ctx, in.ProjectId)

	return &pb.AddOnecProjectQuotaResp{}, nil
}
