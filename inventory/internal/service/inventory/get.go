package inventory

import (
	"context"
	serviceModel "inventory/internal/model"
)

func (s *service) GetPart(ctx context.Context, GetPartRequest serviceModel.GetPartRequest) (serviceModel.GetPartResponse, error) {
	getPartResponses, err := s.inventoryRepository.GetPart(ctx, GetPartRequest)
	if err != nil {
		return serviceModel.GetPartResponse{}, serviceModel.ErrPartNotFound
	}
	return getPartResponses, nil
}
