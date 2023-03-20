package app

import (
	"demo/oms/app/router"
	"demo/oms/app/server"
	"demo/oms/config"

	"fmt"
)

const defaultPort = ":5000"

// InitApp, this function initializes the app
func InitApp() {
	config.LoadConfig(".")

	// _, err := datasources.Get()

	// if err != nil {
	// 	log.Println(err.Error())
	// 	return
	// }

	// Creates a router
	router := router.Create()

	// Checks if there is any port specified in the configuration file
	port := config.Config.Port
	if port == "" {
		port = defaultPort
		fmt.Printf("Connected to default port: %s", port)
	}

	// Starts the server
	server.Serve(port, router)
}
