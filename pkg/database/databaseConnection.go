package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBinstance() *mongo.Client {
	MongoDb := "mongodb://localhost:27017"
	//MongoDb containing the MongoDB connection string,
	//specifying the server's address and port.
	fmt.Print(MongoDb)
	client, err := mongo.NewClient(options.Client().ApplyURI(MongoDb))
	//creates a new MongoDB client using the provided connection string
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	//This line creates a context with a timeout of 10 seconds for connecting to MongoDB
	//The context will be used to manage the connection's lifecycle
	defer cancel()
	//his ensures that the context will be canceled when the function DBinstance exits
	//helping to release any resources related to the context.
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("Connected to MongoDB")
	return client
}

var Client *mongo.Client = DBinstance() //a global variable

func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	var collection *mongo.Collection = client.Database("restaurant").Collection(collectionName)
	return collection
}
