package v1

import (
	"order/internal/service"
	order_v1 "shared/pkg/proto/order/v1"
)

type api struct {
	order_v1.UnimplementedOrderServiceServer

	orderService service.OrderService
}

func NewApi(orderService service.OrderService) *api {
	return &api{
		orderService: orderService,
	}
}
