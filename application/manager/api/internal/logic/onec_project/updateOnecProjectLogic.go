package onec_project

import (
	"context"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/client/onecprojectservice"
	"github.com/yanshicheng/kube-onec/utils"

	"github.com/yanshicheng/kube-onec/application/manager/api/internal/svc"
	"github.com/yanshicheng/kube-onec/application/manager/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateOnecProjectLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateOnecProjectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateOnecProjectLogic {
	return &UpdateOnecProjectLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateOnecProjectLogic) UpdateOnecProject(req *types.UpdateOnecProjectRequest) (resp string, err error) {
	_, err = l.svcCtx.ProjectRpc.UpdateOnecProject(l.ctx, &onecprojectservice.UpdateOnecProjectReq{
		Id:          req.Id,
		Name:        req.Name,
		Description: req.Description,
		UpdatedBy:   utils.GetAccount(l.ctx),
	})
	if err != nil {
		l.Logger.Errorf("更新项目失败: %v", err)
		return "", err
	}
	return "更新项目成功!", nil
}
