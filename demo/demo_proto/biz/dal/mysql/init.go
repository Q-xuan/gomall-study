package mysql

import (
	"fmt"
	"os"

	"github.com/py/biz-demo/gomall/demo/demo_proto/biz/model"
	"github.com/py/biz-demo/gomall/demo/demo_proto/conf"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Init() {
	dsn := fmt.Sprintf(conf.GetConf().MySQL.DSN,
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_DATEBASE"),
	)
	DB, err = gorm.Open(mysql.Open(dsn),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		panic(err)
	}

	DB.AutoMigrate(&model.User{})
	fmt.Printf("%#v\n", DB.Debug().Exec("select version()"))
}
