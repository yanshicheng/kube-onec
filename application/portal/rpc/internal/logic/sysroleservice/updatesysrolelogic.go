package sysroleservicelogic

import (
	"context"
	"errors"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/internal/model"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/internal/svc"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/pb"
	"github.com/yanshicheng/kube-onec/common/handler/errorx"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateSysRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateSysRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateSysRoleLogic {
	return &UpdateSysRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}
func (l *UpdateSysRoleLogic) UpdateSysRole(in *pb.UpdateSysRoleReq) (*pb.UpdateSysRoleResp, error) {
	// 查询角色信息
	role, err := l.svcCtx.SysRole.FindOne(l.ctx, in.Id)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return nil, errorx.DatabaseNotFound // 返回错误：数据未找到（建议在 errorx 中设置中文错误描述）
		}
		// 查询过程中的其他错误
		l.Logger.Errorf("查询角色信息失败，角色ID=%d，错误信息：%v", in.Id, err)
		return nil, errorx.DatabaseQueryErr // 返回错误：查询数据失败
	}

	// 更新字段
	if in.RoleName != "" {
		role.RoleName = in.RoleName

	}
	if in.Description != "" {
		role.Description = in.Description
	}
	role.UpdateBy = in.UpdateBy
	// 更新角色信息
	if err := l.svcCtx.SysRole.Update(l.ctx, role); err != nil {
		l.Logger.Errorf("更新角色信息失败，角色ID=%d，错误信息：%v", in.Id, err)
		return nil, errorx.DatabaseUpdateErr // 返回错误：更新数据失败
	}

	return &pb.UpdateSysRoleResp{}, nil
}
