package entities

type GetOrderByIDInput struct {
	OrderID int64 `json:"order_id" uri:"id"`
}
