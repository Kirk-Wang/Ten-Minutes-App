package database

import (
	"fmt"
	"github.com/lotteryjs/ten-minutes-app/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
)

func TestDatabaseSuite(t *testing.T) {
	suite.Run(t, new(DatabaseSuite))
}

type DatabaseSuite struct {
	suite.Suite
	db *TenDatabase
}

func (s *DatabaseSuite) BeforeTest(suiteName, testName string) {
	s.T().Log("--BeforeTest--")
	db, _ := New("mongodb://root:123456@localhost:27017", "tenapi")
	s.db = db
}

func (s *DatabaseSuite) AfterTest(suiteName, testName string) {
	s.db.Close()
}

func (s *DatabaseSuite) TestPost() {
	s.db.DB.Collection("posts").Drop(nil)

	var err error
	for i := 1; i <= 25; i++ {
		// user1
		UserID, _ := primitive.ObjectIDFromHex("5c99bd941ba7b2304ad8c52a")
		article := (&model.Post{
			UserID: UserID,
			Title:  fmt.Sprintf("tile%d", i),
			Body:   "quia et suscipit\nsuscipit recusandae consequuntur expedita et cum\nreprehenderit molestiae ut ut quas totam\nnostrum rerum est autem sunt rem eveniet architecto",
		}).New()
		s.db.CreatePost(article)
	}
	assert.Nil(s.T(), err)
}

func (s *DatabaseSuite) TestGetPostByID() {
	id, _ := primitive.ObjectIDFromHex("5c92e6199929adef73bceea1")
	post := s.db.GetPostByID(id)
	assert.Equal(s.T(), "tile1", post.Title)
}

func (s *DatabaseSuite) TestUpdatePost() {
	id, _ := primitive.ObjectIDFromHex("5c92e6199929adef73bceea1")
	userID, _ := primitive.ObjectIDFromHex("5c8f9a83da2c3fed4eee9dc1")

	post := &model.Post{
		ID:     id,
		UserID: userID,
		Title:  "title1",
		Body:   "title1bodytitle1body",
	}

	assert.Equal(s.T(), post, s.db.UpdatePost(post))
}
