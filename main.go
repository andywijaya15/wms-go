package main

import (
	"fmt"
	"log"
	"os"
	"wms-go/cmd"
	"wms-go/config"
	"wms-go/models"
	"wms-go/routes"
)

func main() {
	isCommand := len(os.Args) > 1 && os.Args[1] != ""
	if isCommand {
		switch os.Args[1] {
		case "sync-auto-pr":
			cmd.SyncAutoPr()
		case "sync-auto-pr1":
			fmt.Println(2)
		}
	} else {
		config.LoadEnv()
		models.ConnectDatabase()
		router := routes.SetupRouter()

		if err := router.Run(":8080"); err != nil {
			log.Fatal("Error starting the server: ", err)
		}
	}
}
