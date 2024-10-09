package controllers

import (
	"area/models"
	"area/storage"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateAREA(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "AREA created successfully"})
}

func GetApplets(c *gin.Context) {
	token := c.GetHeader("access_token")
	if token == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing Token"})
	}

	collection := storage.DB.Collection("applets")
	var applets []models.Applet
	var applet models.Applet

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	userID := storage.RetrieveUser(token, ctx)
	cur, err := collection.Find(ctx, bson.M{"user_id": userID})
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"applet_array": "[]",
		})
		return
	}
	for cur.Next(ctx) {
		cur.Decode(&applet)
		applets = append(applets, applet)
	}
	c.JSON(http.StatusOK, gin.H{
		"applet_array": applets,
	})
}

func AddApplet(c *gin.Context) {
	token := c.GetHeader("access_token")
	if token == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing Token"})
	}

	var applet models.Applet
	if err := c.ShouldBindJSON(&applet); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	collection := storage.DB.Collection("applets")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var userID primitive.ObjectID = storage.RetrieveUser(token, ctx)
	if userID == primitive.NilObjectID {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid token"})
		return
	}
	applet.ID_User = userID
	_, err := collection.InsertOne(ctx, applet)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create applet"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Applet added successfully"})
}
