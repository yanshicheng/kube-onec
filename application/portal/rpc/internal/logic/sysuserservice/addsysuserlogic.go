package sysuserservicelogic

import (
	"context"
	"errors"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/internal/code"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/internal/model"
	utils2 "github.com/yanshicheng/kube-onec/pkg/utils"
	"time"

	"github.com/yanshicheng/kube-onec/application/portal/rpc/internal/svc"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddSysUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddSysUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddSysUserLogic {
	return &AddSysUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// -----------------------账号信息表-----------------------
func (l *AddSysUserLogic) AddSysUser(in *pb.AddSysUserReq) (*pb.AddSysUserResp, error) {
	// 检查关键字段是否传递 UserName 或者 Account
	if in.UserName == "" || in.Account == "" || in.Mobile == "" || in.Email == "" || in.WorkNumber == "" || in.HireDate == 0 {
		return nil, code.AccountRequiredParams
	}
	//string userName = 1; //用户姓名
	//string account = 2; //用户账号，唯一标识
	//string icon = 3; //用户头像URL
	//string mobile = 4; //用户手机号
	//string email = 5; //用户邮箱地址
	//string workNumber = 6; //用户工号
	//int64 hireDate = 7; //入职时间（时间戳）
	// 获取当前时间戳
	// 获取加密密码
	// 默认 Icon
	// 如果机构存在，检查机构是否存在
	if in.OrganizationId != 0 {
		organization, err := l.svcCtx.SysOrganization.FindOne(l.ctx, in.OrganizationId)
		if err != nil {
			if !errors.Is(err, model.ErrNotFound) {
				l.Logger.Errorf("查询机构信息失败: %v", err)
				return nil, code.GetOrganizationInfoErr
			}
			l.Logger.Errorf("查询机构失败: %v", err)
			return nil, code.GetOrganizationErr
		}
		if organization == nil {
			l.Logger.Errorf("机构不存在: %v", in.OrganizationId)
			return nil, code.GetOrganizationInfoErr
		}
	}

	if in.PositionId != 0 {
		position, err := l.svcCtx.SysPosition.FindOne(l.ctx, in.PositionId)
		if err != nil {
			if !errors.Is(err, model.ErrNotFound) {
				l.Logger.Errorf("查询职位信息失败: %v", err)
				return nil, code.GetPositionInfoErr
			}
			l.Logger.Errorf("查询职位失败: %v", err)
			return nil, code.GetPositionErr
		}
		if position == nil {
			l.Logger.Errorf("职位不存在: %v", in.PositionId)
			return nil, code.GetPositionInfoErr
		}
	}
	icon := "/users/20241223/20241223175322.jpg"
	encryptPassword, err := utils2.EncryptPassword(utils2.GeneratePassword())
	if err != nil {
		l.Logger.Errorf("密码加密失败: %v", err)
		return nil, code.EncryptPasswordErr
	}
	_, err = l.svcCtx.SysUser.Insert(l.ctx, &model.SysUser{
		Account:         in.Account,
		Email:           in.Email,
		HireDate:        utils2.FormattedDate(in.HireDate),
		Icon:            icon,
		Mobile:          in.Mobile,
		UserName:        in.UserName,
		WorkNumber:      in.WorkNumber,
		Password:        encryptPassword,
		IsResetPassword: 1,
		IsDisabled:      0,
		IsLeave:         0,
		OrganizationId:  in.OrganizationId,
		PositionId:      in.PositionId,
		LastLoginTime:   time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
	})
	if err != nil {
		l.Logger.Errorf("创建账号失败: %v", err)
		return nil, code.CreateAccountErr
	}
	l.Logger.Infof("创建账号成功: %v", in.Account)
	return &pb.AddSysUserResp{}, nil
}
