package rpc

import (
	"sync"

	"github.com/cloudwego/kitex/client"
	consul "github.com/kitex-contrib/registry-consul"
	"github.com/py/biz-demo/gomall/app/cart/conf"
	cartUtils "github.com/py/biz-demo/gomall/app/cart/utils"
	"github.com/py/biz-demo/gomall/rpc_gen/kitex_gen/product/productcatalogservice"
)

var (
	ProductClient productcatalogservice.Client
	once          sync.Once
)

func InitClient() {
	once.Do(func() {
		InitProductClient()
	})
}


func InitProductClient() {
	r, err := consul.NewConsulResolver(conf.GetConf().Registry.RegistryAddress[0])
	cartUtils.MustHandleErr(err)
	ProductClient, err = productcatalogservice.NewClient("product", client.WithResolver(r))
	cartUtils.MustHandleErr(err)
}
