package controllers

import (
	"area/middlewares"
	"area/models"
	"area/storage"
	"area/utils"
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func GetMe(c *gin.Context) (string, int) {
	return middlewares.GetClient(c).Hex(), http.StatusOK
}
