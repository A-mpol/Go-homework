package order

import (
	"context"
	serviceModel "order/internal/model"
	"order/internal/repository/converter"
	repoModel "order/internal/repository/model"
)

func (r *repository) Update(ctx context.Context, updateRequest serviceModel.UpdateRequest) error {
	repoUpdateRequest := converter.ServiceUpdateRequestToRepo(updateRequest)
	r.mu.Lock()
	order, _ := r.orders[repoUpdateRequest.OrderUuid]
	if repoUpdateRequest.Status == repoModel.Status_STATUS_CANCELLED {
		order.Status = repoModel.Status_STATUS_CANCELLED
	}
	if repoUpdateRequest.Status == repoModel.Status_STATUS_PAID {
		order.TransactionUuid = repoUpdateRequest.TransactionUuid
		order.PaymentMethod = repoUpdateRequest.PaymentMethod
		order.Status = repoModel.Status_STATUS_PAID
	}
	r.orders[repoUpdateRequest.OrderUuid] = order
	r.mu.Unlock()

	return nil
}
