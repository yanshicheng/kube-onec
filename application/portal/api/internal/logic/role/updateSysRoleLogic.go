package role

import (
	"context"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/client/sysroleservice"

	"github.com/yanshicheng/kube-onec/application/portal/api/internal/svc"
	"github.com/yanshicheng/kube-onec/application/portal/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateSysRoleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateSysRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateSysRoleLogic {
	return &UpdateSysRoleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateSysRoleLogic) UpdateSysRole(req *types.UpdateSysRoleRequest) (resp string, err error) {
	by, ok := l.ctx.Value("account").(string)
	if !ok || by == "" {
		by = "system"
	}
	_, err = l.svcCtx.SysRoleRpc.UpdateSysRole(l.ctx, &sysroleservice.UpdateSysRoleReq{
		Id:          req.Id,
		RoleName:    req.RoleName,
		Description: req.Description,
		UpdatedBy:    by,
	})
	if err != nil {
		// 错误日志，明确上下文
		logx.WithContext(l.ctx).Errorf("更新角色失败: 请求=%+v, 错误=%v", req, err)
		return
	}
	resp = "更新角色成功"
	logx.WithContext(l.ctx).Infof("更新角色成功: 请求=%+v, 响应=%s", req, resp)
	return
}
