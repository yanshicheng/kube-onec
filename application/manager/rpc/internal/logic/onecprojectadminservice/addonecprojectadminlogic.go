package onecprojectadminservicelogic

import (
	"context"

	"github.com/yanshicheng/kube-onec/application/manager/rpc/internal/svc"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddOnecProjectAdminLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddOnecProjectAdminLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddOnecProjectAdminLogic {
	return &AddOnecProjectAdminLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// -----------------------项目管理员表，关联项目与用户的多对多关系-----------------------
func (l *AddOnecProjectAdminLogic) AddOnecProjectAdmin(in *pb.AddOnecProjectAdminReq) (*pb.AddOnecProjectAdminResp, error) {
	// todo: add your logic here and delete this line

	return &pb.AddOnecProjectAdminResp{}, nil
}
