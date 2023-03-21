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

// GetByFilters fetches orders with the provided filters.
func (or *orderRepo) GetByFilters(filters entities.OrderFiltersInput) (*[]models.Order, error) {
	var orders []models.Order

	db := or.db.Table(constants.TABLE_NAME_ORDERS)
	db = or.setFilters(filters, db)

	err := db.Preload("OrderItem").Limit(filters.Limit).Offset(filters.Offset).Order("id desc").Find(&orders).Error
	if err != nil {
		return nil, err
	}

	return &orders, nil
}

func (or *orderRepo) setFilters(filters entities.OrderFiltersInput, db *gorm.DB) *gorm.DB {
	if filters.OrderID != 0 {
		db = db.Where("orders.id = ?", filters.OrderID)
	}

	if filters.CurrencyUnit != "" {
		db = db.Where("orders.currency_unit = ?", filters.CurrencyUnit)
	}

	if filters.Total != 0 {
		db = db.Where("orders.total >?", filters.Total)
	}

	if filters.Status != "" {
		db = db.Where("orders.status = ?", filters.Status)
	}

	return db
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
		return nil, OrderNotFound
	}

	return &order, nil
}

// Create adds an order to the database.
func (or *orderRepo) Create(input entities.CreateOrderInput) (*models.Order, error) {
	// create a new order from input parameters
	order := models.Order{
		Total:        input.Total,
		Status:       input.Status,
		CurrencyUnit: input.CurrencyUnit,
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

// Update changes order statuses as of now, but can be extended to update an order!
func (or *orderRepo) Update(input entities.UpdateOrderStatusInput) (*models.Order, error) {
	result := or.db.Model(&models.Order{}).Where("id = ?", input.OrderID).Updates(models.Order{Status: input.Status})
	if result.Error != nil || result.RowsAffected == 0 {
		return nil, result.Error
	}

	order, err := or.GetByID(input.OrderID)
	if err != nil {
		return nil, err
	}

	return order, nil
}
