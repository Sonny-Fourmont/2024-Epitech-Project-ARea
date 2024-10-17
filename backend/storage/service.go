package storage

import (
	"area/models"
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func ExistService(service models.Service) bool {
	collection := DB.Collection("services")
	var actualService models.Service

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := collection.FindOne(ctx, bson.M{"user_id": service.UserID, "type": service.Type}).Decode(&actualService)
	if err != nil {
		log.Printf("Service not found: %v", err)
		return false
	}
	return true
}

func CreateORUpdateService(newService models.Service) bool {
	if ExistService(newService) {
		return UpdateService(newService)
	}
	return CreateService(newService)
}

func UpdateService(newService models.Service) bool {
	collection := DB.Collection("services")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	update := bson.M{
		"$set": bson.M{
			"user_id": newService.UserID,
			"type":    newService.Type,
			"latest":  newService.Latest,
		},
	}
	_, err := collection.UpdateOne(ctx, bson.M{"user_id": newService.UserID, "type": newService.Type}, update)
	if err != nil {
		log.Printf("Error while updating service: %v", err)
		return false
	}
	return true
}

func CreateService(newService models.Service) bool {
	collection := DB.Collection("services")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := collection.InsertOne(ctx, newService)
	if err != nil {
		log.Printf("Error while creating service: %v", err)
		return false
	}
	return true
}

func DeleteService(service models.Service) bool {
	collection := DB.Collection("services")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := collection.DeleteOne(ctx, bson.M{"user_id": service.UserID, "type": service.Type})
	if err != nil {
		log.Printf("Error while deleting service: %v", err)
		return false
	}
	return true
}

func GetServiceByUserIDAndType(userID string, serviceType string) models.Service {
	collection := DB.Collection("services")
	var service models.Service

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := collection.FindOne(ctx, bson.M{"user_id": userID, "type": serviceType}).Decode(&service)
	if err != nil {
		log.Printf("Error while retrieving service by user_id and type: %v", err)
		return models.Service{}
	}
	return service
}
