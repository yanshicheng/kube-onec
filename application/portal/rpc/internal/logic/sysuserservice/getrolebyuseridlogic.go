package sysuserservicelogic

import (
	"context"
	"errors"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/internal/model"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/internal/svc"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/pb"
	"github.com/yanshicheng/kube-onec/common/handler/errorx"
	"github.com/yanshicheng/kube-onec/pkg/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRoleByUserIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetRoleByUserIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRoleByUserIdLogic {
	return &GetRoleByUserIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetRoleByUserIdLogic) GetRoleByUserId(in *pb.GetRoleByUserIdReq) (*pb.GetRoleByUserIdResp, error) {
	sqlStr := "user_id = ?"
	roleBindall, err := l.svcCtx.SysUserRole.SearchNoPage(l.ctx, "id", false, sqlStr, in.Id)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			l.Logger.Infof("用户未绑定角色，用户ID=%s", in.Id)
			return &pb.GetRoleByUserIdResp{}, nil
		}
		// 查询过程中的其他错误
		l.Logger.Errorf("查询角色信息失败，用户ID=%s，错误信息：%v", in.Id, err)
		return nil, errorx.DatabaseQueryErr
	}
	roleIds := make([]uint64, len(roleBindall))
	for i, roleBind := range roleBindall {
		roleIds[i] = roleBind.RoleId
	}
	// 查询出所有的角色信息
	roleAll, err := l.svcCtx.SysRole.SearchNoPage(l.ctx, "id", false, utils.BuildInCondition("id", len(roleIds)), utils.ConvertUint64SliceToInterfaceSlice(roleIds)...)

	if err != nil {
		// 查询过程中的其他错误
		l.Logger.Errorf("查询角色信息失败，用户ID=%s，错误信息：%v", in.Id, err)
		return nil, errorx.DatabaseQueryErr
	}
	roleNames := make([]string, len(roleAll))
	for i, role := range roleAll {
		roleNames[i] = role.RoleCode
	}
	return &pb.GetRoleByUserIdResp{
		RoleIds:   roleIds,
		RoleNames: roleNames,
	}, nil
}
