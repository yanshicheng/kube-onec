package sysroleservicelogic

import (
	"context"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/internal/svc"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/pb"
	"github.com/yanshicheng/kube-onec/common/handler/errorx"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelSysRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelSysRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelSysRoleLogic {
	return &DelSysRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelSysRoleLogic) DelSysRole(in *pb.DelSysRoleReq) (*pb.DelSysRoleResp, error) {
	// 通过 token 获取用户名
	userName := l.ctx.Value("account")
	if userName == "" {
		userName = "admin"
	}
	err := l.svcCtx.SysRole.DeleteSoft(l.ctx, in.Id) // 使用软删除
	l.Logger.Errorf("删除角色: %v   %v", in.Id, err)
	if err != nil {
		l.Logger.Errorf("删除角色失败: %v", err)
		return nil, errorx.DatabaseQueryErr
	}
	return &pb.DelSysRoleResp{}, nil
}
