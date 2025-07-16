package v1

import (
	"inventory/internal/service"
	inventory_v1 "shared/pkg/proto/inventory/v1"
)

type api struct {
	inventory_v1.UnimplementedInventoryServiceServer

	inventoryService service.InventoryService
}

func NewApi(inventoryService service.InventoryService) *api {
	return &api{
		inventoryService: inventoryService,
	}
}
