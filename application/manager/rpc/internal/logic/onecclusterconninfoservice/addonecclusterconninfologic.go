package onecclusterconninfoservicelogic

import (
	"context"

	"github.com/yanshicheng/kube-onec/application/manager/rpc/internal/svc"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddOnecClusterConnInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddOnecClusterConnInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddOnecClusterConnInfoLogic {
	return &AddOnecClusterConnInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// -----------------------通用的服务连接信息表，动态支持多个服务-----------------------
func (l *AddOnecClusterConnInfoLogic) AddOnecClusterConnInfo(in *pb.AddOnecClusterConnInfoReq) (*pb.AddOnecClusterConnInfoResp, error) {
	// todo: add your logic here and delete this line

	return &pb.AddOnecClusterConnInfoResp{}, nil
}
