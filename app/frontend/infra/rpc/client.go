package rpc

import (
	"context"
	"sync"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/circuitbreak"
	"github.com/cloudwego/kitex/pkg/fallback"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	consulclient "github.com/kitex-contrib/config-consul/client"
	"github.com/kitex-contrib/config-consul/consul"
	"github.com/py/biz-demo/gomall/app/frontend/conf"
	frontendUtils "github.com/py/biz-demo/gomall/app/frontend/utils"
	"github.com/py/biz-demo/gomall/common/clientsuite"
	"github.com/py/biz-demo/gomall/rpc_gen/kitex_gen/cart/cartservice"
	"github.com/py/biz-demo/gomall/rpc_gen/kitex_gen/checkout/checkoutservice"
	"github.com/py/biz-demo/gomall/rpc_gen/kitex_gen/order/orderservice"
	"github.com/py/biz-demo/gomall/rpc_gen/kitex_gen/product"
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
	cbs := circuitbreak.NewCBSuite(func(ri rpcinfo.RPCInfo) string {
		return circuitbreak.RPCInfo2Key(ri)
	})
	cbs.UpdateServiceCBConfig("frontend/product/GetProduct",
		circuitbreak.CBConfig{Enable: true, ErrRate: 0.5, MinSample: 2},
	)
	consulClient, err := consul.NewClient(consul.Options{})
	fb := fallback.NewFallbackPolicy( //配置fallback降级措施
		fallback.UnwrapHelper(func(ctx context.Context, req, resp interface{}, err error) (fbResp interface{}, fbErr error) {
			if err == nil {
				return resp, nil
			}
			methodName := rpcinfo.GetRPCInfo(ctx).To().Method()
			if methodName != "ListProducts" {
				return resp, err
			}
			return &product.ListProductsResp{
				Products: []*product.Product{
					{
						Price:       6.6,
						Id:          3,
						Picture:     "/static/image/t-shirt.jpeg",
						Name:        "T-Shirt",
						Description: "CloudWeGo T-Shirt",
					},
				},
			}, nil
		}),
	)

	opts := []client.Option{
		commonSuite,
		client.WithCircuitBreaker(cbs), //配置熔断策
		client.WithFallback(fb),
		//客户端consul config
		client.WithSuite(consulclient.NewSuite("product", frontendUtils.ServiceName, consulClient)),
	}
	ProductClient, err = productcatalogservice.NewClient("product", opts...)
	frontendUtils.MustHandleErr(err)
}
