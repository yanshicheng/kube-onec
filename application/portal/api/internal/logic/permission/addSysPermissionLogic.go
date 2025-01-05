package permission

import (
	"context"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/pb"

	"github.com/yanshicheng/kube-onec/application/portal/api/internal/svc"
	"github.com/yanshicheng/kube-onec/application/portal/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddSysPermissionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddSysPermissionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddSysPermissionLogic {
	return &AddSysPermissionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddSysPermissionLogic) AddSysPermission(req *types.AddSysPermissionRequest) (resp string, err error) {
	_, err = l.svcCtx.SysPermissionRpc.AddSysPermission(l.ctx, &pb.AddSysPermissionReq{
		Action:   req.Action,
		Name:     req.Name,
		ParentId: req.ParentId,
		Uri:      req.Uri,
	})
	if err != nil {
		l.Logger.Errorf("添加权限失败: %v", err)
		return "", err
	}
	return "权限添加成功!", nil
}
