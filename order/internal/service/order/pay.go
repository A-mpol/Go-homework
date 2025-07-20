package order

import (
	"context"
	serviceModel "order/internal/model"
	payment_v1 "shared/pkg/proto/payment/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *service) Pay(ctx context.Context, in serviceModel.PayRequest) (serviceModel.PayResponse, error) {
	getResponce, err := s.orderRepository.Get(ctx, serviceModel.GetRequest{OrderUuid: in.OrderUuid})
	if err != nil {
		return serviceModel.PayResponse{}, err
	}

	paymentInformation, err := s.paymentClient.PayOrder(ctx, &payment_v1.PayOrderRequest{
		OrderUuid:     in.OrderUuid,
		UserUuid:      getResponce.Order.UserUuid,
		PaymentMethod: payment_v1.PaymentMethod(in.PaymentMethod),
	})
	if err != nil {
		return serviceModel.PayResponse{}, status.Errorf(codes.Internal, "Failed to pay")
	}

	s.orderRepository.Update(ctx, serviceModel.UpdateRequest{
		OrderUuid:       in.OrderUuid,
		TransactionUuid: paymentInformation.TransactionUuid,
		PaymentMethod:   in.PaymentMethod,
		Status:          serviceModel.Status_STATUS_PAID,
	})

	return serviceModel.PayResponse{
		TransactionUuid: paymentInformation.TransactionUuid,
	}, nil
}
