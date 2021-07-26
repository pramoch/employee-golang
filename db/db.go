package DB

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var Client *mongo.Client

func Connect(ctx context.Context) {
	uri := "mongodb://localhost:27017"

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	// Ping the primary
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected and pinged.")

	Client = client
}

func Disconnect(ctx context.Context) {
	if err := Client.Disconnect(ctx); err != nil {
		fmt.Println("Error disconnect")
		panic(err)
	}
}
