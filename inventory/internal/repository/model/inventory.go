package model

import "time"

type GetPartRequest struct {
	Uuid string
}

type GetPartResponse struct {
	Part Part
}

type ListPartsRequest struct {
	Filter PartsFilter
}

type ListPartsResponse struct {
	Parts []Part
}

type Part struct {
	Uuid          string
	Name          string
	Description   string
	Price         float64
	StockQuantity int64
	Category      Category
	Dimensions    Dimensions
	Manufacturer  Manufacturer
	Tags          []string
	Metadata      map[string]Value
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type PartsFilter struct {
	Uuids                 []string
	Names                 []string
	Categories            []Category
	ManufacturerCountries []string
	Tags                  []string
}

type Category int32

const (
	Category_CATEGORY_UNKNOWN_UNSPECIFIED Category = 0
	Category_CATEGORY_ENGINE              Category = 1
	Category_CATEGORY_FUEL                Category = 2
	Category_CATEGORY_PORTHOLE            Category = 3
	Category_CATEGORY_WING                Category = 4
)

type Dimensions struct {
	Length float64
	Width  float64
	Height float64
	Weight float64
}

type Manufacturer struct {
	Name    string
	Country string
	Website string
}

type Value struct {
	StringValue *string
	Int64Value  *int64
	DoubleValue *float64
	BoolValue   *bool
}
