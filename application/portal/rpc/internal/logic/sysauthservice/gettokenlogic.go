package sysauthservicelogic

import (
	"context"
	"fmt"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/internal/code"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/internal/model"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/internal/svc"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/pb"
	"github.com/yanshicheng/kube-onec/common/handler/errorx"
	"github.com/yanshicheng/kube-onec/pkg/jwt"
	utils2 "github.com/yanshicheng/kube-onec/pkg/utils"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

const (
	LoginFailKeyPrefix = "login_fail_" // Redis 键前缀

	MaxLoginFailures    = 3               // 最大登录失败次数
	LoginFailExpire     = 5 * time.Minute // 登录失败记录过期时间
	DefaultRole         = "user"          // 默认角色
	IsResetPasswordFlag = 1               // 需要重置密码标志
	IsDisabledFlag      = 1               // 账号禁用标志
)

// 设置 用户 token redis 前缀
const (
	UuidKeyPrefix   = "account:token:" // Redis 键前缀
	allowMultiLogin = true             // 是否允许多端登录
)

type GetTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTokenLogic {
	return &GetTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取令牌
func (l *GetTokenLogic) GetToken(in *pb.GetTokenRequest) (*pb.GetTokenResponse, error) {
	// 查询账号
	user, err := l.svcCtx.SysUser.FindOneByAccount(l.ctx, in.Account)
	if err != nil {
		l.Logger.Errorf("查询账号失败, 账号: %s, 错误: %v", in.Account, err)
		return nil, code.FindAccountErr
	}
	// 验证是否需要重置密码
	if user.IsResetPassword == 1 {
		l.Logger.Infof("账号需要重置密码, 账号: %s", in.Account)
		return nil, code.ResetPasswordTip
	}
	// 检查账号是否被冻结
	if user.IsDisabled == 1 {
		l.Logger.Infof("账号已被冻结, 账号: %s", in.Account)
		return nil, code.FrozenAccountsErr
	}
	// 查看账号是否离职
	if user.IsLeave == 1 {
		l.Logger.Infof("账号已离职, 禁止登陆。 账号: %s", in.Account)
		return nil, code.AccountLockedTip
	}
	// 处理密码验证逻辑
	if err := l.verifyPassword(in.Password, user.Password); err != nil {
		// 如果密码验证失败，记录登录失败次数
		if err := l.recordLoginFailure(user); err != nil {
			l.Logger.Errorf("记录登录失败次数时发生错误, 账号: %s, 错误: %v", in.Account, err)
			return nil, code.ChangePasswordErr
		}
		return nil, code.LoginErr
	}

	// 查询所有角色信息
	sqlStr := "user_id = ?"
	roles, err := l.svcCtx.SysUserRole.SearchNoPage(l.ctx, "id", false, sqlStr, user.Id)
	var rolesNames []string
	var rolesIds []uint64

	if err != nil {
		// 查询过程中的其他错误
		l.Logger.Errorf("查询角色信息失败，用户=%d，错误信息：%v", user.Id, err)
		return nil, errorx.DatabaseQueryErr
	}

	if len(roles) == 0 {
		rolesNames = append(rolesNames, "user")
	}
	for _, userRole := range roles {
		role, err := l.svcCtx.SysRole.FindOne(l.ctx, userRole.RoleId)
		if err != nil {
			continue
		}
		rolesNames = append(rolesNames, role.RoleCode)
		rolesIds = append(rolesIds, role.Id)
	}
	l.Logger.Infof("账号登录成功, 账号: %s", in.Account)
	uuid, err := utils2.GenerateRandomID()
	if err != nil {
		l.Logger.Errorf("生成UUID失败, 错误: %v", err)
		return nil, code.GenerateUUIDErr
	}
	accessToken, err := jwt.CreateJWTToken(&jwt.AccountInfo{
		AccountId: user.Id,
		Account:   user.Account,
		UserName:  user.UserName,
		Uuid:      uuid,
		Roles:     rolesNames,
	}, l.svcCtx.Config.AuthConfig.AccessSecret, uuid, l.svcCtx.Config.AuthConfig.AccessExpire)
	if err != nil {
		l.Logger.Errorf("生成JWT令牌失败, 错误: %v", err)
		return nil, code.GenerateJWTTokenErr
	}
	// 先判断 redis 是否已经有这个 uuid 访问令牌了
	//_, err = l.svcCtx.Cache.Get(fmt.Sprintf("%s%s", UuidKeyPrefix, user.Account))
	//if err != nil {
	//	l.Logger.Errorf("Redis 查询该 uuid 的访问令牌, 错误: %v, uuid: %v", err, uuid)
	//	return nil, code.UUIDQueryErr
	//}
	//if key != "" {
	//	l.Logger.Infof("Redis 中已存在该 uuid 的访问令牌, 错误: %v , key: %v", key, err)
	//	return nil, code.UUIDExistErr
	//}
	//if err := l.svcCtx.Cache.Setex(fmt.Sprintf("%s%s", UuidKeyPrefix, user.Account), accessToken.AccessToken, int(l.svcCtx.Config.AuthConfig.AccessExpire)); err != nil {
	//	l.Logger.Errorf("存储访问令牌到 Redis 失败: %v", err)
	//	return nil, code.RedisStorageErr
	//}
	// 多端登录模式：允许多端 Token 存储
	redisKey := fmt.Sprintf("%s%s", UuidKeyPrefix, user.Account)
	if allowMultiLogin {
		redisKey = fmt.Sprintf("%s%s:%s", UuidKeyPrefix, user.Account, uuid)
	}
	err = l.svcCtx.Cache.Setex(redisKey, accessToken.AccessToken, int(l.svcCtx.Config.AuthConfig.AccessExpire))
	if err != nil {
		l.Logger.Errorf("存储访问令牌到 Redis 失败: %v", err)
		return nil, code.RedisStorageErr
	}
	refreshToken, err := jwt.CreateJWTToken(&jwt.AccountInfo{
		Account: user.Account,
		Uuid:    uuid,
	}, l.svcCtx.Config.AuthConfig.RefreshSecret, uuid, l.svcCtx.Config.AuthConfig.RefreshExpire)
	if err != nil {
		l.Logger.Errorf("生成JWT令牌失败, 错误: %v", err)
		return nil, code.GenerateJWTTokenErr
	}
	return &pb.GetTokenResponse{
		Account:   user.Account,
		AccountId: user.Id,
		UserName:  user.UserName,
		Roles:     rolesNames,
		Token: &pb.TokenResponse{
			AccessToken:      accessToken.AccessToken,
			RefreshToken:     refreshToken.AccessToken,
			AccessExpiresIn:  accessToken.ExpiresAt,
			RefreshExpiresIn: refreshToken.ExpiresAt,
		},
	}, nil
}

// 验证密码是否正确
func (l *GetTokenLogic) verifyPassword(inputPassword, hashedPassword string) error {
	// 解码密码
	password, err := utils2.DecodeBase64Password(inputPassword)
	if err != nil {
		l.Logger.Errorf("密码解码失败: %v", err)
		return code.DecodeBase64PasswordErr
	}

	// 校验密码
	if !utils2.CheckPasswordHash(password, hashedPassword) {
		l.Logger.Infof("密码验证失败")
		return code.LoginErr
	}

	return nil
}

// 记录登录失败次数，如果 5 分钟内连续失败 3 次则禁用账号
func (l *GetTokenLogic) recordLoginFailure(user *model.SysUser) error {
	// Redis 键，存储账号的登录失败次数
	loginFailKey := fmt.Sprintf("%s%s", LoginFailKeyPrefix, user.Account)

	// 递增登录失败次数，同时检查当前次数
	failCount, err := l.svcCtx.Cache.Incr(loginFailKey)
	if err != nil {
		return fmt.Errorf("redis Incr 操作失败: %w", err)
	}

	// 如果这是第一次失败，设置失败次数的过期时间为 5 分钟
	if failCount == 1 {
		if err := l.svcCtx.Cache.Expire(loginFailKey, int(LoginFailExpire.Seconds())); err != nil {
			return fmt.Errorf("redis 设置过期时间失败: %w", err)
		}
	}

	// 检查是否达到最大失败次数
	if failCount >= MaxLoginFailures {
		// 禁用账号的业务逻辑（可以通过数据库标记账号为禁用）
		if err := l.disableAccount(user); err != nil {
			return fmt.Errorf("禁用账号失败: %w", err)
		}
		// 删除 Redis 中的失败记录，避免重复触发
		if _, err := l.svcCtx.Cache.Del(loginFailKey); err != nil {
			return fmt.Errorf("删除 Redis 登录失败记录失败: %w", err)
		}

		l.Logger.Infof("账号已被禁用, 账号: %s", user.Account)
		return code.AccountLockedErr
	}

	l.Logger.Infof("登录失败计数: %d, 账号: %s", failCount, user.Account)
	return nil
}

// 禁用账号（可以根据具体业务实现，例如更新数据库状态）
func (l *GetTokenLogic) disableAccount(user *model.SysUser) error {
	// 示例实现：将账号设置为禁用状态
	user.IsDisabled = 1
	err := l.svcCtx.SysUser.Update(l.ctx, user)
	if err != nil {
		l.Logger.Errorf("禁用账号失败, 账号: %s, 错误: %v", user.Account, err)
		return err
	}
	return nil
}

// 检查是否有效的刷新令牌
func (l *GetTokenLogic) checkRefreshTokenValid(uuid string) (bool, error) {
	// 获取 Redis 中存储的刷新令牌
	refreshTokenKey := fmt.Sprintf("refresh_token:%s", uuid)
	storedRefreshToken, err := l.svcCtx.Cache.Get(refreshTokenKey)
	if err != nil {
		l.Logger.Errorf("获取刷新令牌失败, 错误: %v", err)
		return false, fmt.Errorf("获取刷新令牌失败")
	}

	if storedRefreshToken == "" {
		l.Logger.Info("刷新令牌已过期或不存在")
		return false, nil
	}

	return true, nil
}
