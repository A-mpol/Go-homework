package inventory

import (
	"context"
	serviceModel "inventory/internal/model"
	repoConverter "inventory/internal/repository/converter"
	repoModel "inventory/internal/repository/model"
)

func (r *repository) ListParts(ctx context.Context, ListPartsRequest serviceModel.ListPartsRequest) (serviceModel.ListPartsResponse, error) {
	repoListPartsRequest := repoConverter.ServiceListPartsRequestToRepo(ListPartsRequest)

	listParts := []repoModel.Part{}
	r.mu.RLock()
	for uuid := range r.inventory {
		fitsParameter := true
		for i, checkUuid := range repoListPartsRequest.Filter.Uuids {
			if uuid == checkUuid {
				break
			}
			if uuid != checkUuid && i == len(repoListPartsRequest.Filter.Uuids)-1 {
				fitsParameter = false
			}
		}

		for i, checkName := range repoListPartsRequest.Filter.Names {
			if !fitsParameter || r.inventory[uuid].Name == checkName {
				break
			}
			if r.inventory[uuid].Name != checkName && i == len(repoListPartsRequest.Filter.Names)-1 {
				fitsParameter = false
			}
		}

		for i, checkCategory := range repoListPartsRequest.Filter.Categories {
			if !fitsParameter || r.inventory[uuid].Category == checkCategory {
				break
			}
			if r.inventory[uuid].Category != checkCategory && i == len(repoListPartsRequest.Filter.Categories)-1 {
				fitsParameter = false
			}
		}

		for i, checkManufacturerCountrie := range repoListPartsRequest.Filter.ManufacturerCountries {
			if !fitsParameter || r.inventory[uuid].Manufacturer.Country == checkManufacturerCountrie {
				break
			}
			if r.inventory[uuid].Manufacturer.Country != checkManufacturerCountrie && i == len(repoListPartsRequest.Filter.ManufacturerCountries)-1 {
				fitsParameter = false
			}
		}

		for i, checkTag := range repoListPartsRequest.Filter.Tags {
			if !fitsParameter {
				break
			}

			noSuchTag := true
			for _, tag := range r.inventory[uuid].Tags {
				if checkTag == tag {
					noSuchTag = false
					break
				}
			}
			if i == len(repoListPartsRequest.Filter.Tags)-1 && noSuchTag {
				fitsParameter = false
			}
		}

		if fitsParameter {
			listParts = append(listParts, r.inventory[uuid])
		}
	}
	r.mu.RUnlock()

	return repoConverter.RepoListPartsResponseToService(repoModel.ListPartsResponse{
		Parts: listParts,
	}), nil
}
