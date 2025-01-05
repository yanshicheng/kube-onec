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

type GetSysRoleByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetSysRoleByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSysRoleByIdLogic {
	return &GetSysRoleByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetSysRoleByIdLogic) GetSysRoleById(in *pb.GetSysRoleByIdReq) (*pb.GetSysRoleByIdResp, error) {
	role, err := l.svcCtx.SysRole.FindOne(l.ctx, in.Id)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			l.Logger.Errorf("查询角色失败: %v", err)
			return nil, errorx.DatabaseNotFound
		}
		l.Logger.Errorf("查询角色失败: %v", err)
		return nil, errorx.DatabaseQueryErr
	}
	l.Logger.Infof("查询角色: %v", role)
	return &pb.GetSysRoleByIdResp{
		Data: &pb.SysRole{
			Id:          role.Id,
			RoleName:    role.RoleName,
			RoleCode:    role.RoleCode,
			Description: role.Description,
			CreateBy:    role.CreateBy,
			CreateTime:  role.CreateTime.Unix(),
			UpdateBy:    role.UpdateBy,
			UpdateTime:  role.UpdateTime.Unix(),
		},
	}, nil
}
