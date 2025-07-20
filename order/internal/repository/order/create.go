package order

import (
	"context"
	serviceModel "order/internal/model"
	"order/internal/repository/converter"
	repoModel "order/internal/repository/model"

	"github.com/google/uuid"
)

func (r *repository) Create(ctx context.Context, in serviceModel.CreateRequest, totalPrice float64) (serviceModel.CreateResponse, error) {
	repoCreateRequest := converter.ServiceCreateRequestToRepo(in)

	orderUuid := uuid.NewString()

	r.mu.Lock()
	r.orders[orderUuid] = repoModel.Order{
		OrderUuid:  orderUuid,
		UserUuid:   repoCreateRequest.UserUuid,
		PartUuids:  repoCreateRequest.PartUuids,
		TotalPrice: totalPrice,
		Status:     repoModel.Status_STATUS_PENDING_PAYMENT,
	}
	r.mu.Unlock()

	return converter.RepoCreateResponseToService(repoModel.CreateResponse{
		OrderUuid:  orderUuid,
		TotalPrice: totalPrice,
	}), nil

}
