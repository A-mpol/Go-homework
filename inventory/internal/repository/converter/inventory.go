package converter

import (
	"inventory/internal/model"
	repoModel "inventory/internal/repository/model"
)

func FiltersToRepoFilters(serviceFilters *model.Filters) *repoModel.Filters {
	var repoCategories []repoModel.Category
	for _, c := range serviceFilters.Categories {
		repoCategories = append(repoCategories, repoModel.Category(c))
	}

	return &repoModel.Filters{
		Uuids:                 serviceFilters.Uuids,
		Names:                 serviceFilters.Names,
		Categories:            repoCategories,
		ManufacturerCountries: serviceFilters.ManufacturerCountries,
		Tags:                  serviceFilters.Tags,
	}
}

func PartToModel(part *repoModel.Part) *model.Part {
	serviceMetadata := make(map[string]*model.Value, len(part.Metadata))
	for k, v := range part.Metadata {
		if v != nil {
			serviceMetadata[k] = &model.Value{
				StringValue: v.StringValue,
				Int64Value:  v.Int64Value,
				DoubleValue: v.DoubleValue,
				BoolValue:   v.BoolValue,
			}
		} else {
			serviceMetadata[k] = nil
		}
	}
	return &model.Part{
		Uuid:          part.Uuid,
		Name:          part.Name,
		Description:   part.Description,
		Price:         part.Price,
		StockQuantity: part.StockQuantity,
		Category:      (*model.Category)(part.Category),
		Dimensions:    (*model.Dimensions)(part.Dimensions),
		Manufacturer:  (*model.Manufacturer)(part.Manufacturer),
		Tags:          part.Tags,
		Metadata:      serviceMetadata,
		CreatedAt:     part.CreatedAt,
		UpdatedAt:     part.UpdatedAt,
	}
}

func ListPartsToModel(listParts *repoModel.ListParts) *model.ListParts {
	serviceParts := make([]*model.Part, 0, len(listParts.Parts))
	for _, p := range listParts.Parts {
		if p == nil {
			serviceParts = append(serviceParts, nil)
		} else {
			serviceParts = append(serviceParts, PartToModel(p))
		}
	}
	return &model.ListParts{
		Parts: serviceParts,
	}
}
