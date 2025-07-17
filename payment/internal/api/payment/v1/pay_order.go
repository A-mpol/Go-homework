package v1

import (
	"context"
	payment_v1 "shared/pkg/proto/payment/v1"
)

func (a *api) PayOrder(ctx context.Context, in *payment_v1.PayOrderRequest) (*payment_v1.PayOrderResponse, error) {
	payOrderResponse, err := a.paymentService.PayOrder(ctx, in)
	if err != nil {
		return nil, err
	}
	return payOrderResponse, nil
}
