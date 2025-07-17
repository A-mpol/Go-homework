package payment

import (
	"context"
	"log"
	payment_v1 "shared/pkg/proto/payment/v1"

	"github.com/google/uuid"
)

func (s *service) PayOrder(ctx context.Context, in *payment_v1.PayOrderRequest) (*payment_v1.PayOrderResponse, error) {
	uuid := uuid.NewString()
	log.Printf("Оплата прошла успешно, transaction_uuid: %s", uuid)
	return &payment_v1.PayOrderResponse{
		TransactionUuid: uuid,
	}, nil
}
