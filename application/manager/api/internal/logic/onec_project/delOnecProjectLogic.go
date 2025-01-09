package onec_project

import (
	"context"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/client/onecprojectservice"

	"github.com/yanshicheng/kube-onec/application/manager/api/internal/svc"
	"github.com/yanshicheng/kube-onec/application/manager/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelOnecProjectLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDelOnecProjectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelOnecProjectLogic {
	return &DelOnecProjectLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DelOnecProjectLogic) DelOnecProject(req *types.DefaultIdRequest) (resp string, err error) {
	_, err = l.svcCtx.ProjectRpc.DelOnecProject(l.ctx, &onecprojectservice.DelOnecProjectReq{
		Id: req.Id,
	})
	if err != nil {
		l.Logger.Errorf("删除项目失败: %v", err)
		return "", err
	}
	return "删除项目成功!", nil
}
