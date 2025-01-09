package svc

import (
	{{.configImport}}
    "github.com/yanshicheng/kube-onec/common/verify"
    "github.com/zeromicro/go-zero/core/stores/redis"
)

type ServiceContext struct {
	Config {{.config}}
    Cache   *redis.Redis
    Validator  *verify.ValidatorInstance
	{{.middleware}}
}

func NewServiceContext(c {{.config}}) *ServiceContext {
	validator, err := verify.InitValidator(verify.LocaleZH)
	if err != nil {
		panic(err)
	}
	return &ServiceContext{
		Config: c,
        Cache:   redis.MustNewRedis(c.Cache),
        Validator:  validator,
		{{.middlewareAssignment}}
	}
}
