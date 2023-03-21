package models

import "demo/oms/order/data/constants"

type Order struct {
	ID          int64       `json:"id" gorm:"primaryKey;column:id" `
	Status      string      `json:"status" gorm:"column:status" `
	Total       float32     `json:"total" gorm:"column:total" `
	CurrentUnit string      `json:"current_unit" gorm:"column:current_unit" `
	OrderItem   []OrderItem `json:"items" gorm:"foreignKey:OrderID; references:ID"`
}

func (Order) TableName() string {
	return constants.TABLE_NAME_ORDERS
}
