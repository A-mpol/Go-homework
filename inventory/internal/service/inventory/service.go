package inventory

import "inventory/internal/repository"

type service struct {
	inventoryRepository repository.InventoryRepository
}

func (s *service) NewRepository(inventoryRepository repository.InventoryRepository) *service {
	return &service{
		inventoryRepository: inventoryRepository,
	}
}
