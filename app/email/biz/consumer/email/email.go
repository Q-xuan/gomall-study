package email

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/server"
	"github.com/golang/protobuf/proto"
	"github.com/nats-io/nats.go"
	"github.com/py/biz-demo/gomall/app/email/infra/mq"
	"github.com/py/biz-demo/gomall/app/email/infra/notify"
	"github.com/py/biz-demo/gomall/rpc_gen/kitex_gen/email"
)

func ConsumerInit() {

	sub, err := mq.Nc.Subscribe("email", func(msg *nats.Msg) {
		var req email.EmailReq
		err := proto.Unmarshal(msg.Data, &req)
		if err != nil {
			klog.Error(err)
			return
		}

		noopEmail := notify.NewNoopEmail()
		_ = noopEmail.Send(&req)
	})

	if err != nil {
		panic(err)
	}

	server.RegisterShutdownHook(func() {
		sub.Unsubscribe()
		mq.Nc.Close()
	})

}
