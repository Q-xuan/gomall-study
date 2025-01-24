package service

import (
	"context"

	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/py/biz-demo/gomall/app/cart/rpc"
	cart "github.com/py/biz-demo/gomall/rpc_gen/kitex_gen/cart"
	"github.com/py/biz-demo/gomall/rpc_gen/kitex_gen/product"
)

type AddItemService struct {
	ctx context.Context
} // NewAddItemService new AddItemService
func NewAddItemService(ctx context.Context) *AddItemService {
	return &AddItemService{ctx: ctx}
}

// Run create note info
func (s *AddItemService) Run(req *cart.AddItemReq) (resp *cart.AddItemResp, err error) {
	// Finish your business logic.
	r, err := rpc.ProductClient.GetProduct(s.ctx, &product.GetProductReq{Id: req.Item.ProductId})
	if err != nil {
		return nil, err
	}
	if r == nil || r.Product.Id == 0 {
		return nil, kerrors.NewBizStatusError(40004, "product not found")
	}
	return
}
