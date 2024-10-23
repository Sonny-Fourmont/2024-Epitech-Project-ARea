package controllers

import (
	"area/middlewares"
	"area/models"
	"area/storage"
	"encoding/json"
	"net/http"

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
	applet.ID_User = middlewares.GetClient(c)
	var status bool = storage.CreateApplet(applet)
	if !status {
		jsonResponseBytes, _ := json.Marshal(map[string]string{"error": "Failed to create applet"})
		return string(jsonResponseBytes), http.StatusInternalServerError
	}
	jsonResponseBytes, _ := json.Marshal(map[string]string{"message": "Applet added successfully"})
	return string(jsonResponseBytes), http.StatusOK
}

// Require Token Middleware
func UpdateApplet(c *gin.Context) (string, int) {
	var applet models.Applet
	if err := c.ShouldBindJSON(&applet); err != nil {
		jsonResponseBytes, _ := json.Marshal(map[string]string{"error": "Invalid JSON"})
		return string(jsonResponseBytes), http.StatusBadRequest
	}
	applet.ID_User = middlewares.GetClient(c)
	var status bool = storage.UpdateApplet(applet)
	if !status {
		jsonResponseBytes, _ := json.Marshal(map[string]string{"error": "Failed to update applet"})
		return string(jsonResponseBytes), http.StatusInternalServerError
	}
	jsonResponseBytes, _ := json.Marshal(map[string]string{"message": "Applet added successfully"})
	return string(jsonResponseBytes), http.StatusOK
}
