package sysuserservicelogic

import (
	"context"
	"errors"
	"fmt"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/internal/model"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/internal/svc"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/pb"
	"github.com/yanshicheng/kube-onec/common/handler/errorx"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSysUserByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetSysUserByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSysUserByIdLogic {
	return &GetSysUserByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetSysUserByIdLogic) GetSysUserById(in *pb.GetSysUserByIdReq) (*pb.GetSysUserByIdResp, error) {
	// 查询用户信息
	user, err := l.svcCtx.SysUser.FindOne(l.ctx, in.Id)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			l.Logger.Errorf("用户不存在: %v", err)
			return nil, errorx.NotFound
		}
		// 查询过程中的其他错误
		l.Logger.Errorf("查询用户信息失败，用户ID=%d，错误信息：%v", in.Id, err)
		return nil, errorx.DatabaseQueryErr
	}

	// 获取 icon 绝对地址
	httpPattern := "http://"
	if l.svcCtx.Config.StorageConf.UseTLS {
		httpPattern = "https://"
	}
	absIcon := fmt.Sprintf("%s%s/%s/%s", httpPattern, l.svcCtx.Config.StorageConf.Endpoints[0], l.svcCtx.Config.StorageConf.BucketName, user.Icon)

	//// 获取职位name
	//position, err := l.svcCtx.SysPosition.FindOne(l.ctx, user.PositionId)
	//positionName := position.Name
	//if err != nil {
	//	if errors.Is(err, model.ErrNotFound) {
	//		positionName = ""
	//	}
	//	// 查询过程中的其他错误
	//	l.Logger.Errorf("查询职位信息失败，职位ID=%d，错误信息：%v", user.PositionId, err)
	//	return nil, errorx.ServerDataQueryErr
	//}
	//// 获取组织机构name
	//organizationFullName, err := l.getOrganizationPath(user.OrganizationId)
	//if err != nil {
	//	if errors.Is(err, model.ErrNotFound) {
	//		organizationFullName = ""
	//	}
	//	// 查询过程中的其他错误
	//	l.Logger.Errorf("查询组织信息失败，组织ID=%d，错误信息：%v", user.OrganizationId, err)
	//	return nil, errorx.ServerDataQueryErr
	//}

	return &pb.GetSysUserByIdResp{
		Data: &pb.SysUser{
			Id:              user.Id,
			UserName:        user.UserName,
			Account:         user.Account,
			Icon:            absIcon,
			Mobile:          user.Mobile,
			Email:           user.Email,
			WorkNumber:      user.WorkNumber,
			HireDate:        user.HireDate.Unix(),
			IsResetPassword: user.IsResetPassword,
			IsDisabled:      user.IsDisabled,
			IsLeave:         user.IsLeave,
			PositionId:      user.PositionId,
			OrganizationId:  user.OrganizationId,
			LastLoginTime:   user.LastLoginTime.Unix(),
			CreatedAt:       user.CreatedAt.Unix(),
			UpdatedAt:       user.UpdatedAt.Unix(),
		},
	}, nil
}
func int64ToBool(value int64) bool {
	// 如果值为 0，返回 false，否则返回 true
	return value != 0
}
