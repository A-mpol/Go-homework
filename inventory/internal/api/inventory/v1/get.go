package v1

import (
	"context"
	"inventory/internal/converter"
	inventory_v1 "shared/pkg/proto/inventory/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (a *api) GetPart(ctx context.Context, in *inventory_v1.GetPartRequest) (*inventory_v1.GetPartResponse, error) {
	serviceGetPartResponse, err := a.inventoryService.GetPart(ctx, converter.GetPartRequestToModel(in))
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "part with UUID %s not found", in.GetUuid())
	}

	return converter.ServiceGetPartResponseToApi(serviceGetPartResponse), nil
}
