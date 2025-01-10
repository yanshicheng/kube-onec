package onecprojectapplicationservicelogic

import (
	"context"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/internal/code"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/internal/shared"

	"github.com/yanshicheng/kube-onec/application/manager/rpc/internal/svc"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOnecProjectApplicationByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetOnecProjectApplicationByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOnecProjectApplicationByIdLogic {
	return &GetOnecProjectApplicationByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetOnecProjectApplicationByIdLogic) GetOnecProjectApplicationById(in *pb.GetOnecProjectApplicationByIdReq) (*pb.GetOnecProjectApplicationByIdResp, error) {
	res, err := l.svcCtx.ProjectApplicationModel.FindOne(l.ctx, in.Id)
	if err != nil {
		l.Logger.Errorf("查询应用失败: %v", err)
		return nil, code.GetApplicationErr
	}
	l.Logger.Infof("查询应用成功: %v", res)

	return &pb.GetOnecProjectApplicationByIdResp{
		Data: shared.ConvertModelToPbApplication(res),
	}, nil
}
