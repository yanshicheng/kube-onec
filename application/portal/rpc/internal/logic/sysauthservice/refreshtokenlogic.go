package sysauthservicelogic

import (
	"context"
	"fmt"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/internal/code"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/internal/svc"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/pb"
	"github.com/yanshicheng/kube-onec/pkg/jwt"
	"github.com/zeromicro/go-zero/core/logx"
)

type RefreshTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRefreshTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RefreshTokenLogic {
	return &RefreshTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 刷新令牌
func (l *RefreshTokenLogic) RefreshToken(in *pb.RefreshTokenRequest) (*pb.RefreshTokenResponse, error) {
	// 拿到刷新令牌
	if in.RefreshToken == "" {
		return nil, code.RefreshTokenEmptyErr
	}
	jwtR, err := jwt.VerifyToken("Bearer "+in.RefreshToken, l.svcCtx.Config.AuthConfig.RefreshSecret)
	if err != nil {
		return nil, err
	}
	// 查询用户信息
	redisKey := fmt.Sprintf("%s%s", UuidKeyPrefix, jwtR.Account.Account)
	if allowMultiLogin {
		redisKey = fmt.Sprintf("%s%s:%s", UuidKeyPrefix, jwtR.Account.Account, jwtR.Account.Uuid)
	}
	key, errs := l.svcCtx.Cache.Get(redisKey)
	if errs != nil {
		l.Logger.Errorf("Redis 查询该 uuid 的访问令牌, 错误: %v, uuid: %v", err, redisKey)
		return nil, code.UUIDQueryErr
	}
	if key == "" {
		return nil, code.UUIDNotExistErr
	}
	jwtA, err := jwt.VerifyToken("Bearer "+in.RefreshToken, l.svcCtx.Config.AuthConfig.RefreshSecret)
	if err != nil {
		return nil, err
	}
	accessToken, errs := jwt.CreateJWTToken(&jwt.AccountInfo{
		AccountId: jwtA.Account.AccountId,
		Account:   jwtA.Account.Account,
		UserName:  jwtA.Account.UserName,
		Uuid:      jwtA.Account.Uuid,
		Roles:     jwtA.Account.Roles,
	}, l.svcCtx.Config.AuthConfig.AccessSecret, jwtA.Account.Uuid, l.svcCtx.Config.AuthConfig.AccessExpire)
	if errs != nil {
		l.Logger.Errorf("生成JWT令牌失败, 错误: %v", err)
		return nil, code.GenerateJWTTokenErr
	}
	errs = l.svcCtx.Cache.Setex(redisKey, accessToken.AccessToken, int(l.svcCtx.Config.AuthConfig.AccessExpire))
	if errs != nil {
		l.Logger.Errorf("存储访问令牌到 Redis 失败: %v", err)
		return nil, code.RedisStorageErr
	}
	return &pb.RefreshTokenResponse{
		AccessToken:     accessToken.AccessToken,
		AccessExpiresIn: accessToken.ExpiresAt,
	}, nil
}
