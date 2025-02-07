package service

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/py/biz-demo/gomall/app/frontend/infra/rpc"
	"github.com/py/biz-demo/gomall/app/frontend/types"
	frontendUtils "github.com/py/biz-demo/gomall/app/frontend/utils"
	"github.com/py/biz-demo/gomall/rpc_gen/kitex_gen/order"
	"github.com/py/biz-demo/gomall/rpc_gen/kitex_gen/product"
	"strconv"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	common "github.com/py/biz-demo/gomall/app/frontend/hertz_gen/frontend/common"
)

type OrderListService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewOrderListService(Context context.Context, RequestContext *app.RequestContext) *OrderListService {
	return &OrderListService{RequestContext: RequestContext, Context: Context}
}

func (h *OrderListService) Run(req *common.Empty) (resp map[string]interface{}, err error) {
	userId := frontendUtils.GetUserIdFromCtx(h.Context)
	listOrder, err := rpc.OrderClient.ListOrder(h.Context, &order.ListOrderReq{UserId: uint32(userId)})
	if err != nil {
		return nil, err
	}

	var list []types.Order
	for _, v := range listOrder.Orders {
		var (
			total float32
			items []types.OrderItem
		)
		for _, i := range v.Items {
			total += i.Cost

			productResp, err := rpc.ProductClient.GetProduct(h.Context, &product.GetProductReq{Id: i.Item.ProductId})
			if err != nil {
				return nil, err
			}
			if productResp == nil || productResp.Product == nil {
				continue
			}

			items = append(items, types.OrderItem{
				ProductID: strconv.Itoa(int(productResp.Product.Id)),
				Picture:   productResp.Product.Picture,
				Qty:       i.Item.Quantity,
				Cost:      i.Cost,
			})
		}

		created := time.Unix(int64(v.CreatedAt), 0)
		list = append(list, types.Order{
			OrderId:     v.OrderId,
			CreatedDate: created.Format("2006-01-01 15:01:10"),
			Cost:        total,
			Items:       items,
		})
	}

	return utils.H{
		"title":  "Order",
		"orders": list,
	}, nil
}
