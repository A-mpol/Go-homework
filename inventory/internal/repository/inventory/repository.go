package inventory

import (
	def "inventory/internal/repository"
	repoModel "inventory/internal/repository/model"
	"sync"
)

var _ def.InventoryRepository = (*repository)(nil)

type repository struct {
	mu        sync.RWMutex
	inventory map[string]repoModel.Part
}

func NewRepository() *repository {
	sensor1 := repoModel.Part{
		Uuid:          "1",
		Name:          "alex",
		Price:         100.0,
		StockQuantity: 15,
		Category:      repoModel.Category_CATEGORY_ENGINE,
		Dimensions: repoModel.Dimensions{
			Length: 0,
			Width:  10,
			Weight: 3,
			Height: 1,
		},
		Manufacturer: repoModel.Manufacturer{
			Name:    "spb",
			Country: "russia",
			Website: "www",
		},
	}
	sensor2 := repoModel.Part{
		Uuid:          "2",
		Name:          "consani",
		Price:         100.0,
		StockQuantity: 15,
		Category:      repoModel.Category_CATEGORY_ENGINE,
		Dimensions: repoModel.Dimensions{
			Length: 0,
			Width:  10,
			Weight: 3,
			Height: 1,
		},
		Manufacturer: repoModel.Manufacturer{
			Name:    "spb",
			Country: "russia",
			Website: "www",
		},
	}
	inv := map[string]repoModel.Part{"1": sensor1, "2": sensor2}
	return &repository{
		inventory: inv,
	}
}
