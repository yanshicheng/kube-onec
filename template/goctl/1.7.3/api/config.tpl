package config

import (
    {{.authImport}}
    "github.com/zeromicro/go-zero/core/stores/redis"
)

type Config struct {
	rest.RestConf
	Cache   redis.RedisConf
	{{.auth}}
	{{.jwtTrans}}
}
