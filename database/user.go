package database

import (
	"go.mongodb.org/mongo-driver/mongo"
)

// GetUsers returns all users.
func (d *mongo.Database) GetUsers() []*model.User {
	var users []*model.User
	d.DB.Find(&users)
	return users
}
