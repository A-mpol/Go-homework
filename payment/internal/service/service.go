package service

import (
	"context"
	payment_v1 "shared/pkg/proto/payment/v1"
)

type PaymentService interface {
	PayOrder(ctx context.Context, in *payment_v1.PayOrderRequest) (*payment_v1.PayOrderResponse, error)
}
