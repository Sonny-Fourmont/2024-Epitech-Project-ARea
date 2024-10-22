package controllers

import (
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
	utils.GithubAuth()
	if utils.GithubOauth == nil {
		jsonResponseBytes, _ := json.Marshal(map[string]string{"error": "OAuth configuration is not initialized"})
		return string(jsonResponseBytes), http.StatusInternalServerError
	}
	url := utils.GithubOauth.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	return url, http.StatusPermanentRedirect
}

func GithubLoggedIn(c *gin.Context) (primitive.ObjectID, string, int) {
	var user models.User
	var token models.Token
	httpClient := utils.GithubOauth.Client(context.Background(), utils.GithubToken)
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

	userFromDB, found := storage.GetUserByEmail(user.Email)
	token.UserID = user.ID
	if found {
		token.UserID = userFromDB.ID
	}
	token.Type = "Github"
	token.TokenData = utils.GithubToken
	token.CreatedAt = time.Now()
	token.UpdatedAt = time.Now()
	if !storage.CreateORUpdateUser(user) {
		jsonResponseBytes, _ := json.Marshal(map[string]string{"error": "Failed to create user"})
		return primitive.NilObjectID, string(jsonResponseBytes), http.StatusInternalServerError
	}

	if !storage.CreateORUpdateToken(token) {
		jsonResponseBytes, _ := json.Marshal(map[string]string{"error": "Failed to create user"})
		return primitive.NilObjectID, string(jsonResponseBytes), http.StatusInternalServerError
	}
	return user.ID, "", http.StatusOK
}
