package onec_project

import (
	"context"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/client/onecprojectservice"
	"github.com/yanshicheng/kube-onec/utils"

	"github.com/yanshicheng/kube-onec/application/manager/api/internal/svc"
	"github.com/yanshicheng/kube-onec/application/manager/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddOnecProjectLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddOnecProjectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddOnecProjectLogic {
	return &AddOnecProjectLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddOnecProjectLogic) AddOnecProject(req *types.AddOnecProjectRequest) (resp string, err error) {
	_, err = l.svcCtx.ProjectRpc.AddOnecProject(l.ctx, &onecprojectservice.AddOnecProjectReq{
		Name:        req.Name,
		Identifier:  req.Identifier,
		Description: req.Description,
		CreatedBy:   utils.GetAccount(l.ctx),
		UpdatedBy:   utils.GetAccount(l.ctx),
	})
	if err != nil {
		l.Logger.Errorf("添加项目失败: %v", err)
		return
	}
	return "添加项目成功!", nil
}
