package inventory

import (
	"context"
	serviceModel "inventory/internal/model"
)

func (s *service) ListParts(ctx context.Context, ListPartsRequest serviceModel.ListPartsRequest) (serviceModel.ListPartsResponse, error) {
	listPartsResponse, _ := s.inventoryRepository.ListParts(ctx, ListPartsRequest)
	return listPartsResponse, nil
}
