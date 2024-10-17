package controllers

import (
	"area/models"
	"area/services"
	"area/storage"
	"area/utils"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/go-github/github"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/oauth2"
	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/option"
)

func RegisterUser(c *gin.Context) (string, int) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		jsonResponseBytes, _ := json.Marshal(map[string]string{"error": "Invalid JSON"})
		return string(jsonResponseBytes), http.StatusBadRequest

	}

	hashedPassword, _ := utils.GenerateHash(user.Password)
	user.Password = hashedPassword

	collection := storage.DB.Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := collection.InsertOne(ctx, user)
	if err != nil {
		jsonResponseBytes, _ := json.Marshal(map[string]string{"error": "Failed to create user"})
		return string(jsonResponseBytes), http.StatusInternalServerError
	}
	jsonResponseBytes, _ := json.Marshal(map[string]string{"message": "User registered successfully"})
	return string(jsonResponseBytes), http.StatusOK
}

func GetUser(c *gin.Context) (string, int) {
	id := c.Param("id")

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		jsonResponseBytes, _ := json.Marshal(map[string]string{"error": "Invalid ID format"})
		return string(jsonResponseBytes), http.StatusBadRequest
	}

	collection := storage.DB.Collection("users")
	var user models.User

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&user)
	if err != nil {
		jsonResponseBytes, _ := json.Marshal(map[string]string{"error": "User not found"})
		return string(jsonResponseBytes), http.StatusBadRequest
	}

	jsonResponseBytes, _ := json.Marshal(map[string]string{
		"id":       user.ID.Hex(),
		"username": user.Username,
		"email":    user.Email,
	})
	return string(jsonResponseBytes), http.StatusOK

}

// ----- GOOGLE ----- //
func GoogleLogin(c *gin.Context) (string, int) {
	utils.GoogleAuth()
	if utils.GoogleOauth == nil {
		jsonResponseBytes, _ := json.Marshal(map[string]string{"error": "OAuth configuration is not initialized"})
		return string(jsonResponseBytes), http.StatusInternalServerError
	}
	url := utils.GoogleOauth.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	return url, http.StatusPermanentRedirect
}

func GoogleLoggedIn(c *gin.Context) (*oauth2.Token, string, int) {
	var user models.User
	var tokens models.Token

	httpClient := utils.GoogleOauth.Client(context.Background(), utils.GoogleToken)
	gmail, _ := gmail.NewService(context.Background(), option.WithHTTPClient(httpClient))
	googleUser, _ := gmail.Users.GetProfile("me").Do()

	user.ID = primitive.NewObjectID()
	user.Username = googleUser.EmailAddress
	user.Email = googleUser.EmailAddress
	hashedPassword, _ := utils.GenerateHash("googleAccount")
	user.Password = hashedPassword

	var status bool = storage.CreateUser(user)
	if !status {
		jsonResponseBytes, _ := json.Marshal(map[string]string{"error": "Failed to create user"})
		return nil, string(jsonResponseBytes), http.StatusInternalServerError
	}

	tokens.UserID = user.ID
	tokens.Type = "Google"
	tokens.TokenData = utils.GoogleToken

	status = storage.CreateToken(tokens)
	if !status {
		jsonResponseBytes, _ := json.Marshal(map[string]string{"error": "Failed to create user"})
		return nil, string(jsonResponseBytes), http.StatusInternalServerError
	}
	log.Output(0, "Refresh token has been created!")
	return utils.GoogleToken, "", http.StatusOK
}

// ----- YOUTUBE ----- //
func YoutubeLogin(c *gin.Context) (string, int) {
	utils.YoutubeLikedAuth()
	if utils.YoutubeOauth == nil {
		jsonResponseBytes, _ := json.Marshal(map[string]string{"error": "OAuth configuration is not initialized"})
		return string(jsonResponseBytes), http.StatusInternalServerError
	}
	url := utils.YoutubeOauth.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	return url, http.StatusPermanentRedirect
}

func YoutubeLoggedIn(c *gin.Context) (string, int) {
	var token models.Token

	httpClient := utils.YoutubeOauth.Client(context.Background(), utils.YoutubeToken)
	gmail, _ := gmail.NewService(context.Background(), option.WithHTTPClient(httpClient))
	youtubeUser, _ := gmail.Users.GetProfile("me").Do()

	token.ID = primitive.NewObjectID()
	token.UserID = storage.GetUserIDByEmail(youtubeUser.EmailAddress)
	token.TokenData = utils.YoutubeToken
	token.Type = "Youtube_liked"

	var status bool = storage.CreateToken(token)
	if !status {
		jsonResponseBytes, _ := json.Marshal(map[string]string{"error": "Failed to create user"})
		return string(jsonResponseBytes), http.StatusInternalServerError
	}

	var videoLikedJSON []string
	var statusCode int
	videoLikedJSON, statusCode = services.GetLastedLiked(token.TokenData)
	jsonResponseBytes, _ := json.Marshal(videoLikedJSON)
	return string(jsonResponseBytes), statusCode
}

// ----- GITHUB ----- //
func GithubLogin(c *gin.Context) (string, int) {
	utils.GithubAuth()
	if utils.GithubOauth == nil {
		jsonResponseBytes, _ := json.Marshal(map[string]string{"error": "OAuth configuration is not initialized"})
		return string(jsonResponseBytes), http.StatusInternalServerError
	}
	url := utils.GithubOauth.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	return url, http.StatusPermanentRedirect
}

func GithubLoggedIn(c *gin.Context) (*oauth2.Token, string, int) {
	var user models.User
	var token models.Token

	httpClient := utils.GithubOauth.Client(context.Background(), utils.GithubToken)
	githubClient := github.NewClient(httpClient)
	userInfo, _, err := githubClient.Users.Get(c, "")
	if err != nil {
		jsonResponseBytes, _ := json.Marshal(map[string]string{"error": "Failed to create user"})
		return nil, string(jsonResponseBytes), http.StatusInternalServerError
	}

	user.ID = primitive.NewObjectID()
	user.Username = *userInfo.Name
	user.Email = *userInfo.Email
	hashedPassword, _ := utils.GenerateHash("githubAccount")
	user.Password = hashedPassword

	token.ID = user.ID
	token.Type = "Github"
	token.TokenData = utils.GithubToken

	var status bool = storage.CreateUser(user)
	if !status {
		jsonResponseBytes, _ := json.Marshal(map[string]string{"error": "Failed to create user"})
		return nil, string(jsonResponseBytes), http.StatusInternalServerError
	}

	status = storage.CreateToken(token)
	if !status {
		jsonResponseBytes, _ := json.Marshal(map[string]string{"error": "Failed to create user"})
		return nil, string(jsonResponseBytes), http.StatusInternalServerError
	}

	return utils.GithubToken, "", http.StatusOK
}
