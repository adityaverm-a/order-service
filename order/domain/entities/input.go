package entities

type GetOrderByIDInput struct {
	OrderID int64 `json:"order_id" uri:"id" binding:"required"`
}

type CreateOrderInput struct {
	Status       string                 `form:"status" json:"status" binding:"required"`
	Total        float32                `form:"total" json:"total" binding:"required"`
	OrderItem    []CreateOrderItemInput `form:"items" json:"items" binding:"required"`
	CurrencyUnit string                 `form:"currency_unit" json:"currency_unit" binding:"required"`
}

type CreateOrderItemInput struct {
	Price       float32 `form:"price" json:"price" binding:"required"`
	Quantity    int64   `form:"quantity" json:"quantity" binding:"required"`
	Description string  `form:"description" json:"description" binding:"required"`
}

type UpdateOrderStatusInput struct {
	OrderID int64  `form:"id" json:"id" binding:"required"`
	Status  string `form:"status" json:"status" binding:"required"`
}

type OrderFiltersInput struct {
	OrderID      int64   `form:"id" json:"id"`
	CurrencyUnit string  `form:"currency_unit" json:"currency_unit"`
	Status       string  `form:"status" json:"status"`
	Total        float32 `form:"total" json:"total"`
	Limit        int     `form:"limit" json:"limit"`
	Offset       int     `form:"offset" json:"offset"`
}
