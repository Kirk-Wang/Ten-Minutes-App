package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

var database *mongo.Database

func New(connection string, dbname string) (*mongo.Database, error) {
	if database == nil {
		client, err := mongo.NewClient(options.Client().ApplyURI(connection))
		if err != nil {
			return nil, err
		}
		ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
		defer cancel()
		err = client.Connect(ctx)
		if err != nil {
			return nil, err
		}
		database = client.Database(dbname)
	}
	return database, nil
}
