package sysuserservicelogic

import (
	"context"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/internal/code"
	"github.com/yanshicheng/kube-onec/utils"

	"github.com/yanshicheng/kube-onec/application/portal/rpc/internal/svc"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChangePasswordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewChangePasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChangePasswordLogic {
	return &ChangePasswordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ChangePasswordLogic) ChangePassword(in *pb.ChangePasswordReq) (*pb.ChangePasswordResp, error) {
	// 修改密码
	account, err := l.svcCtx.SysUser.FindOne(l.ctx, in.Id)
	if err != nil {
		l.Logger.Errorf("查询账号失败: %v", err)
		return nil, code.FindAccountErr
	}
	// 判断输出的旧密码是否正确
	oldPwd, err := utils.DecodeBase64Password(in.OldPassword)
	if err != nil {
		l.Logger.Errorf("解码密码失败: %v", err)
		return nil, code.DecodeBase64PasswordErr
	}
	// 验证旧密码是否正确
	if !utils.CheckPasswordHash(oldPwd, account.Password) {
		l.Logger.Errorf("旧密码错误: %v", err)
		return nil, code.PasswordNotMatchErr
	}
	newPwd, err := utils.DecodeBase64Password(in.NewPassword)
	if err != nil {
		l.Logger.Errorf("解码密码失败: %v", err)
		return nil, code.DecodeBase64PasswordErr
	}
	confirmNewPassword, err := utils.DecodeBase64Password(in.ConfirmPassword)
	if err != nil {
		l.Logger.Errorf("解码密码失败: %v", err)
		return nil, code.DecodeBase64PasswordErr
	}
	if newPwd != confirmNewPassword {
		l.Logger.Errorf("密码不匹配: %v, %v", newPwd, confirmNewPassword)
		return nil, code.PasswordNotMatchErr
	}
	if oldPwd == newPwd {
		l.Logger.Errorf("新旧密码一致: %v", newPwd)
		return nil, code.NewPasswordNotMatchErr
	}
	if !checkPassword(newPwd) {
		l.Logger.Errorf("密码不符合规则: %v", newPwd)
		return nil, code.PasswordIllegal
	}
	// 密码加密修改密码
	newPwdHash, err := utils.EncryptPassword(newPwd)
	if err != nil {
		l.Logger.Errorf("密码加密失败: %v", err)
		return nil, code.EncryptPasswordErr
	}

	account.Password = newPwdHash
	account.IsResetPassword = 0
	l.Logger.Infof("修改密码: %v", account.Account)
	// 保存
	err = l.svcCtx.SysUser.Update(l.ctx, account)
	if err != nil {
		l.Logger.Errorf("修改密码失败: %v", err)
		return nil, code.ChangePasswordErr
	}

	return &pb.ChangePasswordResp{}, nil
}

// 校验函数，必须大于六位，必须包含大小写，数字，特殊字符
// checkPassword 校验密码是否符合规则：必须大于六位，包含大小写字母、数字、特殊字符
func checkPassword(password string) bool {
	// 密码长度必须大于 6 位
	if len(password) <= 6 {
		return false
	}

	// 初始化标志位，用于校验是否包含大写、小写、数字和特殊字符
	var hasUpper, hasLower, hasNumber, hasSpecial bool

	// 遍历密码中的每个字符，进行分类判断
	for _, char := range password {
		// 判断是否是大写字母
		if char >= 'A' && char <= 'Z' {
			hasUpper = true
		} else if char >= 'a' && char <= 'z' { // 判断是否是小写字母
			hasLower = true
		} else if char >= '0' && char <= '9' { // 判断是否是数字
			hasNumber = true
		} else { // 其余字符认为是特殊字符
			hasSpecial = true
		}
	}

	// 只有当所有条件都满足时，返回 true，否则返回 false
	return hasUpper && hasLower && hasNumber && hasSpecial
}
