package database

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MakeMongoDbConnection(ctx context.Context) (*mongo.Client, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://root:example@localhost:27017"))
	if err != nil {
		panic(err)
	}
	err = client.Connect(ctx)
	return client, err
}
