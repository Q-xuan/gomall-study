package rpc

import (
	"sync"

	"github.com/cloudwego/kitex/client"
	consul "github.com/kitex-contrib/registry-consul"
	"github.com/py/biz-demo/gomall/app/frontend/conf"
	frontendUtils "github.com/py/biz-demo/gomall/app/frontend/utils"
	"github.com/py/biz-demo/gomall/rpc_gen/kitex_gen/cart/cartservice"
	"github.com/py/biz-demo/gomall/rpc_gen/kitex_gen/product/productcatalogservice"
	"github.com/py/biz-demo/gomall/rpc_gen/kitex_gen/user/userservice"
)

var (
	UserClient    userservice.Client
	ProductClient productcatalogservice.Client
	CartClient    cartservice.Client
	once          sync.Once
)

func Init() {
	once.Do(func() {
		InitUserClient()
		InitProductClient()
		InitCartClient()
	})
}

func InitCartClient() {
	var opts []client.Option
	r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddr)
	frontendUtils.MustHandleErr(err)
	opts = append(opts, client.WithResolver(r))
	CartClient, err = cartservice.NewClient("cart", opts...)
	frontendUtils.MustHandleErr(err)
}

func InitUserClient() {
	var opts []client.Option
	r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddr)
	frontendUtils.MustHandleErr(err)
	opts = append(opts, client.WithResolver(r))
	UserClient, err = userservice.NewClient("user", opts...)
	frontendUtils.MustHandleErr(err)
}

func InitProductClient() {
	r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddr)
	frontendUtils.MustHandleErr(err)
	ProductClient, err = productcatalogservice.NewClient("product", client.WithResolver(r))
	frontendUtils.MustHandleErr(err)
}
