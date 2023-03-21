package entities

type GetOrderByIDInput struct {
	OrderID int64 `json:"order_id" uri:"id"`
}

type CreateOrderInput struct {
	Total       float32                `json:"total"`
	CurrentUnit string                 `json:"currencyUnit"`
	Status      string                 `json:"status"`
	OrderItem   []CreateOrderItemInput `json:"items"`
}

type CreateOrderItemInput struct {
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	Quantity    int64   `json:"quantity"`
}
