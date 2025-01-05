package user

import (
	"context"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/pb"

	"github.com/yanshicheng/kube-onec/application/portal/api/internal/svc"
	"github.com/yanshicheng/kube-onec/application/portal/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type BindRoleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBindRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BindRoleLogic {
	return &BindRoleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BindRoleLogic) BindRole(req *types.BindRoleRequest) (resp string, err error) {
	l.Logger.Infof("UpdateSysUser: %v", req.RoleIds)

	_, err = l.svcCtx.SysUserRpc.BindRole(l.ctx, &pb.BindRoleReq{
		Id:      req.Id,
		RoleIds: req.RoleIds,
	})
	if err != nil {
		l.Logger.Errorf("绑定角色失败: %v", err)
		return "", err
	}
	return
}
