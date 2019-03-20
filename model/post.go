package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// The Post holds
type Post struct {
	ID     primitive.ObjectID `bson:"_id" json:"id"`
	UserID primitive.ObjectID `bson:"userId" json:"userId"`
	Title  string             `bson:"title" json:"title"`
	Body   string             `bson:"body" json:"body"`
}

// New is an instance
func (p *Post) New() *Post {
	return &Post{
		ID:     primitive.NewObjectID(),
		UserID: p.UserID,
		Title:  p.Title,
		Body:   p.Body,
	}
}
