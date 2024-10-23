package storage

import (
	"area/models"
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func GetAllUsers() ([]models.User, bool) {
	collection := DB.Collection("users")
	var users []models.User

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Printf("Error while getting users: %v", err)
		return users, false
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var user models.User
		cursor.Decode(&user)
		users = append(users, user)
	}
	return users, true
}

func GetUserByEmail(email string) (models.User, bool) {
	collection := DB.Collection("users")
	var user models.User

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := collection.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		log.Printf("Error while getting user: %v", err)
		return user, false
	}
	return user, true
}

func ExistUser(user models.User) bool {
	collection := DB.Collection("users")
	var actualUser models.User

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := collection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&actualUser)
	if err != nil {
		log.Printf("User not found: %v", err)
		return false
	}
	return true
}

func CreateORUpdateUser(newUser models.User) bool {
	if ExistUser(newUser) {
		return UpdateUser(newUser)
	}
	return CreateUser(newUser)
}

func UpdateUser(newUser models.User) bool {
	collection := DB.Collection("users")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	update := bson.M{
		"$set": bson.M{
			"username":   newUser.Username,
			"email":      newUser.Email,
			"password":   newUser.Password,
			"updated_at": time.Now(),
			"created_at": newUser.CreatedAt,
		},
	}
	_, err := collection.UpdateOne(ctx, bson.M{"email": newUser.Email}, update)
	if err != nil {
		log.Printf("Error while updating user: %v", err)
		return false
	}
	return true
}

func CreateUser(newUser models.User) bool {
	collection := DB.Collection("users")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := collection.InsertOne(ctx, newUser)
	if err != nil {
		log.Printf("Error while creating user: %v", err)
		return false
	}
	return true
}

func DeleteUser(user models.User) bool {
	collection := DB.Collection("users")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := collection.DeleteOne(ctx, bson.M{"email": user.Email})
	if err != nil {
		log.Printf("Error while deleting user: %v", err)
		return false
	}
	return true
}
