package order

import (
	"context"
	serviceModel "order/internal/model"
	"order/internal/repository/converter"
	repoModel "order/internal/repository/model"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (r *repository) Get(ctx context.Context, in serviceModel.GetRequest) (serviceModel.GetResponse, error) {
	repoGetRequest := converter.ServiceGetRequestToRepo(in)
	r.mu.RLock()
	order, ok := r.orders[repoGetRequest.OrderUuid]
	r.mu.RUnlock()
	if !ok {
		return serviceModel.GetResponse{}, status.Errorf(codes.NotFound, "Order not found")
	}

	return converter.RepoGetResponseToService(repoModel.GetResponse{
		Order: order,
	}), nil
}
