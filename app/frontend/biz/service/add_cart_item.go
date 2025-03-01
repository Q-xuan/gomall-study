package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	cart "github.com/py/biz-demo/gomall/app/frontend/hertz_gen/frontend/cart"
	rpccart "github.com/py/biz-demo/gomall/rpc_gen/kitex_gen/cart"
	common "github.com/py/biz-demo/gomall/app/frontend/hertz_gen/frontend/common"
	"github.com/py/biz-demo/gomall/app/frontend/infra/rpc"
	frontendUtils "github.com/py/biz-demo/gomall/app/frontend/utils"
)

type AddCartItemService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewAddCartItemService(Context context.Context, RequestContext *app.RequestContext) *AddCartItemService {
	return &AddCartItemService{RequestContext: RequestContext, Context: Context}
}

func (h *AddCartItemService) Run(req *cart.AddCartItemReq) (resp *common.Empty, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	_,err =rpc.CartClient.AddItem(h.Context,&rpccart.AddItemReq{
		UserId: uint32(frontendUtils.GetUserIdFromCtx(h.Context)),
		Item: &rpccart.CartItem{
			ProductId: req.ProductId,
			Quantity: req.ProductNum,
		},
	})
	if err != nil {
		return nil, err
	}
	return
}
