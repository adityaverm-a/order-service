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

	if db.Migrator().HasTable(&models.Order{}) && db.Migrator().HasTable(&models.OrderItem{}) {
		log.Println("`orders` & `order_items` already exists, no need to run migration.")
		return
	}

	db.Set("gorm:table_options", "ENGINE=InnoDB").Migrator().CreateTable(&models.Order{})
	db.Set("gorm:table_options", "ENGINE=InnoDB").Migrator().CreateTable(&models.OrderItem{})

	log.Println("Migration ran succesfully, `orders` & `order_items` are created.")
}
