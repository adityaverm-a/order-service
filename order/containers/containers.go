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

// NewContainer is...
func NewContainer() Container {
	return &container{}
}

type container struct{}

func (c *container) InjectOrderController() controllers.OrderController {
	orderService, err := InjectOrderService()
	if err != nil {
		panic(1)
	}

	return controllers.NewOrderController(orderService)
}

func InjectOrderService() (services.OrderService, error) {
	dataSources, err := datasources.Get()

	if err != nil {
		panic(1)
	}

	sqlClient := dataSources.SQLXClient
	orderRepository := repositories.NewOrderRepository(sqlClient)

	orderService := services.NewOrderService(orderRepository)

	return orderService, nil
}
