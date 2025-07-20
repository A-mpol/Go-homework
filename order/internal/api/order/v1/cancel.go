package v1

import (
	"context"
	"order/internal/converter"
	order_v1 "shared/pkg/proto/order/v1"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (a *api) Cancel(ctx context.Context, in *order_v1.CancelRequest) (*emptypb.Empty, error) {
	serviceCancelRequest := converter.CancelRequestToModel(in)

	if err := a.orderService.Cancel(ctx, serviceCancelRequest); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
