package order

import (
	"context"
	serviceModel "order/internal/model"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *service) Cancel(ctx context.Context, in serviceModel.CancelRequest) error {
	getResponce, err := s.orderRepository.Get(ctx, serviceModel.GetRequest{
		OrderUuid: in.OrderUuid,
	})
	if err != nil {
		return status.Errorf(codes.NotFound, "Order not found")
	}

	if getResponce.Order.Status == serviceModel.Status_STATUS_PAID {
		return status.Errorf(codes.AlreadyExists, "Order already paid, cannot cancel")
	}

	s.orderRepository.Update(ctx, serviceModel.UpdateRequest{
		OrderUuid: in.OrderUuid,
		Status:    serviceModel.Status_STATUS_CANCELLED,
	})

	return nil
}
