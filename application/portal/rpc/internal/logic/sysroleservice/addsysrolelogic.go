package sysroleservicelogic

import (
	"context"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/internal/model"
	"github.com/yanshicheng/kube-onec/common/handler/errorx"

	"github.com/yanshicheng/kube-onec/application/portal/rpc/internal/svc"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddSysRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddSysRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddSysRoleLogic {
	return &AddSysRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// -----------------------角色表-----------------------
func (l *AddSysRoleLogic) AddSysRole(in *pb.AddSysRoleReq) (*pb.AddSysRoleResp, error) {
	// todo: add your logic here and delete this line
	// 通过 token 获取用户名
	_, err := l.svcCtx.SysRole.Insert(l.ctx, &model.SysRole{
		CreatedBy:    in.CreatedBy,
		RoleCode:    in.RoleCode,
		Description: in.Description,
		RoleName:    in.RoleName,
		UpdatedBy:    in.UpdatedBy,
	})
	if err != nil {
		l.Logger.Errorf("添加角色失败: %v", err)
		return nil, errorx.DatabaseCreateErr
	}
	return &pb.AddSysRoleResp{}, nil
}
