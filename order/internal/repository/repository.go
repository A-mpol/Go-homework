package repository

import (
	"context"
	serviceModel "order/internal/model"
)

type OrderRepository interface {
	Create(ctx context.Context, in serviceModel.CreateRequest, totalPrice float64) (serviceModel.CreateResponse, error)
	Get(ctx context.Context, in serviceModel.GetRequest) (serviceModel.GetResponse, error)
	Update(ctx context.Context, updateRequest serviceModel.UpdateRequest) error
}
