package dal

import (
	"github.com/py/biz-demo/gomall/app/email/biz/dal/mysql"
	"github.com/py/biz-demo/gomall/app/email/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
