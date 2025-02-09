package mq

import "github.com/nats-io/nats.go"

var (
	Nc  *nats.Conn
	err error
)

func Init() {
	if Nc, err = nats.Connect(nats.DefaultURL); err != nil {
		return
	}

}
