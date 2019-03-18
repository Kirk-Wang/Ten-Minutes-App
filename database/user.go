package database

import (
	"context"

	"github.com/lotteryjs/ten-minutes-api/model"
	"go.mongodb.org/mongo-driver/bson"
)

// GetUsers returns all users.
func (d *TenDatabase) GetUsers() ([]*model.User, error) {
	var users []*model.User
	cursor, err := d.DB.Collection("users").Find(
		context.Background(),
		bson.D{},
	)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		user := &model.User{}
		if err := cursor.Decode(user); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

// CreateUser creates a user.
func (d *TenDatabase) CreateUser(user *model.User) error {
	if _, err := d.DB.Collection("users").InsertOne(context.Background(), user); err != nil {
		return err
	}
	return nil
}
