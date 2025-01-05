package syspermissionservicelogic

import (
	"context"
	"errors"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/internal/model"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/internal/svc"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/pb"
	"github.com/yanshicheng/kube-onec/common/handler/errorx"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateSysPermissionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateSysPermissionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateSysPermissionLogic {
	return &UpdateSysPermissionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateSysPermissionLogic) UpdateSysPermission(in *pb.UpdateSysPermissionReq) (*pb.UpdateSysPermissionResp, error) {
	// 查询权限
	m, err := l.svcCtx.SysPermission.FindOne(l.ctx, in.Id)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			l.Logger.Infof("权限未找到: 权限ID=%d", in.Id)
			return nil, errorx.DatabaseQueryErr
		}
		l.Logger.Errorf("查询权限信息失败: 权限ID=%d, 错误=%v", in.Id, err)
		return nil, errorx.DatabaseQueryErr
	}
	if in.Name != "" {
		m.Name = in.Name
	}
	if in.Uri != "" {
		m.Uri = in.Uri
	}
	if in.Action != "" {
		m.Action = in.Action
	}
	err = l.svcCtx.SysPermission.Update(l.ctx, m)
	if err != nil {
		l.Logger.Errorf("更新权限信息失败: 权限ID=%d, 错误=%v", in.Id, err)
		return nil, errorx.DatabaseUpdateErr
	}
	return &pb.UpdateSysPermissionResp{}, nil
}
