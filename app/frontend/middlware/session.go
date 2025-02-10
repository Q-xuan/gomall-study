package middlware

import (
	"os"

	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/hertz-contrib/sessions"
	"github.com/hertz-contrib/sessions/redis"
	"github.com/py/biz-demo/gomall/app/frontend/conf"
)

func InitRedisSession(h *server.Hertz) {
	address := conf.GetConf().Redis.Address
	store, err := redis.NewStore(100, "tcp", address, "", []byte(os.Getenv("SESSION_SECRET")))
	if err != nil {
		panic(err)
	}
	h.Use(sessions.New("py-shop", store))
}
