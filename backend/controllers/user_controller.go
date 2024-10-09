package controllers

import (
	"area/models"
	"area/storage"
	"area/utils"
	"context"
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

func RegisterUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	hashedPassword, _ := utils.GenerateHash(user.Password)
	user.Password = hashedPassword

	collection := storage.DB.Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := collection.InsertOne(ctx, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

func GetUser(c *gin.Context) {
	id := c.Param("id")

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	collection := storage.DB.Collection("users")
	var user models.User

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&user)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":       user.ID.Hex(),
		"username": user.Username,
		"email":    user.Email,
	})
}

// ----- GOOGLE ----- //
func GoogleLogin(c *gin.Context) {
	utils.GoogleAuth()
	if utils.GoogleOauth == nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "OAuth configuration is not initialized"})
	}
	url := utils.GoogleOauth.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	c.Redirect(http.StatusPermanentRedirect, url)
}

func GoogleLoggedIn(c *gin.Context) {
	var user models.User
	var tokens models.Token

	httpClient := utils.GoogleOauth.Client(context.Background(), utils.GoogleToken)
	gmail, _ := gmail.NewService(context.Background(), option.WithHTTPClient(httpClient))
	googleUser, _ := gmail.Users.GetProfile("me").Do()

	user.Username = googleUser.EmailAddress
	user.Email = googleUser.EmailAddress
	hashedPassword, _ := utils.GenerateHash("googleAccount")
	user.Password = hashedPassword
	user.Services.GoogleEmail = googleUser.EmailAddress
	user.Services.GoogleId = primitive.NewObjectID()
	tokens.ID = user.Services.GoogleId
	tokens.Type = "Google"
	tokens.RefreshToken = utils.GoogleToken.RefreshToken

	collection := storage.DB.Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := collection.InsertOne(ctx, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}
	log.Output(0, "User has been created!")

	collection = storage.DB.Collection("tokens")
	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = collection.InsertOne(ctx, tokens)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}
	log.Output(0, "Refresh token has been created!")

	c.JSON(http.StatusOK, utils.GoogleToken)
}

// ----- GITHUB ----- //
func GithubLogin(c *gin.Context) {
	utils.GithubAuth()
	if utils.GithubOauth == nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "OAuth configuration is not initialized"})
	}
	url := utils.GithubOauth.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	c.Redirect(http.StatusPermanentRedirect, url)
}

func GithubLoggedIn(c *gin.Context) {
	var user models.User
	var tokens models.Token

	httpClient := utils.GithubOauth.Client(context.Background(), utils.GithubToken)
	githubClient := github.NewClient(httpClient)
	userInfo, _, err := githubClient.Users.Get(c, "")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	user.ID = primitive.NewObjectID()
	user.Username = *userInfo.Name
	user.Email = *userInfo.Email
	hashedPassword, _ := utils.GenerateHash("githubAccount")
	user.Password = hashedPassword
	user.Services.GithubEmail = *userInfo.Email
	tokens.ID = user.ID
	tokens.Type = "Github"
	tokens.Tokens = utils.GithubToken

	collection := storage.DB.Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = collection.InsertOne(ctx, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}
	log.Output(0, "User has been created!")

	collection = storage.DB.Collection("tokens")
	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = collection.InsertOne(ctx, tokens)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}
	log.Output(0, "Token has been created!")

	c.JSON(http.StatusOK, utils.GithubToken)
}
