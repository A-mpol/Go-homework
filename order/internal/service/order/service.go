package order

import (
	"order/internal/repository"
	inventory_v1 "shared/pkg/proto/inventory/v1"
	payment_v1 "shared/pkg/proto/payment/v1"

	def "order/internal/service"
)

var _ def.OrderService = (*service)(nil)

type service struct {
	orderRepository repository.OrderRepository

	inventoryClient inventory_v1.InventoryServiceClient
	paymentClient   payment_v1.PaymentServiceClient
}

func NewService(orderRepository repository.OrderRepository, inventoryClient inventory_v1.InventoryServiceClient, paymentClient payment_v1.PaymentServiceClient) *service {
	return &service{
		orderRepository: orderRepository,
		inventoryClient: inventoryClient,
		paymentClient:   paymentClient,
	}
}
