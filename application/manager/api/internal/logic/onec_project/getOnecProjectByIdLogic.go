package onec_project

import (
	"context"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/client/onecprojectservice"

	"github.com/yanshicheng/kube-onec/application/manager/api/internal/svc"
	"github.com/yanshicheng/kube-onec/application/manager/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOnecProjectByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetOnecProjectByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOnecProjectByIdLogic {
	return &GetOnecProjectByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetOnecProjectByIdLogic) GetOnecProjectById(req *types.DefaultIdRequest) (resp *types.OnecProject, err error) {
	res, err := l.svcCtx.ProjectRpc.GetOnecProjectById(l.ctx, &onecprojectservice.GetOnecProjectByIdReq{
		Id: req.Id,
	})
	if err != nil {
		l.Logger.Errorf("查询项目失败: %v", err)
		return nil, err
	}

	return &types.OnecProject{
		Id:          res.Data.Id,
		Name:        res.Data.Name,
		Identifier:  res.Data.Identifier,
		Description: res.Data.Description,
		IsDefault:   res.Data.IsDefault,
		CreatedAt:   res.Data.CreatedAt,
		UpdatedAt:   res.Data.UpdatedAt,
		CreatedBy:   res.Data.CreatedBy,
		UpdatedBy:   res.Data.UpdatedBy,
	}, nil
}
