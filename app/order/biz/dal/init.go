package dal

import (
	"github.com/py/biz-demo/gomall/app/order/biz/dal/mysql"
	// "github.com/py/biz-demo/gomall/app/order/biz/dal/redis"
)

func Init() {
	// redis.Init()
	mysql.Init()
}
