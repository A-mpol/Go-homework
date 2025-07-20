package service

import (
	"context"
	serviceModel "order/internal/model"
)

type OrderService interface {
	Create(ctx context.Context, in serviceModel.CreateRequest) (serviceModel.CreateResponse, error)
	Pay(ctx context.Context, in serviceModel.PayRequest) (serviceModel.PayResponse, error)
	Get(ctx context.Context, in serviceModel.GetRequest) (serviceModel.GetResponse, error)
	Cancel(ctx context.Context, in serviceModel.CancelRequest) error
}
