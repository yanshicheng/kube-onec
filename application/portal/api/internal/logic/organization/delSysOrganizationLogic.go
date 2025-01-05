package organization

import (
	"context"
	"github.com/yanshicheng/kube-onec/application/portal/api/internal/svc"
	"github.com/yanshicheng/kube-onec/application/portal/api/internal/types"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/client/sysorganizationservice"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelSysOrganizationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDelSysOrganizationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelSysOrganizationLogic {
	return &DelSysOrganizationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DelSysOrganizationLogic) DelSysOrganization(req *types.DefaultIdRequest) (resp string, err error) {
	res, err := l.svcCtx.SysOrganizationRpc.DelSysOrganization(l.ctx, &sysorganizationservice.DelSysOrganizationReq{
		Id: req.Id,
	})
	if err != nil {
		logx.WithContext(l.ctx).Errorf("删除机构失败: %v", err)
		return "", err
	}
	logx.WithContext(l.ctx).Infof("删除机构成功: %v", res)
	return "删除机构成功!", nil
}
