package sysuserservicelogic

import (
	"context"
	"errors"
	"fmt"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/internal/model"
	"github.com/yanshicheng/kube-onec/common/handler/errorx"

	"github.com/yanshicheng/kube-onec/application/portal/rpc/internal/svc"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserInfoLogic) GetUserInfo(in *pb.GetUserInfoReq) (*pb.GetUserInfoResp, error) {
	user, err := l.svcCtx.SysUser.FindOneByAccount(l.ctx, in.Account)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			l.Logger.Errorf("查询用户信息失败，用户ID=%s，错误信息：%v", in.Account, err)
			return nil, errorx.DatabaseNotFound
		}
		// 查询过程中的其他错误
		l.Logger.Errorf("查询用户信息失败，用户ID=%s，错误信息：%v", in.Account, err)
		return nil, errorx.DatabaseQueryErr
	}

	// 获取 icon 绝对地址
	httpPattern := "http://"
	if l.svcCtx.Config.StorageConf.UseTLS {
		httpPattern = "https://"
	}
	absIcon := fmt.Sprintf("%s%s/%s%s", httpPattern, l.svcCtx.Config.StorageConf.Endpoints[0], l.svcCtx.Config.StorageConf.BucketName, user.Icon)

	//// 获取职位name
	position, err := l.svcCtx.SysPosition.FindOne(l.ctx, user.PositionId)
	positionName := position.Name
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			positionName = ""
		}
		// 查询过程中的其他错误
		l.Logger.Errorf("查询职位信息失败，职位ID=%d，错误信息：%v", user.PositionId, err)
		return nil, errorx.DatabaseQueryErr
	}
	// 获取组织机构name
	organizationFullName, err := l.getOrganizationPath(user.OrganizationId)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			organizationFullName = ""
		}
		// 查询过程中的其他错误
		l.Logger.Errorf("查询组织信息失败，组织ID=%d，错误信息：%v", user.OrganizationId, err)
		return nil, errorx.DatabaseQueryErr
	}
	// 查询所有角色信息
	sqlStr := "user_id = ?"
	roles, err := l.svcCtx.SysUserRole.SearchNoPage(l.ctx, "id", false, sqlStr, user.Id)
	var rolesNames []string
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			rolesNames = append(rolesNames, "user")
		}
		// 查询过程中的其他错误
		l.Logger.Errorf("查询角色信息失败，用户ID=%s，错误信息：%v", in.Account, err)
		return nil, errorx.DatabaseQueryErr
	}
	for _, userRole := range roles {
		role, err := l.svcCtx.SysRole.FindOne(l.ctx, userRole.RoleId)
		if err != nil {
			continue
		}
		rolesNames = append(rolesNames, role.RoleCode)
	}
	return &pb.GetUserInfoResp{
		Id:               user.Id,
		UserName:         user.UserName,
		Account:          user.Account,
		Icon:             absIcon,
		Mobile:           user.Mobile,
		Email:            user.Email,
		WorkNumber:       user.WorkNumber,
		HireDate:         user.HireDate.Unix(),
		PositionName:     positionName,
		OrganizationName: organizationFullName,
		RoleNames:        rolesNames,
		LastLoginTime:    user.LastLoginTime.Unix(),
		CreatedAt:        user.CreatedAt.Unix(),
		UpdatedAt:        user.UpdatedAt.Unix(),
	}, nil
}

// 获取组织层级路径的递归函数
func (l *GetUserInfoLogic) getOrganizationPath(id uint64) (string, error) {
	// 查询当前组织
	organization, err := l.svcCtx.SysOrganization.FindOne(l.ctx, id)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return "", errorx.DatabaseNotFound
		}
		// 查询过程中的其他错误
		l.Logger.Errorf("查询组织信息失败，组织ID=%d，错误信息：%v", id, err)
		return "", errorx.DatabaseQueryErr
	}
	// 如果已经是顶层组织（ParentId == 0），返回当前组织名称
	if organization.ParentId == 0 {
		return organization.Name, nil
	}

	// 否则递归查询父级组织路径，并拼接路径
	parentPath, err := l.getOrganizationPath(organization.ParentId)
	if err != nil {
		return "", err
	}

	// 拼接父级路径和当前组织名称
	return fmt.Sprintf("%s/%s", parentPath, organization.Name), nil
}
