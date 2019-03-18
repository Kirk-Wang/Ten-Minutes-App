package database

import (
	"github.com/lotteryjs/ten-minutes-api/model"
	"github.com/stretchr/testify/assert"
)

func (s *DatabaseSuite) TestUser() {
	s.db.DB.Collection("users").Drop(nil)

	kirk := (&model.User{Name: "kirk", Pass: []byte{1, 2, 3, 4}, Admin: true}).New()
	err := s.db.CreateUser(kirk)
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), kirk, s.db.GetUserByName("kirk"))

}
