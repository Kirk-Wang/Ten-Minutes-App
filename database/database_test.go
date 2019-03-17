package database

import (
	"context"
	"github.com/lotteryjs/ten-minutes-api/config"
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
	s.T().Log("--BeforeTest--")
	conf := config.Get()
	s.T().Log(conf.Database.Connection)
	s.T().Log(conf.Database.Dbname)
	db, _ := New(conf.Database.Connection, conf.Database.Dbname)
	assert.Nil(s.T(), nil)

	s.db = db
}

func (s *DatabaseSuite) AfterTest(suiteName, testName string) {
	s.db.Close()
}

func (s *DatabaseSuite) TestGetUsers() {
	s.T().Log("--TestGetUsers--")
	coll := s.db.DB.Collection("inventory_delete")

	err := coll.Drop(context.Background())
	assert.Nil(s.T(), err)
}
