package controllers

import (
	"area/config"
	"area/models"
	"area/storage"
	"area/utils"
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/go-github/github"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/oauth2"
)

func GithubLogin(c *gin.Context) (string, int) {
	if config.GithubOauth == nil {
		jsonResponseBytes, _ := json.Marshal(map[string]string{"error": "OAuth configuration is not initialized"})
		return string(jsonResponseBytes), http.StatusInternalServerError
	}
	url := config.GithubOauth.AuthCodeURL("state-token", oauth2.AccessTypeOffline, oauth2.SetAuthURLParam("prompt", "consent"))
	return url, http.StatusPermanentRedirect
}

func GithubLoggedIn(c *gin.Context) (primitive.ObjectID, string, int) {
	var user models.User
	var token models.Token
	httpClient := config.GithubOauth.Client(context.Background(), config.GithubToken)
	githubClient := github.NewClient(httpClient)
	userInfo, _, err := githubClient.Users.Get(c, "")
	if err != nil {
		jsonResponseBytes, _ := json.Marshal(map[string]string{"error": "Failed to create user"})
		return primitive.NilObjectID, string(jsonResponseBytes), http.StatusInternalServerError
	}

	user.ID = primitive.NewObjectID()
	user.Username = *userInfo.Name
	user.Email = ""
	hashedPassword, _ := utils.GenerateHash("githubAccount")
	user.Password = hashedPassword
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	if !storage.CreateORUpdateUser(user) {
		jsonResponseBytes, _ := json.Marshal(map[string]string{"error": "Failed to create user"})
		return primitive.NilObjectID, string(jsonResponseBytes), http.StatusInternalServerError
	}

	userFromDB, found := storage.GetUserByEmail(user.Email)
	if found {
		user.ID = userFromDB.ID
	}
	token.UserID = user.ID
	token.Type = "Github"
	token.TokenData = config.GithubToken
	token.CreatedAt = time.Now()
	token.UpdatedAt = time.Now()
	if err != nil {
		return primitive.NilObjectID, err.Error(), http.StatusInternalServerError
	}

	if !storage.CreateORUpdateToken(token) {
		jsonResponseBytes, _ := json.Marshal(map[string]string{"error": "Failed to create user"})
		return primitive.NilObjectID, string(jsonResponseBytes), http.StatusInternalServerError
	}
	return user.ID, "", http.StatusOK
}
