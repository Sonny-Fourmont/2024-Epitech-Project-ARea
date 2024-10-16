package storage

import (
	"area/models"
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetUserByID(id primitive.ObjectID) models.User {
	collection := DB.Collection("users")
	var user models.User

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := collection.FindOne(ctx, bson.M{"_id": id}).Decode(&user)
	if err != nil {
		log.Printf("Error while retrieving user by id: %v", err)
		return models.User{}
	}
	return user
}

func GetUserIDByEmail(email string) primitive.ObjectID {
	collection := DB.Collection("users")
	var user models.User

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := collection.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		log.Printf("Error while retrieving user by email: %v", err)
		return primitive.NilObjectID
	}
	return user.ID
}

func GetUserIDByToken(token string) primitive.ObjectID {
	collection := DB.Collection("users")
	var user models.User

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := collection.FindOne(ctx, bson.M{"token_data.accesstoken": token}).Decode(&user)
	if err != nil {
		log.Printf("Error while retrieving user by token: %v", err)
		return primitive.NilObjectID
	}
	return user.ID
}

func UpdateUser(user models.User, newUser models.User) bool {
	collection := DB.Collection("users")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if newUser.Email != "" {
		user.Email = newUser.Email
	}
	if newUser.Password != "" {
		user.Password = newUser.Password
	}
	if newUser.Username != "" {
		user.Username = newUser.Username
	}
	user.UpdatedAt = time.Now()
	_, err := collection.UpdateByID(ctx, user.ID, bson.M{"$set": user})
	if err != nil {
		log.Printf("Error while creating user: %v", err)
		return false
	}
	return true
}

func CreateUser(user models.User) bool {
	collection := DB.Collection("users")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	_, err := collection.InsertOne(ctx, user)
	if err != nil {
		log.Printf("Error while creating user: %v", err)
		return false
	}
	return true
}

func UpdateUserByEmail(email string, newUser models.User) bool {
	var user models.User = GetUserByID(GetUserIDByEmail(email))
	return UpdateUser(user, newUser)
}

func UpdateUserByToken(token string, newUser models.User) bool {
	var user models.User = GetUserByID(GetUserIDByToken(token))
	return UpdateUser(user, newUser)
}

func DeleteUserByID(id primitive.ObjectID) bool {
	collection := DB.Collection("users")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := collection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		log.Printf("Error while deleting user by id: %v", err)
		return false
	}
	return true
}

func DeleteUserByToken(token string) bool {
	return DeleteUserByID(GetUserIDByToken(token))
}
