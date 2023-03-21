package entities

type GetOrderByIDInput struct {
	OrderID int64 `json:"order_id" uri:"id"`
}

type CreateOrderInput struct {
	Status      string                 `json:"status"`
	Total       float32                `json:"total"`
	OrderItem   []CreateOrderItemInput `json:"items"`
	CurrentUnit string                 `json:"currencyUnit"`
}

type CreateOrderItemInput struct {
	Price       float32 `json:"price"`
	Quantity    int64   `json:"quantity"`
	Description string  `json:"description"`
}

type UpdateOrderStatusInput struct {
	OrderID int64  `json:"id"`
	Status  string `json:"status"`
}
