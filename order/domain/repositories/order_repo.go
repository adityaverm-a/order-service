package repositories

import "demo/oms/order/data/models"

// OrderRepository is...
type OrderRepository interface {
	GetByID(id int64) (*models.Persons, error)
}
