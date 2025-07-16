package inventory

import (
	"context"
	serviceModel "inventory/internal/model"
	repoConverter "inventory/internal/repository/converter"
	repoModel "inventory/internal/repository/model"
)

func (r *repository) GetPart(ctx context.Context, GetPartRequest serviceModel.GetPartRequest) (serviceModel.GetPartResponse, error) {
	repoGetPartRequest := repoConverter.ServiceGetPartRequestToRepo(GetPartRequest)
	r.mu.RLock()
	part, ok := r.inventory[repoGetPartRequest.Uuid]
	r.mu.RUnlock()
	if !ok {
		return serviceModel.GetPartResponse{}, serviceModel.ErrPartNotFound
	}

	return repoConverter.RepoGetPartResponseToService(repoModel.GetPartResponse{
		Part: part,
	}), nil
}
