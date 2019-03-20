package database

import (
	"github.com/lotteryjs/ten-minutes-api/model"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *DatabaseSuite) TestPost() {
	s.db.DB.Collection("posts").Drop(nil)

	// user1
	UserID, _ := primitive.ObjectIDFromHex("5c8f9a83da2c3fed4eee9dc1")
	article := (&model.Post{
		UserID: UserID,
		Title:  "sunt aut facere repellat provident occaecati excepturi optio reprehenderit",
		Body:   "quia et suscipit\nsuscipit recusandae consequuntur expedita et cum\nreprehenderit molestiae ut ut quas totam\nnostrum rerum est autem sunt rem eveniet architecto",
	}).New()

	err := s.db.CreatePost(article)
	assert.Nil(s.T(), err)
}
