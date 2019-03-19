package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/lotteryjs/ten-minutes-api/model"
	"go.mongodb.org/mongo-driver/bson"
)

// GetUsers returns all users.
func (d *TenDatabase) GetUsers() []*model.User {
	var users []*model.User
	var skip int64 = 5
	cursor, err := d.DB.Collection("users").
		Find(nil, bson.D{}, &options.FindOptions{Skip: &skip})
	if err != nil {
		return nil
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		user := &model.User{}
		if err := cursor.Decode(user); err != nil {
			return nil
		}
		users = append(users, user)
	}

	return users
}

// CreateUser creates a user.
func (d *TenDatabase) CreateUser(user *model.User) error {
	if _, err := d.DB.Collection("users").
		InsertOne(context.Background(), user); err != nil {
		return err
	}
	return nil
}

// GetUserByName returns the user by the given name or nil.
func (d *TenDatabase) GetUserByName(name string) *model.User {
	var user *model.User
	err := d.DB.Collection("users").
		FindOne(nil, bson.D{{Key: "name", Value: name}}).
		Decode(&user)
	if err != nil {
		return nil
	}
	return user
}
