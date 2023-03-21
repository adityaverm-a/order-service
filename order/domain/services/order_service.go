package services

//go:generate mockgen -destination=mocks/mock_order_service.go -package=mocks droomlabs/api/Cc-Api-Go/order/services OrderService

import (
	"demo/oms/order/data/models"
	"demo/oms/order/domain/entities"
	"demo/oms/order/domain/repositories"
)

// OrderService is...
type OrderService interface {
	GetOrderByID(id int64) (*models.Order, error)
	CreateOrder(input entities.CreateOrderInput) (*models.Order, error)
	UpdateOrderStatus(input entities.UpdateOrderStatusInput) (*models.Order, error)
}

// The NewOrderService function is a factory function that returns a new instance of the orderService
func NewOrderService(repo repositories.OrderRepository) OrderService {
	return &orderService{repo: repo}
}

type orderService struct {
	repo repositories.OrderRepository
}

// The GetOrderByID method of the orderService struct utilizes the OrderRepository instance, by calling the GetByID method on it, and returns any order and errors if received.
func (service *orderService) GetOrderByID(id int64) (*models.Order, error) {
	return service.repo.GetByID(id)
}

// The CreateOrder method of the orderService struct utilizes the OrderRepository instance, by calling the Create method on it, and returns the order created and errors if received.
func (service *orderService) CreateOrder(input entities.CreateOrderInput) (*models.Order, error) {
	return service.repo.Create(input)
}

// The UpdateOrderStatus method of the orderService struct utilizes the OrderRepository instance, by calling the Update method on it, and returns the order with updated status and errors if received.
func (service *orderService) UpdateOrderStatus(input entities.UpdateOrderStatusInput) (*models.Order, error) {
	return service.repo.Update(input)
}
