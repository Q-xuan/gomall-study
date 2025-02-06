package service

import (
	"context"

	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/google/uuid"
	"github.com/py/biz-demo/gomall/app/checkout/infra/rpc"
	"github.com/py/biz-demo/gomall/rpc_gen/kitex_gen/cart"
	checkout "github.com/py/biz-demo/gomall/rpc_gen/kitex_gen/checkout"
	"github.com/py/biz-demo/gomall/rpc_gen/kitex_gen/payment"
	"github.com/py/biz-demo/gomall/rpc_gen/kitex_gen/product"
)

type CheckoutService struct {
	ctx context.Context
} // NewCheckoutService new CheckoutService
func NewCheckoutService(ctx context.Context) *CheckoutService {
	return &CheckoutService{ctx: ctx}
}

// Run create note info
func (s *CheckoutService) Run(req *checkout.CheckoutReq) (resp *checkout.CheckoutResp, err error) {
	// Finish your business logic.
	cartResp, err := rpc.CardClient.GetCart(s.ctx, &cart.GetCartReq{UserId: req.UserId})
	if err != nil {
		return nil, kerrors.NewGRPCBizStatusError(5005001, err.Error())
	}
	if cartResp == nil || cartResp.Items == nil {
		return nil, kerrors.NewGRPCBizStatusError(5004001, "cart is empty")
	}
	var total float32

	for _, cartItem := range cartResp.Items {
		productResp, err := rpc.ProductClient.GetProduct(s.ctx, &product.GetProductReq{
			Id: cartItem.ProductId,
		})
		if err != nil {
			return nil, err
		}
		if productResp.Product == nil {
			continue
		}
		p := productResp.Product.Price
		cost := p * float32(cartItem.Quantity)
		total += cost
	}

	var orderId string
	u, _ := uuid.NewRandom()
	orderId = u.String()

	payReq := &payment.ChargeReq{
		UserId:  req.UserId,
		OrderId: orderId,
		Amount:  total,
		CreditCard: &payment.CreditCardInfo{
			CreditCardNumber:          req.CreditCard.CreditCardNumber,
			CreditCardCvv:             req.CreditCard.CreditCardCvv,
			CreditCardExpirationMonth: req.CreditCard.CreditCardExpirationMonth,
			CreditCardExpirationYear:  req.CreditCard.CreditCardExpirationYear,
		},
	}

	_, err = rpc.CardClient.EmptyCart(s.ctx, &cart.EmptyCartReq{UserId: req.UserId})
	if err != nil {
		return nil, err
	}

	paymentResp, err := rpc.PaymentClient.Charge(s.ctx, payReq)
	if err != nil {
		return nil, err
	}
	klog.Info(paymentResp)

	resp = &checkout.CheckoutResp{
		OrderId:       orderId,
		TransactionId: paymentResp.TransactionId,
	}
	return
}
