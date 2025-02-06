package dal

import (
	"github.com/py/biz-demo/gomall/app/payment/biz/dal/mysql"
	// "github.com/py/biz-demo/gomall/app/payment/biz/dal/redis"
)

func Init() {
	// redis.Init()
	mysql.Init()
}
