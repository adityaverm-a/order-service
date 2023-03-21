package controllers

import (
	"demo/oms/app/controller"
	"demo/oms/order/domain/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// OrderController is...
type OrderController interface {
	// GetOrders(gCtx *gin.Context)
	GetOrder(gCtx *gin.Context)
	// CreateOrder(gCtx *gin.Context)
	// UpdateOrder(gCtx *gin.Context)
}

// NewOrderController is...
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

type URI struct {
	OrderID int64 `json:"order_id" uri:"id"`
}

// ResetCart removes all items from cart and resets all the fields and returns updated cart
func (c *orderController) GetOrder(gCtx *gin.Context) {
	// cart, err := c.cartService.ResetCart(gCtx.Request.Context())
	// if err != nil {
	// 	c.SendWithError(gCtx, err)
	// 	return
	// }

	// cartModel, err := mappers.NewCartMapper().ToModel(*cart)
	// if err != nil {
	// 	c.SendWithError(gCtx, err)
	// 	return
	// }

	// var orderID = gCtx.Param("id")
	uri := URI{}

	if err := gCtx.BindUri(&uri); err != nil {
		gCtx.JSON(http.StatusBadRequest, gin.H{"code": "failed", "msg": err.Error()})
		return
	}

	order, _ := c.orderService.GetOrderByID(uri.OrderID)

	c.Send(gCtx, order)

}

// AddToCart adds an item in cart based on inputs and returns updated cart or error
// func (c *orderController) CreateOrder(gCtx *gin.Context) {
// 	var input models.AddItemInput
// 	if err := gCtx.ShouldBindQuery(&input); err != nil {
// 		gCtx.JSON(http.StatusBadRequest, gin.H{"code": "failed", "msg": err.Error()})
// 		return
// 	}

// 	if err := gCtx.ShouldBind(&input); err != nil {
// 		gCtx.JSON(http.StatusBadRequest, gin.H{"code": "failed", "msg": err.Error()})
// 		return
// 	}

// 	addItemInput, err := mappers.NewAddItemInputMapper().ToEntity(input)
// 	if err != nil {
// 		c.SendWithError(gCtx, err)
// 	}

// 	cart, err := c.cartService.AddItem(gCtx.Request.Context(), *addItemInput)
// 	if err != nil {
// 		c.SendWithError(gCtx, err)
// 		return
// 	}

// 	cartModel, err := mappers.NewCartMapper().ToModel(*cart)
// 	if err != nil {
// 		c.SendWithError(gCtx, err)
// 		return
// 	}

// 	c.Send(gCtx, cartModel)
// }

// AddMultipleItemsToCart adds multiple items in cart based on inputs and returns updated cart or error
// func (c *orderController) UpdateOrder(gCtx *gin.Context) {
// 	var input models.AddMultipleItemsInput
// 	if err := gCtx.ShouldBindQuery(&input); err != nil {
// 		c.SendWithError(gCtx, err)
// 		return
// 	}

// 	if err := gCtx.ShouldBind(&input); err != nil {
// 		c.SendWithError(gCtx, err)
// 		return
// 	}

// 	addMultipleItemsInput, err := mappers.NewAddMultipleItemsInputMapper().ToEntity(input)
// 	if err != nil {
// 		c.SendWithError(gCtx, err)
// 	}

// 	cart, err := c.cartService.AddMultipleItems(gCtx.Request.Context(), *addMultipleItemsInput)
// 	if err != nil {
// 		c.SendWithError(gCtx, err)
// 		return
// 	}

// 	cartModel, err := mappers.NewCartMapper().ToModel(*cart)
// 	if err != nil {
// 		c.SendWithError(gCtx, err)
// 		return
// 	}

// 	c.Send(gCtx, cartModel)
// }
