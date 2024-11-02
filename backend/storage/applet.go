package storage

import (
	"area/models"
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ExistApplet(applet models.Applet) bool {
	collection := DB.Collection("applets")
	var actualApplet models.Applet

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := collection.FindOne(ctx, bson.M{"user_id": applet.ID_User, "if_type": applet.IfType, "that_type": applet.ThatType}).Decode(&actualApplet)
	if err != nil {
		log.Printf("Applet not found: %v", err)
		return false
	}
	return true
}

func UpdateApplet(newApplet models.Applet) bool {
	collection := DB.Collection("applets")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	oldApplet, status := GetApplet(newApplet.ID, newApplet.ID_User)
	if !status {
		return false
	}

	update := bson.M{
		"$set": bson.M{
			"user_id":    oldApplet.ID_User,
			"that":       newApplet.That,
			"that_type":  newApplet.ThatType,
			"if":         newApplet.If,
			"if_type":    newApplet.IfType,
			"updated_at": time.Now(),
			"created_at": oldApplet.CreatedAt,
		},
	}

	_, err := collection.UpdateOne(ctx, bson.M{"user_id": newApplet.ID_User, "if_type": newApplet.IfType, "that_type": newApplet.ThatType}, update)
	if err != nil {
		log.Printf("Error while updating applet: %v", err)
		return false
	}
	return true
}

func CreateApplet(newApplet models.Applet) bool {
	collection := DB.Collection("applets")
	newApplet.ID = primitive.NewObjectID()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := collection.InsertOne(ctx, newApplet)
	if err != nil {
		log.Printf("Error while creating applet: %v", err)
		return false
	}
	return true
}

func DeleteApplet(applet models.Applet) bool {
	collection := DB.Collection("applets")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := collection.DeleteOne(ctx, bson.M{"user_id": applet.ID_User, "if_type": applet.IfType, "that_type": applet.ThatType})
	if err != nil {
		log.Printf("Error while deleting applet: %v", err)
		return false
	}
	return true
}

func GetApplet(appletID primitive.ObjectID, userID primitive.ObjectID) (models.Applet, bool) {
	collection := DB.Collection("applets")
	var applet models.Applet

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := collection.FindOne(ctx, bson.M{"_id": appletID}).Decode(&applet)
	if err != nil {
		return models.Applet{}, false
	}
	return applet, true
}

func GetAppletUserIfThat(userID string, ifType string, thatType string) models.Applet {
	collection := DB.Collection("applets")
	var applet models.Applet

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := collection.FindOne(ctx, bson.M{"user_id": userID, "if_type": ifType, "that_type": thatType}).Decode(&applet)
	if err != nil {
		log.Printf("Error while retrieving applet by user_id and types (if and that): %v", err)
		return models.Applet{}
	}
	return applet
}

func GetApplets(userID primitive.ObjectID) []models.Applet {
	collection := DB.Collection("applets")
	var applets []models.Applet
	var applet models.Applet

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cur, err := collection.Find(ctx, bson.M{"user_id": userID})
	if err != nil {
		log.Printf("Error while retrieving applets by user_id: %v", err)
		return []models.Applet{}
	}
	for cur.Next(ctx) {
		cur.Decode(&applet)
		applets = append(applets, applet)
	}
	return applets
}
