package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/py/biz-demo/gomall/app/frontend/hertz_gen/frontend/common"
	"github.com/py/biz-demo/gomall/app/frontend/infra/rpc"
	"github.com/py/biz-demo/gomall/rpc_gen/kitex_gen/product"
)

type HomeService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewHomeService(Context context.Context, RequestContext *app.RequestContext) *HomeService {
	return &HomeService{RequestContext: RequestContext, Context: Context}
}

func (h *HomeService) Run(req *common.Empty) (data map[string]any, err error) {
	products, err := rpc.ProductClient.ListProducts(h.Context, &product.ListProductsReq{})
	if err != nil {
		return nil, err
	}
	//items := []map[string]any{
	//	{"Name": "T-shirt-1", "Price": 100, "Picture": "https://fakeimg.pl/200x300/?text=T-shirt-1"},
	//	{"Name": "T-shirt-2", "Price": 110, "Picture": "https://fakeimg.pl/200x300/?text=T-shirt-2"},
	//	{"Name": "T-shirt-3", "Price": 120, "Picture": "https://fakeimg.pl/200x300/?text=T-shirt-3"},
	//	{"Name": "T-shirt-4", "Price": 130, "Picture": "https://fakeimg.pl/200x300/?text=T-shirt-4"},
	//	{"Name": "T-shirt-5", "Price": 140, "Picture": "https://fakeimg.pl/200x300/?text=T-shirt-5"},
	//	{"Name": "T-shirt-6", "Price": 150, "Picture": "https://fakeimg.pl/200x300/?text=T-shirt-6"},
	//}
	return utils.H{
		"Title": "Hot sale",
		"Items": products.Products,
	}, nil
}
