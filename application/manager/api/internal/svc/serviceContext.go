package svc

import (
	"github.com/yanshicheng/kube-onec/application/manager/api/internal/config"
	"github.com/yanshicheng/kube-onec/application/manager/api/internal/middleware"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/client/onecclusterconninfoservice"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/client/onecclusterservice"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/client/onecnodeservice"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/client/sysauthservice"
	"github.com/yanshicheng/kube-onec/common/interceptors"
	"github.com/yanshicheng/kube-onec/common/verify"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config             config.Config
	Cache              *redis.Redis
	Validator          *verify.ValidatorInstance
	JWTAuthMiddleware  rest.Middleware
	ClusterRpc         onecclusterservice.OnecClusterService
	ClusterConnInfoRpc onecclusterconninfoservice.OnecClusterConnInfoService
	NodeRpc            onecnodeservice.OnecNodeService
}

func NewServiceContext(c config.Config) *ServiceContext {
	validator, err := verify.InitValidator(verify.LocaleZH)
	if err != nil {
		panic(err)
	}
	portalRpc := zrpc.MustNewClient(c.PortalRPC, zrpc.WithUnaryClientInterceptor(interceptors.ClientErrorInterceptor()))
	managerRpc := zrpc.MustNewClient(c.ManagerRPC, zrpc.WithUnaryClientInterceptor(interceptors.ClientErrorInterceptor()))
	return &ServiceContext{
		Config:             c,
		Cache:              redis.MustNewRedis(c.Cache),
		Validator:          validator,
		JWTAuthMiddleware:  middleware.NewJWTAuthMiddleware(sysauthservice.NewSysAuthService(portalRpc)).Handle,
		ClusterRpc:         onecclusterservice.NewOnecClusterService(managerRpc),
		ClusterConnInfoRpc: onecclusterconninfoservice.NewOnecClusterConnInfoService(managerRpc),
		NodeRpc:            onecnodeservice.NewOnecNodeService(managerRpc),
	}
}
