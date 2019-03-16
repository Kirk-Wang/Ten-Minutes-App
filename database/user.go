package database

import (
	"context"

	"github.com/lotteryjs/ten-minutes-api/model"
	"go.mongodb.org/mongo-driver/bson"
)

// GetUsers returns all users.
func (d *TenDatabase) GetUsers() []*model.User {
	// var users []*model.User
	// d.DB.Find(&users)
	// return users
	cursor, err := d.DB.Collection("users").Find(
		context.Background(),
		bson.D{},
	)
	defer cursor.Close(context.Background())

	users := []*model.User
	for cursor.Next(context.Background()) {
		user := &model.User{}
		if err := cursor.Decode(user); err != nil {
			return nil, err
		}
		feedbacks = append(feedbacks, feedback)
	}
}
