package storage

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

func ResetDatabase() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := DB.Collection("users").DeleteMany(ctx, bson.M{})
	if err != nil {
		log.Fatalf("Erreur lors de la suppression des documents de la collection 'users' : %v", err)
	}
}

func ConnectDatabase() {
	mongoURI := os.Getenv("MONGODB_URI_DEV")
	if os.Getenv("GIN_MODE") == "test" {
		mongoURI = os.Getenv("MONGODB_URI_TEST")
	}
	dbName := os.Getenv("DB_NAME")

	clientOpts := options.Client().ApplyURI(mongoURI)
	client, err := mongo.Connect(context.TODO(), clientOpts)
	if err != nil {
		log.Fatalf("Erreur lors de la connexion à MongoDB : %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("Impossible de se connecter à MongoDB : %v", err)
	}

	DB = client.Database(dbName)
}
