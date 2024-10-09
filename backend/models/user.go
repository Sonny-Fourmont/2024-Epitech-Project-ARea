package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Services struct {
	GoogleEmail    string `bson:"google_email,omitempty"`
	GithubEmail    string `bson:"github_email,omitempty"`
	MicrosoftEmail string `bson:"microsoft_email,omitempty"`
}

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at,omitempty"`
	Username  string             `bson:"username" gorm:"unique"`
	Email     string             `bson:"email" gorm:"unique"`
	Password  string             `bson:"password"`
	Services  Services           `bson:"services"`
}
