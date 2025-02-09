package rpc

import (
	"sync"

	"github.com/cloudwego/kitex/client"
	"github.com/py/biz-demo/gomall/app/cart/conf"
	cartUtils "github.com/py/biz-demo/gomall/app/cart/utils"
	"github.com/py/biz-demo/gomall/common/clientsuite"
	"github.com/py/biz-demo/gomall/rpc_gen/kitex_gen/product/productcatalogservice"
)

var (
	ProductClient productcatalogservice.Client
	once          sync.Once
	ServiceName   = conf.GetConf().Kitex.Service
	RegistryAddr  = conf.GetConf().Registry.RegistryAddress[0]
	err           error
)

func InitClient() {
	once.Do(func() {
		InitProductClient()
	})
}

func InitProductClient() {
	opt := []client.Option{
		client.WithSuite(clientsuite.CommonClientSuite{
			CurrentServiceName: ServiceName,
			RegistryAddr:       RegistryAddr,
		}),
	}
	ProductClient, err = productcatalogservice.NewClient("product", opt...)
	cartUtils.MustHandleErr(err)
}
