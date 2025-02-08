package dal

import (
	// "github.com/py/biz-demo/gomall/app/frontend/biz/dal/mysql"
	"github.com/py/biz-demo/gomall/app/frontend/biz/dal/redis"
)

func Init() {
	redis.Init()
	// mysql.Init()
}
