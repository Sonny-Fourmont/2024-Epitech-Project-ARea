package controllers

import (
	"area/config"
	"area/models"
	"area/storage"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	CreatedAt time.Time          `bson:"created_at, omitempty"`
	UpdatedAt time.Time          `bson:"updated_at, omitempty"`
	Username  string             `bson:"username" gorm:"unique"`
	Email     string             `bson:"email" gorm:"unique"`
	Password  string             `bson:"password"`
}

func TestGetMe(t *testing.T) {
	config.LoadConfig()
	storage.ConnectDatabase()
	defer storage.ResetDatabase()

	user := models.User{
		ID:        primitive.NewObjectID(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Username:  "test",
		Email:     "",
		Password:  "test",
	}
	storage.CreateUser(user)
	status := storage.ExistUser(user)
	assert.True(t, status)
}
