package dal

import (
	"github.com/py/biz-demo/gomall/app/user/biz/dal/mysql"
)

func Init() {
	// redis.Init()
	mysql.Init()
}
