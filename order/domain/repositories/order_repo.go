package repositories

import (
	"demo/oms/order/data/models"
	"demo/oms/order/domain/entities"
)

// OrderRepository is...
type OrderRepository interface {
	GetByID(id int64) (*models.Order, error)
	Create(input entities.CreateOrderInput) (*models.Order, error)
	Update(input entities.UpdateOrderStatusInput) (*models.Order, error)
}
