package database

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

var (
	Client *mongo.Client
	err    error
)

func CreateConnection(ctx context.Context, uri string) error {
	Client, err = mongo.Connect(options.Client().ApplyURI(uri))
	if err != nil {
		return err
	}

	err = Client.Ping(ctx, readpref.Primary())
	if err != nil {
		return err
	}

	return nil
}
