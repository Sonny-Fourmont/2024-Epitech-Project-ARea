package controllers

import (
	"area/config"
	"area/middlewares"
	"area/models"
	"area/storage"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

func SpotifyLogin(c *gin.Context) (string, int) {
	if config.SpotifyOauth == nil {
		jsonResponseBytes, _ := json.Marshal(map[string]string{"error": "OAuth configuration is not initialized"})
		return string(jsonResponseBytes), http.StatusInternalServerError
	}
	url := config.SpotifyOauth.AuthCodeURL("state-token", oauth2.AccessTypeOffline, oauth2.SetAuthURLParam("prompt", "consent"))
	return url, http.StatusPermanentRedirect
}

func SpotifyLoggedIn(c *gin.Context) (string, int) {
	var token models.Token
	token.UserID = middlewares.GetClient(c)
	token.Type = "Spotify"
	token.TokenData = config.SpotifyToken
	token.CreatedAt = time.Now()
	token.UpdatedAt = time.Now()
	if !storage.CreateORUpdateToken(token) {
		jsonResponseBytes, _ := json.Marshal(map[string]string{"error": "Failed to create token"})
		return string(jsonResponseBytes), http.StatusInternalServerError
	}

	var statusCode int
	jsonResponseBytes, _ := json.Marshal(map[string]string{"message": "Spotify login successfully"})
	return string(jsonResponseBytes), statusCode
}
