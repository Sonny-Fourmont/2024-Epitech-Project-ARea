package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Services struct {
	GoogleEmail    string             `bson:"google_email,omitempty"`
	GoogleId       primitive.ObjectID `bson:"google_id,omitempty"`
	GithubEmail    string             `bson:"github_email,omitempty"`
	GithubId       primitive.ObjectID `bson:"github_id,omitempty"`
	MicrosoftEmail string             `bson:"microsoft_email,omitempty"`
	MicrosoftId    primitive.ObjectID `bson:"microsoft_id,omitempty"`
}

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at,omitempty"`
	Username  string             `bson:"username" gorm:"unique"`
	Email     string             `bson:"email" gorm:"unique"`
	Password  string             `bson:"password"`
	Services  Services           `bson:"services"`
	AppletId  primitive.ObjectID `bson:"applet_id,omitempty"`
}
