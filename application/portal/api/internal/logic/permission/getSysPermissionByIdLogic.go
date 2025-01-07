package permission

import (
	"context"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/client/syspermissionservice"

	"github.com/yanshicheng/kube-onec/application/portal/api/internal/svc"
	"github.com/yanshicheng/kube-onec/application/portal/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSysPermissionByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetSysPermissionByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSysPermissionByIdLogic {
	return &GetSysPermissionByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSysPermissionByIdLogic) GetSysPermissionById(req *types.DefaultIdRequest) (resp *types.SysPermission, err error) {
	res, err := l.svcCtx.SysPermissionRpc.GetSysPermissionById(l.ctx, &syspermissionservice.GetSysPermissionByIdReq{
		Id: req.Id,
	})
	if err != nil {
		l.Logger.Errorf("获取权限失败: %v", err)
		return
	}
	resp = &types.SysPermission{
		Action:    res.Data.Action,
		CreatedAt: res.Data.CreatedAt,
		Id:        res.Data.Id,
		Level:     res.Data.Level,
		Name:      res.Data.Name,
		ParentId:  res.Data.ParentId,
		UpdatedAt: res.Data.UpdatedAt,
		Uri:       res.Data.Uri,
	}
	return
}
