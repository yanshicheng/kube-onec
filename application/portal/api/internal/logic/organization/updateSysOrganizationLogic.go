package organization

import (
	"context"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/client/sysorganizationservice"
	"github.com/yanshicheng/kube-onec/common/handler/errorx"

	"github.com/yanshicheng/kube-onec/application/portal/api/internal/svc"
	"github.com/yanshicheng/kube-onec/application/portal/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateSysOrganizationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateSysOrganizationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateSysOrganizationLogic {
	return &UpdateSysOrganizationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateSysOrganizationLogic) UpdateSysOrganization(req *types.UpdateSysOrganizationRequest) (resp string, err error) {
	_, err = l.svcCtx.SysOrganizationRpc.UpdateSysOrganization(l.ctx, &sysorganizationservice.UpdateSysOrganizationReq{
		Id:          req.Id,
		Name:        req.Name,
		Description: req.Description,
	})
	if err != nil {
		logx.WithContext(l.ctx).Errorf("更新职位信息失败: 职位ID=%d, 错误=%v", req.Id, err)
		return "", errorx.DatabaseUpdateErr
	}
	logx.WithContext(l.ctx).Infof("更新职位信息成功: 职位ID=%d, 职位名称=%s, 职位描述=%s", req.Id, req.Name, req.Description)
	return "更新成功!", nil
}
