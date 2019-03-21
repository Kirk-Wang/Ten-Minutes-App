package database

import (
	"context"
	"github.com/lotteryjs/ten-minutes-api/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// GetPosts returns all posts.
// start, end int, order, sort string
func (d *TenDatabase) GetPosts(paging *model.Paging) []*model.Post {
	var posts []*model.Post
	cursor, err := d.DB.Collection("posts").
		Find(context.Background(), bson.D{},
			&options.FindOptions{
				Skip:  paging.Skip,
				Sort:  bson.D{bson.E{Key: paging.SortKey, Value: paging.SortVal}},
				Limit: paging.Limit,
			})
	if err != nil {
		return nil
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		post := &model.Post{}
		if err := cursor.Decode(post); err != nil {
			return nil
		}
		posts = append(posts, post)
	}

	return posts
}

// CreatePost creates a post.
func (d *TenDatabase) CreatePost(post *model.Post) (*model.Post, error) {
	if _, err := d.DB.Collection("posts").
		InsertOne(context.Background(), post); err != nil {
		return nil, err
	}
	return post, nil
}

// GetPostByID returns the post by the given id or nil.
func (d *TenDatabase) GetPostByID(id primitive.ObjectID) *model.Post {
	var post *model.Post
	err := d.DB.Collection("posts").
		FindOne(nil, bson.D{{Key: "_id", Value: id}}).
		Decode(&post)
	if err != nil {
		return nil
	}
	return post
}

// UpdatePost updates a post.
func (d *TenDatabase) UpdatePost(post *model.Post) *model.Post {
	result := d.DB.Collection("posts").
		FindOneAndReplace(nil,
			bson.D{{Key: "_id", Value: post.ID}},
			post,
			&options.FindOneAndReplaceOptions{},
		)
	if result != nil {
		return post
	}
	return nil
}
