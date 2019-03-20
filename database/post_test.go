package database

import (
	"github.com/lotteryjs/ten-minutes-api/model"
	"github.com/stretchr/testify/assert"
)

func (s *DatabaseSuite) TestPost() {
	s.db.DB.Collection("posts").Drop(nil)

	article := (&model.Post{}).new()
}
