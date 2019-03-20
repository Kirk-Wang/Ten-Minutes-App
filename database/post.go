package database

import (
	"context"
	"github.com/lotteryjs/ten-minutes-api/model"
)

// CreatePost creates a post.
func (d *TenDatabase) CreatePost(post *model.Post) error {
	if _, err := d.DB.Collection("posts").
		InsertOne(context.Background(), post); err != nil {
		return err
	}
	return nil
}
