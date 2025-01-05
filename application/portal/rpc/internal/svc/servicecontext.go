package svc

import (
	"github.com/yanshicheng/kube-onec/application/portal/rpc/internal/config"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/internal/model"
	"github.com/yanshicheng/kube-onec/pkg/storage"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"log"
)

type ServiceContext struct {
	Config            config.Config
	Cache             *redis.Redis
	Storage           storage.Uploader
	SysUser           model.SysUserModel
	SysRole           model.SysRoleModel
	SysOrganization   model.SysOrganizationModel
	SysPosition       model.SysPositionModel
	SysMenu           model.SysMenuModel
	SysPermission     model.SysPermissionModel
	SysDict           model.SysDictModel
	SysDictItem       model.SysDictItemModel
	SysRolePermission model.SysRolePermissionModel
	SysUserRole       model.SysUserRoleModel
	SysRoleMenu       model.SysRoleMenuModel
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
	uploader, err := storage.NewUploader(storage.UploaderOptions{
		AccessKey:    c.StorageConf.AccessKey,
		AccessSecret: c.StorageConf.AccessSecret,
		CAFile:       c.StorageConf.CAFile,
		CAKey:        c.StorageConf.CAKey,
		Endpoints:    c.StorageConf.Endpoints,
		Provider:     c.StorageConf.Provider,
		UseTLS:       c.StorageConf.UseTLS,
		BucketName:   c.StorageConf.BucketName,
	})
	if err != nil {
		log.Fatalf("初始化上传器失败: %v", err)
	}
	return &ServiceContext{
		Config:  c,
		Cache:   redis.MustNewRedis(c.Cache),
		Storage: uploader, // 直接赋值，不需要取地址
		// BookModel: models.NewBooksModel(sqlx.NewMysql(c.Mysql.DataSource), c.DBCache),
		SysUser:           model.NewSysUserModel(sqlConn, c.DBCache),
		SysRole:           model.NewSysRoleModel(sqlConn, c.DBCache),
		SysOrganization:   model.NewSysOrganizationModel(sqlConn, c.DBCache),
		SysPosition:       model.NewSysPositionModel(sqlConn, c.DBCache),
		SysMenu:           model.NewSysMenuModel(sqlConn, c.DBCache),
		SysPermission:     model.NewSysPermissionModel(sqlConn, c.DBCache),
		SysDict:           model.NewSysDictModel(sqlConn, c.DBCache),
		SysDictItem:       model.NewSysDictItemModel(sqlConn, c.DBCache),
		SysRolePermission: model.NewSysRolePermissionModel(sqlConn, c.DBCache),
		SysUserRole:       model.NewSysUserRoleModel(sqlConn, c.DBCache),
		SysRoleMenu:       model.NewSysRoleMenuModel(sqlConn, c.DBCache),
	}
}
