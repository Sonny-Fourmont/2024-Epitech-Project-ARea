package storage

import (
	"area/models"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func RetrieveUser(token_name string, ctx context.Context) primitive.ObjectID {
	collection := DB.Collection("tokens")
	var token models.Token

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := collection.FindOne(ctx, bson.M{"token_data.accesstoken": token_name}).Decode(&token)
	if err != nil {
		println(token_name)
		return primitive.NilObjectID
	}

	collection = DB.Collection("users")
	var user models.User
	err = collection.FindOne(ctx, bson.M{"_id": token.ID}).Decode(&user)
	if err != nil {
		return primitive.NilObjectID
	}
	return user.ID
}
