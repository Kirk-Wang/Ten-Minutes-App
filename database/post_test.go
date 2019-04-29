package database

import (
	"github.com/lotteryjs/ten-minutes-app/model"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *DatabaseSuite) TestCreatePost() {
	s.db.DB.Collection("posts").Drop(nil)

	user := s.db.GetUserByName("Graham")

	article := (&model.Post{
		UserID: user.ID,
		Title:  "title1",
		Body:   "quia et suscipit\nsuscipit recusandae consequuntur expedita et cum\nreprehenderit molestiae ut ut quas totam\nnostrum rerum est autem sunt rem eveniet architecto",
	}).New()

	s.db.CreatePost(article)
	post := s.db.GetPostByID(article.ID)

	assert.Equal(s.T(), article, post)
}

func (s *DatabaseSuite) TestCountPost() {
	assert.Equal(s.T(), "1", s.db.CountPost(nil))
}

func (s *DatabaseSuite) TestGetPostByID() {
	id, _ := primitive.ObjectIDFromHex("5cc5ca2f6a670dd59ea3a590")
	post := s.db.GetPostByID(id)
	assert.Equal(s.T(), "title1", post.Title)
}
