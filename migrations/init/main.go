package main

import (
	"demo/oms/config"
	"demo/oms/datasources"
	"demo/oms/order/data/models"
	"log"
)

func main() {
	config.LoadConfig(".")

	ds, err := datasources.Get()
	if err != nil {
		log.Println(err.Error())
		return
	}

	db := ds.GORMClient

	var users = []models.Order{
		{
			Status:       "PENDING",
			Total:        54990.82,
			CurrencyUnit: "INR",
			OrderItem: []models.OrderItem{
				{
					Description: "PS 5",
					Quantity:    1,
					Price:       54990.82,
				},
			},
		},
		{
			Status:       "SHIPPED",
			Total:        590.2,
			CurrencyUnit: "USD",
			OrderItem: []models.OrderItem{
				{
					Description: "Macbook Pro",
					Quantity:    1,
					Price:       590.2,
				},
			},
		},
		{
			Status:       "DELIVERED",
			Total:        500.2,
			CurrencyUnit: "USD",
			OrderItem: []models.OrderItem{
				{
					Description: "Macbook Air",
					Quantity:    2,
					Price:       250.1,
				},
			},
		},
	}

	err = db.Create(&users).Error
	if err != nil {
		log.Println(err.Error())
		return
	}
}
