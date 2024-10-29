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
	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/option"
)

func GoogleLogin(c *gin.Context) (string, int) {
	if config.GoogleOauth == nil {
		jsonResponseBytes, _ := json.Marshal(map[string]string{"error": "OAuth configuration is not initialized"})
		return string(jsonResponseBytes), http.StatusInternalServerError
	}
	url := config.GoogleOauth.AuthCodeURL("state-token", oauth2.AccessTypeOffline, oauth2.SetAuthURLParam("prompt", "consent"))
	return url, http.StatusPermanentRedirect
}

func GoogleLoggedIn(c *gin.Context) (primitive.ObjectID, string, int) {
	var user models.User
	var token models.Token

	httpClient := config.GoogleOauth.Client(context.Background(), config.GoogleToken)
	gmail, _ := gmail.NewService(context.Background(), option.WithHTTPClient(httpClient))
	googleUser, _ := gmail.Users.GetProfile("me").Do()

	user.ID = primitive.NewObjectID()
	user.Username = googleUser.EmailAddress
	user.Email = googleUser.EmailAddress
	hashedPassword, _ := utils.GenerateHash("googleAccount")
	user.Password = hashedPassword
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	if !storage.CreateORUpdateUser(user) {
		jsonResponseBytes, _ := json.Marshal(map[string]string{"error": "Failed to create user"})
		return primitive.NilObjectID, string(jsonResponseBytes), http.StatusInternalServerError
	}
	userFromDB, found := storage.GetUserByEmail(user.Email)
	token.UserID = user.ID
	if found {
		token.UserID = userFromDB.ID
	}
	token.Type = "Google"
	token.TokenData = config.GoogleToken
	token.CreatedAt = time.Now()
	token.UpdatedAt = time.Now()

	if !storage.CreateORUpdateToken(token) {
		jsonResponseBytes, _ := json.Marshal(map[string]string{"error": "Failed to create user"})
		return primitive.NilObjectID, string(jsonResponseBytes), http.StatusInternalServerError
	}
	log.Output(0, "Refresh token has been created!")
	return user.ID, "", http.StatusOK
}
