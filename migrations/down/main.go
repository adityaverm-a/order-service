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

	if !db.Migrator().HasTable(&models.Order{}) && !db.Migrator().HasTable(&models.OrderItem{}) {
		log.Println("`orders` & `order_items` doesn't exists, no need to run migration.")
		return
	}

	db.Migrator().DropTable(&models.Order{})
	db.Migrator().DropTable(&models.OrderItem{})

	log.Println("Migration ran succesfully, `orders` & `order_items` are removed.")
}
