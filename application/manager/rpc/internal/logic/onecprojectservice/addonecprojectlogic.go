package onecprojectservicelogic

import (
	"context"

	"github.com/yanshicheng/kube-onec/application/manager/rpc/internal/svc"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddOnecProjectLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddOnecProjectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddOnecProjectLogic {
	return &AddOnecProjectLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// -----------------------项目表，记录项目信息-----------------------
func (l *AddOnecProjectLogic) AddOnecProject(in *pb.AddOnecProjectReq) (*pb.AddOnecProjectResp, error) {
	// todo: add your logic here and delete this line

	return &pb.AddOnecProjectResp{}, nil
}
