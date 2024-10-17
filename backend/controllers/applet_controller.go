package controllers

import (
	"area/middlewares"
	"area/models"
	"area/storage"
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Require Token Middleware
func GetApplets(c *gin.Context) (string, int) {
	var clientID string = middlewares.GetClient(c).Hex()
	var applets []models.Applet = storage.GetApplets(clientID)

	response := struct {
		Applets []models.Applet `json:"applet_array"`
	}{Applets: applets}
	jsonResponseBytes, _ := json.Marshal(response)
	return string(jsonResponseBytes), http.StatusOK
}

// Require Token Middleware
func AddApplet(c *gin.Context) (string, int) {
	var applet models.Applet
	if err := c.ShouldBindJSON(&applet); err != nil {
		jsonResponseBytes, _ := json.Marshal(map[string]string{"error": "Invalid JSON"})
		return string(jsonResponseBytes), http.StatusBadRequest
	}

	collection := storage.DB.Collection("applets")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	applet.ID_User = middlewares.GetClient(c)
	_, err := collection.InsertOne(ctx, applet)
	if err != nil {
		jsonResponseBytes, _ := json.Marshal(map[string]string{"error": "Failed to create applet"})
		return string(jsonResponseBytes), http.StatusInternalServerError
	}
	jsonResponseBytes, _ := json.Marshal(map[string]string{"message": "Applet added successfully"})
	return string(jsonResponseBytes), http.StatusOK
}
