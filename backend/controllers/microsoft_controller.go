package controllers

import (
	"area/config"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/oauth2"
)

func AzureLogin(c *gin.Context) (string, int) {
	config.AzureAuth()
	if config.AzureOauth == nil {
		jsonResponseBytes, _ := json.Marshal(map[string]string{"error": "OAuth configuration is not initialized"})
		return string(jsonResponseBytes), http.StatusInternalServerError
	}
	url := config.AzureOauth.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	return url, http.StatusPermanentRedirect
}

func AzureLoggedIn(c *gin.Context) (primitive.ObjectID, string, int) {
	// var user models.User
	// var token models.Token

	// httpClient := config.AzureOauth.Client(context.Background(), config.AzureToken)

	return primitive.NewObjectID(), "", 0
}
