package middlware

import (
	"fmt"
	"os"

	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/hertz-contrib/sessions"
	"github.com/hertz-contrib/sessions/redis"
	"github.com/py/biz-demo/gomall/app/frontend/conf"
)


func InitRedisSession(h * server.Hertz){
	address := fmt.Sprintf(conf.GetConf().Redis.Address, os.Getenv("REDIS_HOST"))
	password := fmt.Sprintf(conf.GetConf().Redis.Password, os.Getenv("REDIS_PASSWORD"))
	store, _ := redis.NewStore(10, "tcp", address, password, []byte(os.Getenv("SESSION_SECRET")))
	h.Use(sessions.New("py-shop", store))
}