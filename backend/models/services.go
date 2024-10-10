package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/oauth2"
)

type YoutubeLikedService struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	ID_User   primitive.ObjectID `bson:"user_id, omitempty"`
	Latest    string             `bson:"latest_liked"`
	TokenData *oauth2.Token      `bson:"token_data,omitempty"`
}
