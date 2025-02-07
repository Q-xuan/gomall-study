package rpc

import (
	"github.com/py/biz-demo/gomall/rpc_gen/kitex_gen/order/orderservice"
	"sync"

	"github.com/cloudwego/kitex/client"
	consul "github.com/kitex-contrib/registry-consul"
	"github.com/py/biz-demo/gomall/app/checkout/conf"
	"github.com/py/biz-demo/gomall/rpc_gen/kitex_gen/cart/cartservice"
	"github.com/py/biz-demo/gomall/rpc_gen/kitex_gen/payment/paymentservice"
	"github.com/py/biz-demo/gomall/rpc_gen/kitex_gen/product/productcatalogservice"
)

var (
	CardClient    cartservice.Client
	ProductClient productcatalogservice.Client
	PaymentClient paymentservice.Client
	OrderClient   orderservice.Client
	once          sync.Once
	err           error
)

func InitClient() {
	once.Do(func() {
		initCartClient()
		initProductClient()
		initPaymentClient()
		initOrderClient()
	})
}

func initOrderClient() {
	r, err := consul.NewConsulResolver(conf.GetConf().Registry.RegistryAddress[0])
	if err != nil {
		panic(err)
	}
	OrderClient, err = orderservice.NewClient("order", client.WithResolver(r))
	if err != nil {
		panic(err)
	}
}

func initCartClient() {
	r, err := consul.NewConsulResolver(conf.GetConf().Registry.RegistryAddress[0])
	if err != nil {
		panic(err)
	}
	CardClient, err = cartservice.NewClient("cart", client.WithResolver(r))
	if err != nil {
		panic(err)
	}
}

func initProductClient() {
	r, err := consul.NewConsulResolver(conf.GetConf().Registry.RegistryAddress[0])
	if err != nil {
		panic(err)
	}
	ProductClient, err = productcatalogservice.NewClient("product", client.WithResolver(r))
	if err != nil {
		panic(err)
	}
}

func initPaymentClient() {
	r, err := consul.NewConsulResolver(conf.GetConf().Registry.RegistryAddress[0])
	if err != nil {
		panic(err)
	}
	PaymentClient, err = paymentservice.NewClient("payment", client.WithResolver(r))
	if err != nil {
		panic(err)
	}
}
