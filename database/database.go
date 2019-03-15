package database

import (
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/clientopt"
)

var database *mongo.Database

func New(connection, dbname) (*mongo.Database, error) {
	if database == nil {
		client, err := mongo.NewClientWithOptions(
			env.Get().MongoURL,
			clientopt.ServerSelectionTimeout(time.Second))
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		err = client.Connect(context.TODO())
		if err != nil {
			log.Fatal(err)
			return nil, err
		}

		database = client.Database(dbname)
	}
	return database, nil
}
