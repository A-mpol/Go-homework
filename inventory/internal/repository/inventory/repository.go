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
	return &repository{
		inventory: make(map[string]repoModel.Part),
	}
}
