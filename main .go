package main

import (
	"log"
	"wms-go/config"
	"wms-go/models"
	"wms-go/routes"
)

func main() {
	config.LoadEnv()
	models.ConnectDatabase()
	router := routes.SetupRouter()

	if err := router.Run(":8080"); err != nil {
		log.Fatal("Error starting the server: ", err)
	}
}
