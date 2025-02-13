package dal

import (
	"github.com/py/biz-demo/gomall/app/product/biz/dal/mysql"
	"github.com/py/biz-demo/gomall/app/product/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
