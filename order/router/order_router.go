package router

import (
	"demo/oms/order/containers"
	"demo/oms/order/controllers"

	"github.com/gin-gonic/gin"
)

var orderController controllers.OrderController

// InjectOrderRoutes is defined to set up the routes and inject the dependencies for the router to work correctly.
func InjectOrderRoutes(router *gin.RouterGroup) {
	setupRoutes(router)
}

// setupRoutes is defined to set up the router's endpoints.
func setupRoutes(r *gin.RouterGroup) {

	// create a new instance of the container
	container := containers.NewContainer()
	orderController = container.InjectOrderController()

	// a GET request to /orders will fetch  all orders
	// r.GET("/orders", func(c *gin.Context) {
	// 	orderController.GetOrder(c)
	// })

	// a POST request to /order will create an order
	r.POST("/order", func(c *gin.Context) {
		orderController.CreateOrder(c)
	})

	// a GET request to /order/:id will fetch an order
	r.GET("/order/:id", func(c *gin.Context) {
		orderController.GetOrder(c)
	})

	// a POST request to /order/change-status/:id will update status for that order
	// r.POST("/order/change-status/:id", func(c *gin.Context) {
	// 	orderController.GetOrder(c)
	// })
}
