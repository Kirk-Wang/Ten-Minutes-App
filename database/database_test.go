package database

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
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
	db, err := New("mongodb://root:123456@localhost:27017", "typicode")
	assert.Nil(s.T(), err)
	s.db = db
}
