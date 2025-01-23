package dal

import (
	"github.com/py/biz-demo/gomall/app/product/biz/dal/mysql"
)

func Init() {
	// redis.Init()
	mysql.Init()
}
