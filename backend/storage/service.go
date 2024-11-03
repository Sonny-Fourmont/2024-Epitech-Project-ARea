package storage

import (
	"area/models"
	"context"
	"log"
	"reflect"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func StoreAndCheckResponse(appletID primitive.ObjectID, response []string, ifType string) bool {
	service := GetServiceByAppletIDAndType(appletID, ifType)

	if service.AppletID != primitive.NilObjectID {
		if reflect.DeepEqual(service.Latest, response) {
			return false
		}
	}
	CreateORUpdateService(models.Service{AppletID: appletID, Type: ifType, Latest: response})
	return true
}

func ExistService(service models.Service) bool {
	collection := DB.Collection("services")
	var actualService models.Service

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := collection.FindOne(ctx, bson.M{"applet_id": service.AppletID, "type": service.Type}).Decode(&actualService)
	if err != nil {
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

func UpdateService(service models.Service) bool {
	collection := DB.Collection("services")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	update := bson.M{
		"$set": bson.M{
			"applet_id": service.AppletID,
			"type":      service.Type,
			"latest":    service.Latest,
		},
	}
	_, err := collection.UpdateOne(ctx, bson.M{"applet_id": service.AppletID, "type": service.Type}, update)
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

	_, err := collection.DeleteOne(ctx, bson.M{"applet_id": service.AppletID, "type": service.Type})
	if err != nil {
		log.Printf("Error while deleting service: %v", err)
		return false
	}
	return true
}

func GetService(serviceID primitive.ObjectID) (models.Service, bool) {
	collection := DB.Collection("services")
	var service models.Service

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := collection.FindOne(ctx, bson.M{"_id": serviceID}).Decode(&service)
	if err != nil {
		return models.Service{}, false
	}
	return service, true
}

func GetServiceByAppletIDAndType(appletID primitive.ObjectID, serviceType string) models.Service {
	collection := DB.Collection("services")
	var service models.Service

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := collection.FindOne(ctx, bson.M{"applet_id": appletID, "type": serviceType}).Decode(&service)
	if err != nil {
		return models.Service{}
	}
	return service
}
