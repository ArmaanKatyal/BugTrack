package config

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ClientConnection returns a connection to the database
func ClientConnection() *mongo.Client {
	// create the client connection
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(ViperEnvVariable("MONGO_DB_URL")))
	// check for errors
	if err != nil {
		panic(err)
	}
	// check the connection
	if err := client.Ping(context.TODO(), nil); err != nil {
		panic(err)
	}
	// return the client connection
	return client
}
