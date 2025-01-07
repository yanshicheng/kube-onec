package svc

import (
	"github.com/yanshicheng/kube-onec/application/portal/api/internal/config"
	"github.com/yanshicheng/kube-onec/application/portal/api/internal/middleware"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/client/imageservice"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/client/sysauthservice"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/client/sysdictitemservice"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/client/sysdictservice"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/client/sysorganizationservice"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/client/syspermissionservice"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/client/syspositionservice"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/client/sysroleservice"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/client/sysuserservice"
	"github.com/yanshicheng/kube-onec/common/interceptors"
	"github.com/yanshicheng/kube-onec/common/verify"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config             config.Config
	Cache              *redis.Redis
	SysAuthRpc         sysauthservice.SysAuthService
	SysUserRpc         sysuserservice.SysUserService
	SysRoleRpc         sysroleservice.SysRoleService
	SysPositionRpc     syspositionservice.SysPositionService
	SysOrganizationRpc sysorganizationservice.SysOrganizationService
	SysPermissionRpc   syspermissionservice.SysPermissionService
	SysDictRpc         sysdictservice.SysDictService
	SysDictItemRpc     sysdictitemservice.SysDictItemService
	Validator          *verify.ValidatorInstance
	StoreRpc           imageservice.ImageService
	JWTAuthMiddleware  rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	validator, err := verify.InitValidator(verify.LocaleZH)
	if err != nil {
		panic(err)
	}
	// 自定义拦截器
	protalRpc := zrpc.MustNewClient(c.PortalRPC, zrpc.WithUnaryClientInterceptor(interceptors.ClientErrorInterceptor()))
	return &ServiceContext{
		Config:             c,
		Cache:              redis.MustNewRedis(c.Cache),
		Validator:          validator,
		SysAuthRpc:         sysauthservice.NewSysAuthService(protalRpc),
		SysUserRpc:         sysuserservice.NewSysUserService(protalRpc),
		SysRoleRpc:         sysroleservice.NewSysRoleService(protalRpc),
		SysPositionRpc:     syspositionservice.NewSysPositionService(protalRpc),
		SysOrganizationRpc: sysorganizationservice.NewSysOrganizationService(protalRpc),
		SysPermissionRpc:   syspermissionservice.NewSysPermissionService(protalRpc),
		SysDictRpc:         sysdictservice.NewSysDictService(protalRpc),
		SysDictItemRpc:     sysdictitemservice.NewSysDictItemService(protalRpc),
		StoreRpc:           imageservice.NewImageService(protalRpc),
		JWTAuthMiddleware:  middleware.NewJWTAuthMiddleware(sysauthservice.NewSysAuthService(protalRpc)).Handle,
	}
}
