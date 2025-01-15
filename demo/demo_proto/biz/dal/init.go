package dal

import (
	"github.com/py/biz-demo/gomall/demo/demo_proto/biz/dal/mysql"
	"github.com/py/biz-demo/gomall/demo/demo_proto/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
