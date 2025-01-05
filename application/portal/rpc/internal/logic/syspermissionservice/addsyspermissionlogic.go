package syspermissionservicelogic

import (
	"context"
	"errors"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/internal/code"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/internal/model"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/internal/svc"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddSysPermissionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddSysPermissionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddSysPermissionLogic {
	return &AddSysPermissionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// -----------------------权限表-----------------------
func (l *AddSysPermissionLogic) AddSysPermission(in *pb.AddSysPermissionReq) (*pb.AddSysPermissionResp, error) {
	// 要求 action 只能是 GET POST PUT DELETE *
	if in.Action != "GET" && in.Action != "POST" && in.Action != "PUT" && in.Action != "DELETE" && in.Action != "*" && in.Action != "" {
		l.Logger.Errorf("action 只能是 GET POST PUT DELETE *")
		return nil, code.ActionIllegal
	}
	// 如果 ParentId 不是0 则查询父级是否存在
	level := uint64(1)
	if in.ParentId != 0 {
		perm, err := l.svcCtx.SysPermission.FindOne(l.ctx, in.ParentId)
		if err != nil {
			if errors.Is(err, model.ErrNotFound) {
				l.Logger.Errorf("父级权限不存在: %v", err)
				return nil, code.ParentPermissionNotExist
			}
			l.Logger.Errorf("父级权限查询异常: %v", err)
			return nil, code.ParentPermissionNotExist
		}
		level = perm.Level + 1
		if level > 3 {
			l.Logger.Errorf("组织层级不能超过3级: 父节点ID=%d, 父节点层级=%d", in.ParentId, perm.Level)
			return nil, code.CreateOrganizationErr
		}
	}
	_, err := l.svcCtx.SysPermission.Insert(l.ctx, &model.SysPermission{
		Action:   in.Action,
		Level:    level,
		Name:     in.Name,
		ParentId: in.ParentId,
		Uri:      in.Uri,
	})
	if err != nil {
		l.Logger.Errorf("添加权限失败: %v", err)
		return nil, err
	}
	return &pb.AddSysPermissionResp{}, nil
}
