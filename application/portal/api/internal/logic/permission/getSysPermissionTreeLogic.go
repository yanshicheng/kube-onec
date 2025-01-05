package permission

import (
	"context"

	"github.com/yanshicheng/kube-onec/application/portal/api/internal/svc"
	"github.com/yanshicheng/kube-onec/application/portal/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSysPermissionTreeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetSysPermissionTreeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSysPermissionTreeLogic {
	return &GetSysPermissionTreeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSysPermissionTreeLogic) GetSysPermissionTree() (resp []types.SysPermissionTreeNode, err error) {
	// todo: add your logic here and delete this line

	return
}
