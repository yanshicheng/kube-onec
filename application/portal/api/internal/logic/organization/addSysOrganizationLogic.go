package organization

import (
	"context"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/client/sysorganizationservice"

	"github.com/yanshicheng/kube-onec/application/portal/api/internal/svc"
	"github.com/yanshicheng/kube-onec/application/portal/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddSysOrganizationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddSysOrganizationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddSysOrganizationLogic {
	return &AddSysOrganizationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddSysOrganizationLogic) AddSysOrganization(req *types.AddSysOrganizationRequest) (resp string, err error) {
	res, err := l.svcCtx.SysOrganizationRpc.AddSysOrganization(l.ctx, &sysorganizationservice.AddSysOrganizationReq{
		Name:        req.Name,
		ParentId:    req.ParentId,
		Description: req.Description,
	})
	if err != nil {
		logx.WithContext(l.ctx).Errorf("添加机构失败: 请求=%+v, 错误=%v", req, err)
		return "", err
	}
	logx.WithContext(l.ctx).Infof("添加机构成功: 请求=%+v, 响应=%s", req, res)
	return "添加机构成功!", nil
}
