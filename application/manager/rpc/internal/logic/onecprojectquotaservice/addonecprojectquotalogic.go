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
	// todo: add your logic here and delete this line

	return &pb.AddOnecProjectQuotaResp{}, nil
}
