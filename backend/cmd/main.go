package main

import (
	"log"
	"time"

	"area/config"
	"area/routes"
	"area/services"
	storage "area/storage"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadConfig()
	storage.ConnectDatabase()

	router := gin.Default()
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowMethods = []string{"POST", "GET", "PUT", "DELETE"}
	corsConfig.AllowHeaders = []string{"Origin", "Content-Type", "Authorization", "Accept", "User-Agent"}
	corsConfig.AllowCredentials = true
	corsConfig.MaxAge = 12 * time.Hour

	router.Use(cors.New(corsConfig))

	routes.InitRoutes(router)

	port := config.ConfigGin.Port
	log.Printf("Starting server on port %s...", port)
	go services.LaunchServices()

	router.Run(":" + port)
}
