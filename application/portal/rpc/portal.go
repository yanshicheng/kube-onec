package main

import (
	"flag"
	"fmt"

	"github.com/yanshicheng/kube-onec/application/portal/rpc/internal/config"
	imageserviceServer "github.com/yanshicheng/kube-onec/application/portal/rpc/internal/server/imageservice"
	sysauthserviceServer "github.com/yanshicheng/kube-onec/application/portal/rpc/internal/server/sysauthservice"
	sysdictitemserviceServer "github.com/yanshicheng/kube-onec/application/portal/rpc/internal/server/sysdictitemservice"
	sysdictserviceServer "github.com/yanshicheng/kube-onec/application/portal/rpc/internal/server/sysdictservice"
	sysmenuserviceServer "github.com/yanshicheng/kube-onec/application/portal/rpc/internal/server/sysmenuservice"
	sysorganizationserviceServer "github.com/yanshicheng/kube-onec/application/portal/rpc/internal/server/sysorganizationservice"
	syspermissionserviceServer "github.com/yanshicheng/kube-onec/application/portal/rpc/internal/server/syspermissionservice"
	syspositionserviceServer "github.com/yanshicheng/kube-onec/application/portal/rpc/internal/server/syspositionservice"
	sysroleserviceServer "github.com/yanshicheng/kube-onec/application/portal/rpc/internal/server/sysroleservice"
	sysuserserviceServer "github.com/yanshicheng/kube-onec/application/portal/rpc/internal/server/sysuserservice"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/internal/svc"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/pb"

	"github.com/yanshicheng/kube-onec/common/interceptors"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/portal.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		pb.RegisterSysUserServiceServer(grpcServer, sysuserserviceServer.NewSysUserServiceServer(ctx))
		pb.RegisterSysMenuServiceServer(grpcServer, sysmenuserviceServer.NewSysMenuServiceServer(ctx))
		pb.RegisterSysOrganizationServiceServer(grpcServer, sysorganizationserviceServer.NewSysOrganizationServiceServer(ctx))
		pb.RegisterSysPermissionServiceServer(grpcServer, syspermissionserviceServer.NewSysPermissionServiceServer(ctx))
		pb.RegisterSysPositionServiceServer(grpcServer, syspositionserviceServer.NewSysPositionServiceServer(ctx))
		pb.RegisterSysRoleServiceServer(grpcServer, sysroleserviceServer.NewSysRoleServiceServer(ctx))
		pb.RegisterSysAuthServiceServer(grpcServer, sysauthserviceServer.NewSysAuthServiceServer(ctx))
		pb.RegisterSysDictServiceServer(grpcServer, sysdictserviceServer.NewSysDictServiceServer(ctx))
		pb.RegisterSysDictItemServiceServer(grpcServer, sysdictitemserviceServer.NewSysDictItemServiceServer(ctx))
		pb.RegisterImageServiceServer(grpcServer, imageserviceServer.NewImageServiceServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	// 自定义拦截器
	s.AddUnaryInterceptors(interceptors.ServerErrorInterceptor())

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
