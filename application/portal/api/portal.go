package main

import (
	"flag"
	"fmt"
	"github.com/yanshicheng/kube-onec/application/portal/api/internal/config"
	"github.com/yanshicheng/kube-onec/application/portal/api/internal/handler"
	"github.com/yanshicheng/kube-onec/application/portal/api/internal/svc"
	"github.com/yanshicheng/kube-onec/common/handler/errorx"
	"github.com/yanshicheng/kube-onec/common/handler/okx"
	middlewarex "github.com/yanshicheng/kube-onec/common/middleware"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/portal-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf, rest.WithUnauthorizedCallback(JwtUnauthorizedResult))
	defer server.Stop()

	// 自定义全局中间件
	server.Use(middlewarex.PanicRecoveryMiddleware)

	// 自定义错误
	httpx.SetErrorHandler(errorx.ErrHandler)
	httpx.SetOkHandler(okx.OkHandler)

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}

func JwtUnauthorizedResult(w http.ResponseWriter, r *http.Request, err error) {
	// 根据不同的 jwt 错误返回不同的中文提示
	logx.Infof("jwt 错误：%v", err)
}
