package v1

import (
	"context"
	"order/internal/converter"
	order_v1 "shared/pkg/proto/order/v1"
)

func (a *api) Pay(ctx context.Context, in *order_v1.PayRequest) (*order_v1.PayResponse, error) {
	servicePayRequest := converter.PayRequestToModel(in)
	payResponse, err := a.orderService.Pay(ctx, servicePayRequest)
	if err != nil {
		return nil, err
	}

	return converter.ServicePayResponseToApi(payResponse), nil
}
