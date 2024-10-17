package controllers

import (
	"area/models"
	"area/storage"
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateAREA(c *gin.Context) (string, int) {
	jsonResponseBytes, _ := json.Marshal(map[string]string{"message": "AREA created successfully"})
	return string(jsonResponseBytes), http.StatusOK
}

func GetApplets(c *gin.Context) (string, int) {
	var err error
	token := c.GetHeader("access_token")
	if token == "" {
		jsonResponseBytes, _ := json.Marshal(map[string]string{"error": "Missing Token"})
		return string(jsonResponseBytes), http.StatusBadRequest
	}

	collection := storage.DB.Collection("applets")
	var applets []models.Applet
	var applet models.Applet

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var userID primitive.ObjectID = storage.GetUserIDByToken(token)
	cur, err := collection.Find(ctx, bson.M{"user_id": userID})
	if err != nil {
		jsonReponseBytes, _ := json.Marshal("[]")
		return string(jsonReponseBytes), http.StatusOK
	}
	for cur.Next(ctx) {
		cur.Decode(&applet)
		applets = append(applets, applet)
	}

	response := struct {
		Applets []models.Applet `json:"applet_array"`
	}{Applets: applets}
	jsonResponseBytes, _ := json.Marshal(response)
	return string(jsonResponseBytes), http.StatusOK
}

func AddApplet(c *gin.Context) (string, int) {
	token := c.GetHeader("access_token")
	if token == "" {
		jsonResponseBytes, _ := json.Marshal(map[string]string{"error": "Missing Token"})
		return string(jsonResponseBytes), http.StatusBadRequest
	}

	var applet models.Applet
	if err := c.ShouldBindJSON(&applet); err != nil {
		jsonResponseBytes, _ := json.Marshal(map[string]string{"error": "Invalid JSON"})
		return string(jsonResponseBytes), http.StatusBadRequest
	}

	collection := storage.DB.Collection("applets")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var userID primitive.ObjectID = storage.GetUserIDByToken(token)
	if userID == primitive.NilObjectID {
		jsonResponseBytes, _ := json.Marshal(map[string]string{"error": "Invalid Token"})
		return string(jsonResponseBytes), http.StatusInternalServerError
	}
	applet.ID_User = userID
	_, err := collection.InsertOne(ctx, applet)
	if err != nil {
		jsonResponseBytes, _ := json.Marshal(map[string]string{"error": "Failed to create applet"})
		return string(jsonResponseBytes), http.StatusInternalServerError
	}
	jsonResponseBytes, _ := json.Marshal(map[string]string{"message": "Applet added successfully"})
	return string(jsonResponseBytes), http.StatusOK
}
