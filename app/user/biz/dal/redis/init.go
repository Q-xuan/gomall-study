package redis

import (
	"context"
	"fmt"
	"os"

	"github.com/py/biz-demo/gomall/app/user/conf"
	"github.com/redis/go-redis/v9"
)

var (
	RedisClient *redis.Client
)

func Init() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf(conf.GetConf().Redis.Address, os.Getenv("REDIS_HOST")),
		Username: conf.GetConf().Redis.Username,
		Password: fmt.Sprintf(conf.GetConf().Redis.Password,os.Getenv("REDIS_PASSWORD")),
		DB:       conf.GetConf().Redis.DB,
	})
	if err := RedisClient.Ping(context.Background()).Err(); err != nil {
		panic(err)
	}
}
