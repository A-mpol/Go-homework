package v1

import (
	"context"
	"order/internal/converter"
	order_v1 "shared/pkg/proto/order/v1"
)

func (a *api) Get(ctx context.Context, in *order_v1.GetRequest) (*order_v1.GetResponse, error) {
	serviceGetRequest := converter.GetRequestToModel(in)
	getResponse, err := a.orderService.Get(ctx, serviceGetRequest)
	if err != nil {
		return nil, err
	}
	return converter.ServiceGetResponseToApi(getResponse), nil
}
