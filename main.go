package main

import (
	"log"
	"os"
	"wms-go/cmd"
	"wms-go/config"
	"wms-go/models"
	"wms-go/routes"
)

const (
	SyncAutoPrCommand = "sync-auto-pr"
)

var DB = models.DB

func handleCommand(command string) {
	switch command {
	case SyncAutoPrCommand:
		cmd.SyncAutoPr()
	default:
		log.Fatalf("Unknown command: %v", command)
	}
}

func startServer() {
	appPort := os.Getenv("APP_PORT")
	router := routes.SetupRouter()

	if err := router.Run(":" + appPort); err != nil {
		log.Fatal("Error starting the server: ", err)
	}
}

func main() {
	config.LoadEnv()
	models.ConnectDatabase()
	isCommand := len(os.Args) > 1 && os.Args[1] != ""
	if isCommand {
		handleCommand(os.Args[1])
	} else {
		startServer()
	}
}
