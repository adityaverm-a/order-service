package router

import (
	"demo/oms/order/containers"
	"demo/oms/order/controllers"

	"github.com/gin-gonic/gin"
)

var orderController controllers.OrderController

// InjectOrderRoutes is...
func InjectOrderRoutes(router *gin.RouterGroup) {
	setupRoutes(router)
}

func setupRoutes(r *gin.RouterGroup) {

	container := containers.NewContainer()
	orderController = container.InjectOrderController()

	// r.GET("/orders", func(c *gin.Context) {
	// 	orderController.GetOrders(c)
	// })

	// r.POST("/orders", func(c *gin.Context) {
	// 	orderController.CreateOrder(c)
	// })

	r.GET("/orders/:id", func(c *gin.Context) {
		orderController.GetOrder(c)
	})

	// r.PATCH("/orders/:id", func(c *gin.Context) {
	// 	orderController.UpdateOrder(c)
	// })
}
