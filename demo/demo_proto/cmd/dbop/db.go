package main

import (
	"github.com/joho/godotenv"
	"github.com/py/biz-demo/gomall/demo/demo_proto/biz/dal"
	"github.com/py/biz-demo/gomall/demo/demo_proto/biz/dal/mysql"
	"github.com/py/biz-demo/gomall/demo/demo_proto/biz/model"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	dal.Init()

	//CURD

	//CREATE
	// mysql.DB.Create(&model.User{Email: "demo@example.com",Password: "demo123"})
	// mysql.DB.Model(&model.User{}).Where("email = ?", "demo@example.com").Update("password", "demo345")


	// READ
	// var row model.User
	// mysql.DB.Model(&model.User{}).Where("email = ?", "demo@example.com").First(&row)

	// fmt.Printf("row: %+v\n", row)


	//DELETE
	// mysql.DB.Where("email = ?", "demo@example.com").Delete(&model.User{})

	// mysql.DB.Unscoped().Where("email = ?", "demo@example.com").Delete(&model.User{})

}
