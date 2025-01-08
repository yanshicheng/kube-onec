package onecprojectservicelogic

import (
	"context"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/internal/code"

	"github.com/yanshicheng/kube-onec/application/manager/rpc/internal/svc"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOnecProjectByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetOnecProjectByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOnecProjectByIdLogic {
	return &GetOnecProjectByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetOnecProjectByIdLogic) GetOnecProjectById(in *pb.GetOnecProjectByIdReq) (*pb.GetOnecProjectByIdResp, error) {
	res, err := l.svcCtx.ProjectModel.FindOne(l.ctx, in.Id)
	if err != nil {
		l.Logger.Errorf("查询项目失败: %v", err)
		return nil, code.GetProjectErr
	}
	return &pb.GetOnecProjectByIdResp{
		Data: &pb.OnecProject{
			Id:          res.Id,
			Name:        res.Name,
			Identifier:  res.Identifier,
			Description: res.Description,
			CreatedAt:   res.CreatedAt.Unix(),
			UpdatedAt:   res.UpdatedAt.Unix(),
			CreatedBy:   res.CreatedBy,
			UpdatedBy:   res.UpdatedBy,
		},
	}, nil
}
