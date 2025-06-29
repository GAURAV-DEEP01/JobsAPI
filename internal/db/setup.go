package db

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	JobDB         string
	JobCollection string = "Job"
)

var MongoClient *mongo.Client

func InitMongo(uri string) *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}

	if err := client.Ping(ctx, nil); err != nil {
		log.Fatal("MongoDB ping failed:", err)
	}

	JobDB = os.Getenv("DATABASE_NAME")
	log.Println("MongoDB connected")
	MongoClient = client
	return client
}

func GetCollection(collectionName string) *mongo.Collection {
	return MongoClient.Database(JobDB).Collection(collectionName)
}
