package svc

import (
	"github.com/yanshicheng/kube-onec/application/manager/rpc/internal/config"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/internal/model"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/client/sysdictitemservice"
	"github.com/yanshicheng/kube-onec/common/interceptors"
	"github.com/yanshicheng/kube-onec/common/k8swrapper/manager"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
	"log"
)

type ServiceContext struct {
	Config                   config.Config
	Cache                    *redis.Redis
	SysDictItemRpc           sysdictitemservice.SysDictItemService
	OnecClient               *manager.OnecK8sClientManager
	ClusterModel             model.OnecClusterModel
	NodeModel                model.OnecNodeModel
	LabelsResourceModel      model.OnecResourceLabelsModel
	TaintsResourceModel      model.OnecResourceTaintsModel
	AnnotationsResourceModel model.OnecResourceAnnotationsModel
	ProjectModel             model.OnecProjectModel
	ProjectAdminModel        model.OnecProjectAdminModel
	ProjectApplicationModel  model.OnecProjectApplicationModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.Mysql.DataSource)
	rawDB, err := sqlConn.RawDB()
	if err != nil {
		log.Fatal(err) // 处理错误
	}
	// 配置连接池参数
	rawDB.SetMaxOpenConns(c.Mysql.MaxOpenConns)       // 最大打开连接数
	rawDB.SetMaxIdleConns(c.Mysql.MaxIdleConns)       // 最大空闲连接数
	rawDB.SetConnMaxLifetime(c.Mysql.ConnMaxLifetime) // 连接的最大生命周期

	// 自定义拦截器
	protalRpc := zrpc.MustNewClient(c.PortalRPC, zrpc.WithUnaryClientInterceptor(interceptors.ClientErrorInterceptor()))
	return &ServiceContext{
		Config:                   c,
		Cache:                    redis.MustNewRedis(c.Cache),
		OnecClient:               manager.NewOnecK8sClientManager(),
		ClusterModel:             model.NewOnecClusterModel(sqlConn, c.DBCache),
		NodeModel:                model.NewOnecNodeModel(sqlConn, c.DBCache),
		SysDictItemRpc:           sysdictitemservice.NewSysDictItemService(protalRpc),
		LabelsResourceModel:      model.NewOnecResourceLabelsModel(sqlConn, c.DBCache),
		TaintsResourceModel:      model.NewOnecResourceTaintsModel(sqlConn, c.DBCache),
		AnnotationsResourceModel: model.NewOnecResourceAnnotationsModel(sqlConn, c.DBCache),
		ProjectModel:             model.NewOnecProjectModel(sqlConn, c.DBCache),
		ProjectAdminModel:        model.NewOnecProjectAdminModel(sqlConn, c.DBCache),
		ProjectApplicationModel:  model.NewOnecProjectApplicationModel(sqlConn, c.DBCache),
		// BookModel: models.NewBooksModel(sqlx.NewMysql(c.Mysql.DataSource), c.DBCache),
	}
}
