package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/oauth2"
)

type Token struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	Type         string             `bson:"type"`
	Tokens       *oauth2.Token      `bson:tokens,omitempty`
	RefreshToken string             `bson:"refresh_tokens,omitempty"`
}
