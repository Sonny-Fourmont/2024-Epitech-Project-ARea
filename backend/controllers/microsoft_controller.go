package controllers

import (
	"area/config"
	"area/models"
	"area/storage"
	"area/utils"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

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
	var user models.User
	var token models.Token

	client := config.AzureOauth.Client(context.Background(), config.AzureToken)
	userInfos, err := client.Get("https://graph.microsoft.com/v1.0/me")
	if err != nil {
		return primitive.NilObjectID, err.Error(), http.StatusInternalServerError
	}
	defer userInfos.Body.Close()
	if userInfos.StatusCode != http.StatusOK {
		return primitive.NilObjectID, "failed to fetch user profile, status:", userInfos.StatusCode
	}
	// err = json.NewDecoder(userInfos.Body).Decode(&user)
	// if err != nil {
	// 	return primitive.NilObjectID, err.Error(), http.StatusInternalServerError
	// }

	if !storage.CreateORUpdateUser(user) {
		return primitive.NilObjectID, "Failed to create user", http.StatusInternalServerError
	}
	userFromDB, found := storage.GetUserByEmail(user.Email)
	if found {
		user.ID = userFromDB.ID
	}

	token.UserID = user.ID
	token.Type = "Microsoft"
	token.TokenData = config.AzureToken
	token.CreatedAt = time.Now()
	token.UpdatedAt = time.Now()
	token, err = utils.RefreshToken(token)
	if err != nil {
		return primitive.NilObjectID, err.Error(), http.StatusInternalServerError
	}

	if !storage.CreateORUpdateToken(token) {
		return primitive.NilObjectID, "Failed to create user", http.StatusInternalServerError
	}
	log.Output(0, "Refresh token has been created!")

	return user.ID, "User login successfully", http.StatusOK
}
