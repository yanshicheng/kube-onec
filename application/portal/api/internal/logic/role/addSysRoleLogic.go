package role

import (
	"context"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/client/sysroleservice"

	"github.com/yanshicheng/kube-onec/application/portal/api/internal/svc"
	"github.com/yanshicheng/kube-onec/application/portal/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddSysRoleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddSysRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddSysRoleLogic {
	return &AddSysRoleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddSysRoleLogic) AddSysRole(req *types.AddSysRoleRequest) (resp string, err error) {
	// 获取上下文中的 account 值，并检查其类型是否为字符串
	account, ok := l.ctx.Value("account").(string)
	if !ok || account == "" {
		// 如果不存在或为空字符串，则默认使用 "system"
		account = "system"
	}
	res, err := l.svcCtx.SysRoleRpc.AddSysRole(l.ctx, &sysroleservice.AddSysRoleReq{
		RoleName:    req.RoleName,
		Description: req.Description,
		RoleCode:    req.RoleCode,
		CreatedBy:    account,
		UpdatedBy:    account,
	})
	if err != nil {
		logx.WithContext(l.ctx).Errorf("添加角色失败: %v", err)
		return "", err
	}
	logx.WithContext(l.ctx).Infof("添加角色成功: %v", res)
	return "添加用户成功", nil
}
