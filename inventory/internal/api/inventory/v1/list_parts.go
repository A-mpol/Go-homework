package v1

import (
	"context"
	"inventory/internal/converter"
	inventory_v1 "shared/pkg/proto/inventory/v1"
)

func (a *api) ListParts(ctx context.Context, in *inventory_v1.ListPartsRequest) (*inventory_v1.ListPartsResponse, error) {
	serviceListPartsResponse, _ := a.inventoryService.ListParts(ctx, converter.ListPartsRequestToModel(in))
	return converter.ServiceListPartsResponseToApi(serviceListPartsResponse), nil
}
