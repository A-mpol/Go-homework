package order

import (
	"context"
	serviceModel "order/internal/model"
)

func (s *service) Get(ctx context.Context, in serviceModel.GetRequest) (serviceModel.GetResponse, error) {
	serviceGetResponse, err := s.orderRepository.Get(ctx, in)
	if err != nil {
		return serviceModel.GetResponse{}, err
	}

	return serviceGetResponse, nil
}
