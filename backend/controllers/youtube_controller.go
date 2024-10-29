package controllers

import (
	"area/middlewares"
	"area/models"
	"area/services"
	"area/storage"
	"area/utils"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/oauth2"
)

// Require Token Middleware
func YoutubeLogin(c *gin.Context) (string, int) {
	utils.YoutubeLikedAuth()
	if utils.YoutubeOauth == nil {
		jsonResponseBytes, _ := json.Marshal(map[string]string{"error": "OAuth configuration is not initialized"})
		return string(jsonResponseBytes), http.StatusInternalServerError
	}
	url := utils.YoutubeOauth.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	return url, http.StatusPermanentRedirect
}

// Require Token Middleware
func YoutubeLoggedIn(c *gin.Context) (string, int) {
	var token models.Token

	token.ID = primitive.NewObjectID()
	token.UserID = middlewares.GetClient(c)
	token.TokenData = utils.YoutubeToken
	token.Type = "Youtube_liked"
	token.CreatedAt = time.Now()
	token.UpdatedAt = time.Now()

	if !storage.CreateORUpdateToken(token) {
		jsonResponseBytes, _ := json.Marshal(map[string]string{"error": "Failed to create user"})
		return string(jsonResponseBytes), http.StatusInternalServerError
	}

	var videoLikedJSON []string
	var statusCode int
	videoLikedJSON, statusCode = services.GetLastedLiked(token.TokenData)
	jsonResponseBytes, _ := json.Marshal(videoLikedJSON)
	return string(jsonResponseBytes), statusCode
}
