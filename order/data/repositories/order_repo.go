package repositories

import (
	"demo/oms/order/data/constants"
	"demo/oms/order/data/models"
	"demo/oms/order/domain/entities"
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
func (or *orderRepo) GetByID(id int64) (*models.Order, error) {
	order := models.Order{}

	db := or.db.Table(constants.TABLE_NAME_ORDERS)
	err := db.Preload("OrderItem").Where("orders.id = ?", id).Find(&order).Error

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return &order, nil
}

// Create adds an order to the database.
func (or *orderRepo) Create(input entities.CreateOrderInput) (*models.Order, error) {
	// create a new order from input parameters
	order := models.Order{
		Total:       input.Total,
		Status:      input.Status,
		CurrentUnit: input.CurrentUnit,
	}

	// insert order into the database
	orderResult := or.db.Create(&order)
	if orderResult.Error != nil || orderResult.RowsAffected == 0 {
		return nil, orderResult.Error
	}

	var orderItems []models.OrderItem
	for _, item := range input.OrderItem {
		// creating each orderItem from input parameters
		orderItem := models.OrderItem{
			OrderID:     order.ID,
			Price:       item.Price,
			Quantity:    item.Quantity,
			Description: item.Description,
		}

		// insert the orderItem into the database
		orderItemResult := or.db.Create(&orderItem)
		if orderItemResult.Error != nil || orderItemResult.RowsAffected == 0 {
			return nil, orderItemResult.Error
		}

		orderItems = append(orderItems, orderItem)
	}

	// assign the orderItems slice to the OrderItem field of the order
	order.OrderItem = orderItems

	return &order, nil
}
