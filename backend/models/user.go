package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at,omitempty"`
	Username  string             `bson:"username" gorm:"unique"`
	Email     string             `bson:"email" gorm:"unique"`
	Password  string             `bson:"password"`
}
