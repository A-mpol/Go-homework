package order

import (
	def "order/internal/repository"
	repoModel "order/internal/repository/model"
	"sync"
)

var _ def.OrderRepository = (*repository)(nil)

type repository struct {
	mu     sync.RWMutex
	orders map[string]repoModel.Order
}

func NewRepository() *repository {
	return &repository{
		orders: make(map[string]repoModel.Order),
	}
}
