package repository

import (
	"context"
	"inventory/internal/model"
)

type InventoryRepository interface {
	GetPart(ctx context.Context, uuid string) (*model.Part, error)
	ListParts(ctx context.Context, filtres *model.Filters) (*model.ListParts, error)
}
