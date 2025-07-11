package start

import (
	"context"
	"github.com/go-redis/redis/v8"
	"weikang/global"
)

func InitRedis() {
	appends := global.NacosConfig.Redis
	global.Rds = redis.NewClient(&redis.Options{
		Addr:     appends.Host,
		Password: appends.Password, // no password set
		DB:       appends.Db,       // use default DB
	})
	err := global.Rds.Ping(context.Background()).Err()
	if err != nil {
		return
	}

}
