package permission

import (
	"context"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/client/syspermissionservice"

	"github.com/yanshicheng/kube-onec/application/portal/api/internal/svc"
	"github.com/yanshicheng/kube-onec/application/portal/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateSysPermissionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateSysPermissionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateSysPermissionLogic {
	return &UpdateSysPermissionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateSysPermissionLogic) UpdateSysPermission(req *types.UpdateSysPermissionRequest) (resp string, err error) {
	_, err = l.svcCtx.SysPermissionRpc.UpdateSysPermission(l.ctx, &syspermissionservice.UpdateSysPermissionReq{
		Id:     req.Id,
		Name:   req.Name,
		Uri:    req.Uri,
		Action: req.Action,
	})
	if err != nil {
		l.Logger.Errorf("更新权限失败: %v", err)
		return "", err
	}
	return "更新权限成功!", nil
}
