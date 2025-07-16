package converter

import (
	serviceModel "inventory/internal/model"
	repoModel "inventory/internal/repository/model"
)

func ServiceGetPartRequestToRepo(serviceModelGetPartRequest serviceModel.GetPartRequest) repoModel.GetPartRequest {
	return repoModel.GetPartRequest{
		Uuid: serviceModelGetPartRequest.Uuid,
	}
}

func RepoGetPartResponseToService(repoModelGetPartResponse repoModel.GetPartResponse) serviceModel.GetPartResponse {
	return serviceModel.GetPartResponse{
		Part: RepoPartToService(repoModelGetPartResponse.Part),
	}
}

func RepoPartToService(repoModelPart repoModel.Part) serviceModel.Part {
	return serviceModel.Part{
		Uuid:          repoModelPart.Uuid,
		Name:          repoModelPart.Name,
		Description:   repoModelPart.Description,
		Price:         repoModelPart.Price,
		StockQuantity: repoModelPart.StockQuantity,
		Category:      serviceModel.Category(repoModelPart.Category),
		Dimensions:    serviceModel.Dimensions(repoModelPart.Dimensions),
		Manufacturer:  serviceModel.Manufacturer(repoModelPart.Manufacturer),
		Tags:          repoModelPart.Tags,
		Metadata:      RepoMetadataToService(repoModelPart.Metadata),
		CreatedAt:     repoModelPart.CreatedAt,
		UpdatedAt:     repoModelPart.UpdatedAt,
	}
}

func RepoValueToService(v repoModel.Value) serviceModel.Value {
	return serviceModel.Value{
		StringValue: v.StringValue,
		Int64Value:  v.Int64Value,
		DoubleValue: v.DoubleValue,
		BoolValue:   v.BoolValue,
	}
}

func RepoMetadataToService(src map[string]repoModel.Value) map[string]serviceModel.Value {
	dst := make(map[string]serviceModel.Value, len(src))
	for k, v := range src {
		dst[k] = RepoValueToService(v)
	}
	return dst
}

func ServiceListPartsRequestToRepo(serviceModelListPartsRequest serviceModel.ListPartsRequest) repoModel.ListPartsRequest {
	return repoModel.ListPartsRequest{
		Filter: repoModel.PartsFilter{
			Uuids:                 serviceModelListPartsRequest.Filter.Uuids,
			Names:                 serviceModelListPartsRequest.Filter.Names,
			Categories:            ServiceCategoriesToRepo(serviceModelListPartsRequest.Filter.Categories),
			ManufacturerCountries: serviceModelListPartsRequest.Filter.ManufacturerCountries,
			Tags:                  serviceModelListPartsRequest.Filter.Tags,
		},
	}
}

func ServiceCategoriesToRepo(src []serviceModel.Category) []repoModel.Category {
	dst := make([]repoModel.Category, len(src))
	for i, cat := range src {
		dst[i] = repoModel.Category(cat)
	}
	return dst
}

func RepoListPartsResponseToService(repoModelListPartsResponse repoModel.ListPartsResponse) serviceModel.ListPartsResponse {
	return serviceModel.ListPartsResponse{
		Parts: RepoPartsToService(repoModelListPartsResponse.Parts),
	}
}

func RepoPartsToService(src []repoModel.Part) []serviceModel.Part {
	dst := make([]serviceModel.Part, len(src))
	for i, item := range src {
		dst[i] = RepoPartToService(item)
	}
	return dst
}
