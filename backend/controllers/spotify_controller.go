package controllers

import (
	"area/models"
	"area/storage"
	"area/utils"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/zmb3/spotify"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/oauth2"
)

func SpotifyLogin(c *gin.Context) (string, int) {
	utils.SpotifyAuth()
	if utils.SpotifyOauth == nil {
		jsonResponseBytes, _ := json.Marshal(map[string]string{"error": "OAuth configuration is not initialized"})
		return string(jsonResponseBytes), http.StatusInternalServerError
	}
	url := utils.SpotifyOauth.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	return url, http.StatusPermanentRedirect
}

func SpotifyLoggedIn(c *gin.Context) (primitive.ObjectID, string, int) {
	var user models.User
	var token models.Token

	client := spotify.Authenticator{}.NewClient(utils.SpotifyToken)
	userInfo, err := client.CurrentUser()
	if err != nil {
		jsonResponseBytes, _ := json.Marshal(map[string]string{"error": "Failed to create user"})
		return primitive.NilObjectID, string(jsonResponseBytes), http.StatusInternalServerError
	}

	user.ID = primitive.NewObjectID()
	user.Username = userInfo.DisplayName
	user.Email = userInfo.Email
	hashedPassword, _ := utils.GenerateHash("spotifyAccount")
	user.Password = hashedPassword
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	userFromDB, found := storage.GetUserByEmail(user.Email)
	token.UserID = user.ID
	if found {
		token.UserID = userFromDB.ID
	}
	token.Type = "Spotify"
	token.TokenData = utils.SpotifyToken
	token.CreatedAt = time.Now()
	token.UpdatedAt = time.Now()
	if !storage.CreateORUpdateUser(user) {
		jsonResponseBytes, _ := json.Marshal(map[string]string{"error": "Failed to create user"})
		return primitive.NilObjectID, string(jsonResponseBytes), http.StatusInternalServerError
	}
	if !storage.CreateORUpdateToken(token) {
		jsonResponseBytes, _ := json.Marshal(map[string]string{"error": "Failed to create token"})
		return primitive.NilObjectID, string(jsonResponseBytes), http.StatusInternalServerError
	}
	return user.ID, "", http.StatusOK
}
