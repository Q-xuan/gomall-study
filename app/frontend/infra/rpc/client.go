package rpc

import (
	"sync"

	"github.com/py/biz-demo/gomall/rpc_gen/kitex_gen/order/orderservice"

	"github.com/cloudwego/kitex/client"
	"github.com/py/biz-demo/gomall/app/frontend/conf"
	frontendUtils "github.com/py/biz-demo/gomall/app/frontend/utils"
	"github.com/py/biz-demo/gomall/common/clientsuite"
	"github.com/py/biz-demo/gomall/rpc_gen/kitex_gen/cart/cartservice"
	"github.com/py/biz-demo/gomall/rpc_gen/kitex_gen/checkout/checkoutservice"
	"github.com/py/biz-demo/gomall/rpc_gen/kitex_gen/product/productcatalogservice"
	"github.com/py/biz-demo/gomall/rpc_gen/kitex_gen/user/userservice"
)

var (
	UserClient     userservice.Client
	ProductClient  productcatalogservice.Client
	CartClient     cartservice.Client
	CheckoutClinet checkoutservice.Client
	OrderClient    orderservice.Client
	once           sync.Once
	err            error
	registryAddr   string
	commonSuite    client.Option
)

func Init() {
	once.Do(func() {
		registryAddr = conf.GetConf().Hertz.RegistryAddr
		commonSuite = client.WithSuite(clientsuite.CommonClientSuite{
			RegistryAddr:       registryAddr,
			CurrentServiceName: frontendUtils.ServiceName,
		})
		InitUserClient()
		InitProductClient()
		InitCartClient()
		InitCheckoutClient()
		InitOrderClient()
	})
}

func InitOrderClient() {
	opts := []client.Option{
		commonSuite,
	}
	OrderClient, err = orderservice.NewClient("order", opts...)
	frontendUtils.MustHandleErr(err)
}

func InitCheckoutClient() {
	opts := []client.Option{
		commonSuite,
	}
	CheckoutClinet, err = checkoutservice.NewClient("checkout", opts...)
	frontendUtils.MustHandleErr(err)
}

func InitCartClient() {
	opts := []client.Option{
		commonSuite,
	}
	CartClient, err = cartservice.NewClient("cart", opts...)
	frontendUtils.MustHandleErr(err)
}

func InitUserClient() {
	opts := []client.Option{
		commonSuite,
	}
	UserClient, err = userservice.NewClient("user", opts...)
	frontendUtils.MustHandleErr(err)
}

func InitProductClient() {
	opts := []client.Option{
		commonSuite,
	}
	ProductClient, err = productcatalogservice.NewClient("product", opts...)
	frontendUtils.MustHandleErr(err)
}
