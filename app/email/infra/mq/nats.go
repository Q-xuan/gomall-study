package mq

import (
	"os"

	"github.com/nats-io/nats.go"
	"github.com/py/biz-demo/gomall/app/email/conf"
)

var (
	Nc  *nats.Conn
	err error
)

func Init() {
	if os.Getenv("GO_ENV") == "test"{
		if Nc, err = nats.Connect(nats.DefaultURL); err != nil {
			return
		}
	}else{
		if Nc, err = nats.Connect(conf.GetConf().Nats.Address); err != nil {
			return
		}
	}
}
