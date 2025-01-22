package dal

import (
	"github.com/py/biz-demo/gomall/app/user/biz/dal/mysql"
	"github.com/py/biz-demo/gomall/app/user/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
