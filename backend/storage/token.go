package storage

import (
	"area/models"
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetToken(UserId primitive.ObjectID, serviceType string) models.Token {
	collection := DB.Collection("tokens")
	var token models.Token

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := collection.FindOne(ctx, bson.M{"user_id": UserId, "type": serviceType}).Decode(&token)
	if err != nil {
		log.Printf("Error while retrieving token by id and type: %v", err)
		return models.Token{}
	}
	return token
}

func GetTokenByID(id primitive.ObjectID) models.Token {
	collection := DB.Collection("tokens")
	var token models.Token

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := collection.FindOne(ctx, bson.M{"_id": id}).Decode(&token)
	if err != nil {
		log.Printf("Error while retrieving token by id and type: %v", err)
		return models.Token{}
	}
	return token
}

func UpdateToken(token models.Token, newToken models.Token) bool {
	collection := DB.Collection("tokens")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if newToken.ID != primitive.NilObjectID {
		token.ID = newToken.ID
	}
	if newToken.Type != "" {
		token.Type = newToken.Type
	}
	if newToken.TokenData != nil {
		token.TokenData = newToken.TokenData
	}
	if newToken.UserID != primitive.NilObjectID {
		token.UserID = newToken.UserID
	}
	token.UpdatedAt = time.Now()
	_, err := collection.UpdateByID(ctx, token.ID, bson.M{"$set": token})
	if err != nil {
		log.Printf("Error while updating token: %v", err)
		return false
	}
	return true
}

func CreateToken(token models.Token) bool {
	collection := DB.Collection("tokens")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	token.CreatedAt = time.Now()
	token.UpdatedAt = time.Now()
	_, err := collection.InsertOne(ctx, token)
	if err != nil {
		log.Printf("Error while creating token: %v", err)
		return false
	}
	return true
}

func DeleteToken(UserId primitive.ObjectID, serviceType string) bool {
	collection := DB.Collection("token")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := collection.DeleteOne(ctx, bson.M{"user_id": UserId, "type": serviceType})
	if err != nil {
		log.Printf("Error while deleting token by id and type: %v", err)
		return false
	}
	return true
}
