package entities

type GetOrderByIDInput struct {
	OrderID int64 `json:"order_id" uri:"id"`
}

type CreateOrderInput struct {
	Status       string                 `form:"status" json:"status"`
	Total        float32                `form:"total" json:"total"`
	OrderItem    []CreateOrderItemInput `form:"items" json:"items"`
	CurrencyUnit string                 `form:"currency_unit" json:"currency_unit"`
}

type CreateOrderItemInput struct {
	Price       float32 `form:"price" json:"price"`
	Quantity    int64   `form:"quantity" json:"quantity"`
	Description string  `form:"description" json:"description"`
}

type UpdateOrderStatusInput struct {
	OrderID int64  `form:"id" json:"id"`
	Status  string `form:"status" json:"status"`
}

type OrderFiltersInput struct {
	OrderID      int64   `form:"id" json:"id"`
	CurrencyUnit string  `form:"currency_unit" json:"currency_unit"`
	Status       string  `form:"status" json:"status"`
	Total        float32 `form:"total" json:"total"`
	Limit        int     `form:"limit" json:"limit"`
	Offset       int     `form:"offset" json:"offset"`
}
