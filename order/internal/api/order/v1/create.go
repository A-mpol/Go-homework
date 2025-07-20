package v1

import (
	"context"
	"order/internal/converter"
	order_v1 "shared/pkg/proto/order/v1"
)

func (a *api) Create(ctx context.Context, in *order_v1.CreateRequest) (*order_v1.CreateResponse, error) {
	serviceCreateRequest := converter.CreateRequestToModel(in)
	createResponse, err := a.orderService.Create(ctx, serviceCreateRequest)
	if err != nil {
		return nil, err
	}
	return converter.ServiceCreateResponseToApi(createResponse), nil
}
