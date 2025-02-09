package rpc

import (
	"sync"

	"github.com/py/biz-demo/gomall/rpc_gen/kitex_gen/order/orderservice"

	"github.com/cloudwego/kitex/client"
	"github.com/py/biz-demo/gomall/app/checkout/conf"
	checkoututils "github.com/py/biz-demo/gomall/app/checkout/utils"
	"github.com/py/biz-demo/gomall/common/clientsuite"
	"github.com/py/biz-demo/gomall/rpc_gen/kitex_gen/cart/cartservice"
	"github.com/py/biz-demo/gomall/rpc_gen/kitex_gen/payment/paymentservice"
	"github.com/py/biz-demo/gomall/rpc_gen/kitex_gen/product/productcatalogservice"
)

var (
	CartClient    cartservice.Client
	ProductClient productcatalogservice.Client
	PaymentClient paymentservice.Client
	OrderClient   orderservice.Client
	once          sync.Once
	err           error
	RegistryAddr  string
	ServiceName   string
	commonSuite   client.Option
)

func InitClient() {
	once.Do(func() {
		RegistryAddr = conf.GetConf().Registry.RegistryAddress[0]
		ServiceName = conf.GetConf().Kitex.Service
		commonSuite = client.WithSuite(clientsuite.CommonClientSuite{
			CurrentServiceName: ServiceName,
			RegistryAddr:       RegistryAddr,
		})
		initCartClient()
		initProductClient()
		initPaymentClient()
		initOrderClient()
	})
}

func initProductClient() {
	ProductClient, err = productcatalogservice.NewClient("product", commonSuite)
	checkoututils.MustHandleError(err)
}

func initCartClient() {
	CartClient, err = cartservice.NewClient("cart", commonSuite)
	checkoututils.MustHandleError(err)
}

func initPaymentClient() {
	PaymentClient, err = paymentservice.NewClient("payment", commonSuite)
	checkoututils.MustHandleError(err)
}

func initOrderClient() {
	OrderClient, err = orderservice.NewClient("order", commonSuite)
	checkoututils.MustHandleError(err)
}
