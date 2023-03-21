package repositories

import (
	"demo/oms/order/data/constants"
	"demo/oms/order/data/models"
	"demo/oms/order/domain/repositories"
	"errors"

	"gorm.io/gorm"
)

// The NewOrderRepository function is a factory function that returns a new instance of the orderRepo
func NewOrderRepository(db *gorm.DB) repositories.OrderRepository {
	return &orderRepo{db: db}
}

type orderRepo struct {
	db *gorm.DB
}

// GetByID fetches order with the provided id.
func (or *orderRepo) GetByID(id int64) (models.Order, error) {
	order := models.Order{}

	db := or.db.Table(constants.TABLE_NAME_ORDERS)
	err := db.Preload("OrderItem").Where("orders.id = ?", id).Find(&order).Error

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return order, err
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return order, nil
	}

	return order, nil
}
