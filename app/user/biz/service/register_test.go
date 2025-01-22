package service

import (
	"context"
	"testing"

	user "github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/user"
	"github.com/joho/godotenv"
	"github.com/py/biz-demo/gomall/app/user/biz/dal/mysql"
)

func TestRegister_Run(t *testing.T) {
	godotenv.Load("../../.env")
	mysql.Init()
	ctx := context.Background()
	s := NewRegisterService(ctx)
	// init req and assert value

	req := &user.RegisterReq{
		Email:    "1admin@123.com",
		Password: "123456",
		ConfirmPassword: "123456",
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
