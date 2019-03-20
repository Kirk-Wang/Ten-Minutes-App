package database

import (
	"github.com/lotteryjs/ten-minutes-api/model"
	"github.com/stretchr/testify/assert"
)

func (s *DatabaseSuite) TestUser() {
	s.db.DB.Collection("posts").Drop(nil)

	article := (&model.Post{}).new()
}
