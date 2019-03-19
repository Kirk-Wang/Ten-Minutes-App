package database

import (
	"context"
	"github.com/lotteryjs/ten-minutes-api/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// GetUsersOptions is query params
type GetUsersOptions struct {
	Skip    *int64
	Limit   *int64
	sortKey string
	sortVal string
}

func (d *TenDatabase) NewGetUsersOptions() *GetUsersOptions {
	return &GetUsersOptions{}
}

// GetUsers returns all users.
// start, end int, order, sort string
func (d *TenDatabase) GetUsers(opts *GetUsersOptions) []*model.User {
	var users []*model.User
	cursor, err := d.DB.Collection("users").
		Find(context.Background(), bson.D{},
			&options.FindOptions{
				Skip:  opts.Skip,
				Sort:  bson.D{bson.E{Key: opts.sortKey, Value: opts.sortVal}},
				Limit: opts.Limit,
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
