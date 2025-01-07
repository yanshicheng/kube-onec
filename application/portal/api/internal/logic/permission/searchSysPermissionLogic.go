package permission

import (
	"context"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/client/syspermissionservice"

	"github.com/yanshicheng/kube-onec/application/portal/api/internal/svc"
	"github.com/yanshicheng/kube-onec/application/portal/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchSysPermissionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchSysPermissionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchSysPermissionLogic {
	return &SearchSysPermissionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchSysPermissionLogic) SearchSysPermission(req *types.SearchSysPermissionRequest) (resp *types.SearchSysPermissionResponse, err error) {
	res, err := l.svcCtx.SysPermissionRpc.SearchSysPermission(l.ctx, &syspermissionservice.SearchSysPermissionReq{
		Action:   req.Action,
		IsAsc:    req.IsAsc,
		Name:     req.Name,
		OrderStr: req.OrderStr,
		Page:     req.Page,
		PageSize: req.PageSize,
		ParentId: req.ParentId,
		Uri:      req.Uri,
	})
	if err != nil {
		l.Logger.Errorf("查询权限失败: %v", err)
		return nil, err
	}
	data := make([]types.SysPermission, len(res.Data))
	for i, v := range res.Data {
		data[i] = types.SysPermission{
			Id:        v.Id,
			Action:    v.Action,
			Name:      v.Name,
			ParentId:  v.ParentId,
			Uri:       v.Uri,
			Level:     v.Level,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		}
	}
	resp = &types.SearchSysPermissionResponse{
		Items: data,
		Total: res.Total,
	}
	return
}
