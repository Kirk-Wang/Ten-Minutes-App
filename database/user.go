package database

import (
	"context"
	"github.com/lotteryjs/ten-minutes-api/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// GetUsers returns all users.
// _end=5&_order=DESC&_sort=id&_start=0
// start, end int, order, sort string
func (d *TenDatabase) GetUsers() []*model.User {
	var users []*model.User
	var skip int64 = 40
	var limit int64 = 5
	cursor, err := d.DB.Collection("users").
		Find(context.Background(), bson.D{},
			&options.FindOptions{
				Skip:  &skip,
				Sort:  bson.D{bson.E{Key: "_id", Value: -1}},
				Limit: &limit,
			})
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
