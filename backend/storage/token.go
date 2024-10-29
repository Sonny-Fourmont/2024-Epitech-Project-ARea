package storage

import (
	"area/models"
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetTokens(userID primitive.ObjectID) ([]models.Token, bool) {
	collection := DB.Collection("tokens")
	var tokens []models.Token

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{"user_id": userID})
	if err != nil {
		log.Printf("Error while getting tokens: %v", err)
		return tokens, false
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var token models.Token
		cursor.Decode(&token)
		tokens = append(tokens, token)
	}
	return tokens, true
}

func ExistToken(token models.Token) bool {
	collection := DB.Collection("tokens")
	var actualToken models.Token

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := collection.FindOne(ctx, bson.M{"user_id": token.UserID, "type": token.Type}).Decode(&actualToken)
	if err != nil {
		log.Printf("Token not found: %v", err)
		return false
	}
	return true
}

func CreateORUpdateToken(newToken models.Token) bool {
	if ExistToken(newToken) {
		return UpdateToken(newToken)
	}
	return CreateToken(newToken)
}

func UpdateToken(newToken models.Token) bool {
	collection := DB.Collection("tokens")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	oldToken := GetTokenByUserIDAndType(newToken.UserID.Hex(), newToken.Type)

	update := bson.M{
		"$set": bson.M{
			"user_id":    newToken.UserID,
			"type":       newToken.Type,
			"token_data": newToken.TokenData,
			"updated_at": time.Now(),
			"created_at": oldToken.CreatedAt,
		},
	}
	_, err := collection.UpdateOne(ctx, bson.M{"user_id": newToken.UserID, "type": newToken.Type}, update)
	if err != nil {
		log.Printf("Error while updating token: %v", err)
		return false
	}
	return true
}

func CreateToken(newToken models.Token) bool {
	collection := DB.Collection("tokens")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := collection.InsertOne(ctx, newToken)
	if err != nil {
		log.Printf("Error while creating token: %v", err)
		return false
	}
	return true
}

func DeleteToken(token models.Token) bool {
	collection := DB.Collection("tokens")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := collection.DeleteOne(ctx, bson.M{"user_id": token.UserID, "type": token.Type})
	if err != nil {
		log.Printf("Error while deleting token: %v", err)
		return false
	}
	return true
}

func GetTokenByUserIDAndType(userID string, tokenType string) models.Token {
	collection := DB.Collection("tokens")
	var token models.Token

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := collection.FindOne(ctx, bson.M{"user_id": userID, "type": tokenType}).Decode(&token)
	if err != nil {
		log.Printf("Error while retrieving token by user_id and type: %v", err)
		return models.Token{}
	}
	return token
}
