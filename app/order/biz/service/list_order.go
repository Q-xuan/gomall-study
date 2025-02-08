package service

import (
	"context"

	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/py/biz-demo/gomall/app/order/biz/dal/mysql"
	"github.com/py/biz-demo/gomall/app/order/biz/model"
	"github.com/py/biz-demo/gomall/rpc_gen/kitex_gen/cart"
	order "github.com/py/biz-demo/gomall/rpc_gen/kitex_gen/order"
)

type ListOrderService struct {
	ctx context.Context
} // NewListOrderService new ListOrderService
func NewListOrderService(ctx context.Context) *ListOrderService {
	return &ListOrderService{ctx: ctx}
}

// Run create note info
func (s *ListOrderService) Run(req *order.ListOrderReq) (resp *order.ListOrderResp, err error) {
	// Finish your business logic.
	list, err := model.ListOrder(s.ctx, mysql.DB, req.UserId)
	if err != nil {
		return nil, kerrors.NewBizStatusError(500001, err.Error())
	}

	var orders []*order.Order
	for _, v := range list {
		var items []*order.OrderItem
		for _, i := range v.OrderItems {
			items = append(items, &order.OrderItem{
				Item: &cart.CartItem{
					ProductId: i.ProductId,
					Quantity:  i.Quantity,
				},
				Cost: i.Cost,
			})
		}
		orders = append(orders, &order.Order{
			OrderId:      v.OrderId,
			UserId:       v.UserId,
			UserCurrency: v.UserCurrency,
			Email:        v.Consignee.Email,
			Address: &order.Address{
				Street:  v.Consignee.Street,
				Country: v.Consignee.Country,
				State:   v.Consignee.State,
				ZipCode: v.Consignee.ZipCode,
			},
			Items: items,
			CreatedAt: int32(v.CreatedAt.UnixMilli()),
		})

	}
	resp = &order.ListOrderResp{
		Orders: orders,
	}
	return
}
