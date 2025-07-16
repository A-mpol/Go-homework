package service

import (
	"context"
	serviceModel "inventory/internal/model"
)

type InventoryService interface {
	GetPart(ctx context.Context, GetPartRequest serviceModel.GetPartRequest) (serviceModel.GetPartResponse, error)
	ListParts(ctx context.Context, ListPartsRequest serviceModel.ListPartsRequest) (serviceModel.ListPartsResponse, error)
}
