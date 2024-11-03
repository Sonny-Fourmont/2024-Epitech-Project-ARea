package controllers

import (
	"area/config"
	"area/middlewares"
	"area/models"
	"area/storage"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetServices(c *gin.Context) (string, int) {
	var statusCode int = http.StatusOK

	var services models.ServiceAvailable = config.AllServices

	jsonResponseBytes, _ := json.Marshal(services)
	return string(jsonResponseBytes), statusCode
}

func HaveService(c *gin.Context) (string, int) {
	var statusCode int = http.StatusOK
	var serviceName string = c.Param("service")
	var token primitive.ObjectID = middlewares.GetClient(c)
	var tempToken models.Token = models.Token{UserID: token, Type: serviceName}

	var tokenExist bool = storage.ExistToken(tempToken)
	if !tokenExist {
		statusCode = http.StatusConflict
		jsonResponseBytes, _ := json.Marshal(map[string]string{"error": "Token doesnt exists"})
		return string(jsonResponseBytes), statusCode
	}

	return "", statusCode
}
