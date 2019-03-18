package database

import "github.com/stretchr/testify/assert"

func (s *DatabaseSuite) TestUser() {
	users, _ := s.db.GetUsers()
	s.T().Log(users)
	assert.Len(s.T(), users, 1)
}
