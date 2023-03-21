package controllers

import (
	"demo/oms/app/controller"
	"demo/oms/order/domain/entities"
	"demo/oms/order/domain/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// OrderController is...
type OrderController interface {
	// GetOrders(gCtx *gin.Context)
	GetOrder(gCtx *gin.Context)
	CreateOrder(gCtx *gin.Context)
	UpdateOrderStatus(gCtx *gin.Context)
}

// NewOrderController creates a new instance of the orderController
func NewOrderController(orderService services.OrderService) OrderController {
	return &orderController{orderService: orderService}
}

type orderController struct {
	controller.Controller
	orderService services.OrderService
}

// GetCart gets the cart of a specific user (logged in or logged out)
// func (c *orderController) GetOrders(gCtx *gin.Context) {
// var input services.GetCartInput
// if err := gCtx.ShouldBindQuery(&input); err != nil {
// 	gCtx.JSON(http.StatusBadRequest, gin.H{"code": "failed", "msg": err.Error()})
// 	return
// }
// cart, err := c.cartService.GetCart(gCtx.Request.Context(), &input)
// if err != nil {
// 	c.SendWithError(gCtx, err)
// 	return
// }

// cartModel, err := mappers.NewCartMapper().ToModel(*cart)
// if err != nil {
// 	c.SendWithError(gCtx, err)
// 	return
// }

// c.Send(gCtx, order)
// }

// GetOrder handles incoming requests to retrieve an order by its ID.
func (c *orderController) GetOrder(gCtx *gin.Context) {
	input := entities.GetOrderByIDInput{}

	if err := gCtx.BindUri(&input); err != nil {
		gCtx.JSON(http.StatusBadRequest, gin.H{"code": "failed", "msg": err.Error()})
		return
	}

	order, err := c.orderService.GetOrderByID(input.OrderID)
	if err != nil {
		c.SendWithError(gCtx, err)
		return
	}

	c.Send(gCtx, order)
}

// AddToCart adds an item in cart based on inputs and returns updated order or error
func (c *orderController) CreateOrder(gCtx *gin.Context) {
	var input entities.CreateOrderInput
	if err := gCtx.ShouldBindQuery(&input); err != nil {
		gCtx.JSON(http.StatusBadRequest, gin.H{"code": "failed", "msg": err.Error()})
		return
	}

	if err := gCtx.ShouldBind(&input); err != nil {
		gCtx.JSON(http.StatusBadRequest, gin.H{"code": "failed", "msg": err.Error()})
		return
	}

	order, err := c.orderService.CreateOrder(input)
	if err != nil {
		c.SendWithError(gCtx, err)
		return
	}

	c.Send(gCtx, order)
}

// UpdateOrderStatus updates an order's status based on inputs and returns updated order or error
func (c *orderController) UpdateOrderStatus(gCtx *gin.Context) {
	var input entities.UpdateOrderStatusInput
	if err := gCtx.ShouldBindQuery(&input); err != nil {
		c.SendWithError(gCtx, err)
		return
	}

	if err := gCtx.ShouldBind(&input); err != nil {
		c.SendWithError(gCtx, err)
		return
	}

	order, err := c.orderService.UpdateOrderStatus(input)
	if err != nil {
		c.SendWithError(gCtx, err)
		return
	}

	c.Send(gCtx, order)
}
