package database

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func ConnectMongoDB() (*mongo.Database, error) {
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		uri = "mongodb://bitaksi-mongodb:27017"
	}

	clientOpts := options.Client().ApplyURI(uri)
	clientOpts.SetBSONOptions(&options.BSONOptions{
		ObjectIDAsHexString: true,
	})

	client, err := mongo.Connect(clientOpts)
	if err != nil {
		return nil, err
	}

	if err := client.Ping(context.Background(), nil); err != nil {
		return nil, err
	}

	dbName := os.Getenv("MONGODB_DATABASE")
	if dbName == "" {
		dbName = "driver_service"
	}

	return client.Database(dbName), nil
}
