package containers

import (
	"demo/oms/datasources"
	"demo/oms/order/controllers"
	"demo/oms/order/data/repositories"
	"demo/oms/order/domain/services"
)

// Container is...
type Container interface {
	InjectOrderController() controllers.OrderController
}

// NewContainer creates a new Container instance
func NewContainer() Container {
	return &container{}
}

type container struct{}

// InjectOrderController injects an instance of the OrderController
func (c *container) InjectOrderController() controllers.OrderController {
	orderService, err := InjectOrderService()
	if err != nil {
		panic(1)
	}

	return controllers.NewOrderController(orderService)
}

// InjectOrderService injects an instance of the OrderService
func InjectOrderService() (services.OrderService, error) {

	sqlClient := datasources.GetSQLClient()
	orderRepository := repositories.NewOrderRepository(sqlClient)

	orderService := services.NewOrderService(orderRepository)

	return orderService, nil
}
