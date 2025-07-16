package converter

import (
	serviceModel "inventory/internal/model"
	inventory_v1 "shared/pkg/proto/inventory/v1"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func GetPartRequestToModel(in *inventory_v1.GetPartRequest) serviceModel.GetPartRequest {
	return serviceModel.GetPartRequest{
		Uuid: in.Uuid,
	}
}

func ServiceGetPartResponseToApi(serviceModelGetPartResponse serviceModel.GetPartResponse) *inventory_v1.GetPartResponse {
	return &inventory_v1.GetPartResponse{
		Part: ServicePartToApi(serviceModelGetPartResponse.Part),
	}
}

func ServicePartToApi(serviceModelPart serviceModel.Part) *inventory_v1.Part {
	return &inventory_v1.Part{
		Uuid:          serviceModelPart.Uuid,
		Name:          serviceModelPart.Name,
		Description:   serviceModelPart.Description,
		Price:         serviceModelPart.Price,
		StockQuantity: serviceModelPart.StockQuantity,
		Category:      inventory_v1.Category(serviceModelPart.Category),
		Dimensions:    ServiceDimensionsToApi(serviceModelPart.Dimensions),
		Manufacturer:  ServiceManufacturerToApi(serviceModelPart.Manufacturer),
		Tags:          serviceModelPart.Tags,
		Metadata:      ServiceMetadataToApi(serviceModelPart.Metadata),
		CreatedAt:     ServiceTimeToApi(serviceModelPart.CreatedAt),
		UpdatedAt:     ServiceTimeToApi(serviceModelPart.UpdatedAt),
	}
}

func ServiceDimensionsToApi(serviceModelDimensions serviceModel.Dimensions) *inventory_v1.Dimensions {
	return &inventory_v1.Dimensions{
		Length: serviceModelDimensions.Length,
		Width:  serviceModelDimensions.Width,
		Height: serviceModelDimensions.Height,
		Weight: serviceModelDimensions.Weight,
	}
}

func ServiceManufacturerToApi(serviceModelManufacturer serviceModel.Manufacturer) *inventory_v1.Manufacturer {
	return &inventory_v1.Manufacturer{
		Name:    serviceModelManufacturer.Name,
		Country: serviceModelManufacturer.Country,
		Website: serviceModelManufacturer.Website,
	}
}

func ServiceMetadataToApi(meta map[string]serviceModel.Value) map[string]*inventory_v1.Value {
	if meta == nil {
		return nil
	}
	result := make(map[string]*inventory_v1.Value, len(meta))
	for k, v := range meta {
		result[k] = ServiceValueToApi(v)
	}
	return result
}

func ServiceValueToApi(v serviceModel.Value) *inventory_v1.Value {
	switch {
	case v.StringValue != nil:
		return &inventory_v1.Value{
			Kind: &inventory_v1.Value_StringValue{StringValue: *v.StringValue},
		}
	case v.Int64Value != nil:
		return &inventory_v1.Value{
			Kind: &inventory_v1.Value_Int64Value{Int64Value: *v.Int64Value},
		}
	case v.DoubleValue != nil:
		return &inventory_v1.Value{
			Kind: &inventory_v1.Value_DoubleValue{DoubleValue: *v.DoubleValue},
		}
	case v.BoolValue != nil:
		return &inventory_v1.Value{
			Kind: &inventory_v1.Value_BoolValue{BoolValue: *v.BoolValue},
		}
	default:
		return nil
	}
}

func ServiceTimeToApi(serviceTime time.Time) *timestamppb.Timestamp {
	if !serviceTime.IsZero() {
		return timestamppb.New(serviceTime)
	}
	return nil
}

func ListPartsRequestToModel(in *inventory_v1.ListPartsRequest) serviceModel.ListPartsRequest {
	return serviceModel.ListPartsRequest{
		Filter: serviceModel.PartsFilter{
			Uuids:                 in.Filter.Uuids,
			Names:                 in.Filter.Names,
			Categories:            ApiCategoriesToService(in.Filter.Categories),
			ManufacturerCountries: in.Filter.ManufacturerCountries,
			Tags:                  in.Filter.Tags,
		},
	}
}

func ApiCategoriesToService(src []inventory_v1.Category) []serviceModel.Category {
	dst := make([]serviceModel.Category, len(src))
	for i, cat := range src {
		dst[i] = serviceModel.Category(cat)
	}
	return dst
}

func ServiceListPartsResponseToApi(serviceModelListPartsResponse serviceModel.ListPartsResponse) *inventory_v1.ListPartsResponse {
	return &inventory_v1.ListPartsResponse{
		Parts: ServicePartsToApi(serviceModelListPartsResponse.Parts),
	}
}

func ServicePartsToApi(src []serviceModel.Part) []*inventory_v1.Part {
	dst := make([]*inventory_v1.Part, len(src))
	for i, item := range src {
		dst[i] = ServicePartToApi(item)
	}
	return dst
}
