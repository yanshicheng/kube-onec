package onecprojectservicelogic

import (
	"context"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/internal/code"

	"github.com/yanshicheng/kube-onec/application/manager/rpc/internal/svc"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateOnecProjectLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateOnecProjectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateOnecProjectLogic {
	return &UpdateOnecProjectLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateOnecProjectLogic) UpdateOnecProject(in *pb.UpdateOnecProjectReq) (*pb.UpdateOnecProjectResp, error) {
	res, err := l.svcCtx.ProjectModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}
	if in.Name != "" {
		res.Name = in.Name
	}
	if in.Description != "" {
		res.Description = in.Description
	}
	res.UpdatedBy = in.UpdatedBy
	err = l.svcCtx.ProjectModel.Update(l.ctx, res)
	if err != nil {
		l.Logger.Errorf("更新项目失败: %v", err)
		return nil, code.UpdateProjectErr
	}
	l.Logger.Infof("更新项目成功: %v", res)
	return &pb.UpdateOnecProjectResp{}, nil
}
