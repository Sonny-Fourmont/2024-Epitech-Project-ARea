package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/oauth2"
)

type Token struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	UserID    primitive.ObjectID `bson:"user_id, omitempty"`
	Type      string             `bson:"type"`
	CreatedAt time.Time          `bson:"created_at, omitempty"`
	UpdatedAt time.Time          `bson:"updated_at, omitempty"`
	TokenData *oauth2.Token      `bson:"token_data, omitempty"`
}
