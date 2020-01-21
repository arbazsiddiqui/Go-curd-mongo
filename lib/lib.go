package lib

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var Client *mongo.Client

func init() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	
	Client, _ = mongo.Connect(context.TODO(), clientOptions)
	
	err := Client.Ping(context.TODO(), nil)
	
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Println("Connected to MongoDB!")
}

func GetClient() *mongo.Client {
	return Client
}
