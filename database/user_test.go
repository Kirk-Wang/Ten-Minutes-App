package database

import (
	"github.com/lotteryjs/ten-minutes-api/model"
	"github.com/stretchr/testify/assert"
)

func (s *DatabaseSuite) TestUser() {
	nicories := &model.User{Name: "kirk", Pass: []byte{1, 2, 3, 4}, Admin: true}
	err := s.db.CreateUser(nicories)
	// s.T().Log(err)
	assert.Nil(s.T(), err)
}
