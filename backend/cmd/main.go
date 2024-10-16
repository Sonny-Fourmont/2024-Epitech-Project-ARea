package main

import (
	"log"

	"area/config"
	"area/routes"
	"area/services"
	storage "area/storage"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadConfig()
	storage.ConnectDatabase()

	router := gin.Default()

	routes.InitRoutes(router)

	port := config.ConfigGin.Port
	log.Printf("Starting server on port %s...", port)
	go services.LaunchServices()

	router.Run(":" + port)
}
