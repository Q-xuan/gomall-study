package dal

import (
	"github.com/py/biz-demo/gomall/app/checkout/biz/dal/mysql"
	"github.com/py/biz-demo/gomall/app/checkout/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
