package main

import (
	"log"
	"os"

	api "area/api"
	storage "area/storage"

	"github.com/gin-gonic/gin"
)

func main() {
	storage.ConnectDatabase()

	router := gin.Default()

	api.InitRoutes(router)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Starting server on port %s...", port)
	router.Run(":" + port)
}
