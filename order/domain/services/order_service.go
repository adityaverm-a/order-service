package services

//go:generate mockgen -destination=mocks/mock_order_service.go -package=mocks droomlabs/api/Cc-Api-Go/order/services OrderService

import (
	"demo/oms/order/data/models"
	"demo/oms/order/domain/repositories"
)

type OrderService interface {
	GetOrderByID(id int64) (*models.Persons, error)
}

func NewOrderService(repo repositories.OrderRepository) OrderService {
	return &orderService{repo: repo}
}

type orderService struct {
	repo repositories.OrderRepository
}

func (service *orderService) GetOrderByID(id int64) (*models.Persons, error) {

	order, err := service.repo.GetByID(id)
	return order, err
}
