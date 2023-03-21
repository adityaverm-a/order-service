package repositories

import (
	"demo/oms/order/data/models"
	"demo/oms/order/domain/entities"
)

// OrderRepository is...
//
//go:generate mockgen -destination=mocks/mock_order_repo.go -package=mocks demo/oms/order/domain/repositories OrderRepository
type OrderRepository interface {
	GetByID(id int64) (*models.Order, error)
	GetByFilters(filters entities.OrderFiltersInput) (*[]models.Order, error)
	Create(input entities.CreateOrderInput) (*models.Order, error)
	Update(input entities.UpdateOrderStatusInput) (*models.Order, error)
}
