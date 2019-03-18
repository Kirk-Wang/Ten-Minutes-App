package testdb

import (
	"testing"

	"github.com/lotteryjs/ten-minutes-api/database"
	"github.com/stretchr/testify/assert"
)

// Database is the wrapper for the gorm database with sleek helper methods.
type Database struct {
	*database.TenDatabase
	t *testing.T
}

// NewDB creates a new test db instance.
func NewDB(t *testing.T) *Database {
	db, err := database.New("mongodb://root:123456@localhost:27017", "tenapi")
	assert.Nil(t, err)
	assert.NotNil(t, db)
	return &Database{TenDatabase: db, t: t}
}
