package main

import (
	"flag"
	"fmt"

	"github.com/yanshicheng/kube-onec/application/manager/rpc/internal/config"
	onecclusterconninfoserviceServer "github.com/yanshicheng/kube-onec/application/manager/rpc/internal/server/onecclusterconninfoservice"
	onecclusterserviceServer "github.com/yanshicheng/kube-onec/application/manager/rpc/internal/server/onecclusterservice"
	onecnodeserviceServer "github.com/yanshicheng/kube-onec/application/manager/rpc/internal/server/onecnodeservice"
	onecprojectadminserviceServer "github.com/yanshicheng/kube-onec/application/manager/rpc/internal/server/onecprojectadminservice"
	onecprojectapplicationserviceServer "github.com/yanshicheng/kube-onec/application/manager/rpc/internal/server/onecprojectapplicationservice"
	onecprojectquotaserviceServer "github.com/yanshicheng/kube-onec/application/manager/rpc/internal/server/onecprojectquotaservice"
	onecprojectserviceServer "github.com/yanshicheng/kube-onec/application/manager/rpc/internal/server/onecprojectservice"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/internal/svc"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/pb"

	"github.com/yanshicheng/kube-onec/common/interceptors"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/manager.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		pb.RegisterOnecClusterServiceServer(grpcServer, onecclusterserviceServer.NewOnecClusterServiceServer(ctx))
		pb.RegisterOnecClusterConnInfoServiceServer(grpcServer, onecclusterconninfoserviceServer.NewOnecClusterConnInfoServiceServer(ctx))
		pb.RegisterOnecNodeServiceServer(grpcServer, onecnodeserviceServer.NewOnecNodeServiceServer(ctx))
		pb.RegisterOnecProjectServiceServer(grpcServer, onecprojectserviceServer.NewOnecProjectServiceServer(ctx))
		pb.RegisterOnecProjectQuotaServiceServer(grpcServer, onecprojectquotaserviceServer.NewOnecProjectQuotaServiceServer(ctx))
		pb.RegisterOnecProjectAdminServiceServer(grpcServer, onecprojectadminserviceServer.NewOnecProjectAdminServiceServer(ctx))
		pb.RegisterOnecProjectApplicationServiceServer(grpcServer, onecprojectapplicationserviceServer.NewOnecProjectApplicationServiceServer(ctx))

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
