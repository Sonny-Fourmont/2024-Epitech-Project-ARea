package controllers

import (
	"area/config"
	"area/models"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetServices(c *gin.Context) (string, int) {
	var statusCode int = http.StatusOK

	var services models.ServiceAvailable = config.AllServices

	jsonResponseBytes, _ := json.Marshal(services)
	return string(jsonResponseBytes), statusCode
}
