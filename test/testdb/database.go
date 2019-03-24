package testdb

import (
	"github.com/lotteryjs/ten-minutes-app/model"
	"testing"

	"github.com/lotteryjs/ten-minutes-app/database"
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

// NewUser creates a user and returns the user.
func (d *Database) NewUser(name string) *model.User {
	return d.NewUserWithName(name)
}

// NewUserWithName creates a user with a name and returns the user.
func (d *Database) NewUserWithName(name string) *model.User {
	user := (&model.User{
		Name:     name,
		UserName: "Bret",
		Email:    "Sincere@april.biz",
		Address: model.UserAddress{
			Street:  "Kulas Light",
			Suite:   "Apt. 556",
			City:    "Gwenborough",
			Zipcode: "92998-3874",
			Geo: model.UserAddressGeo{
				Lat: "-37.3159",
				Lng: "81.1496",
			},
		},
		Phone:   "1-770-736-8031 x56442",
		Website: "hildegard.org",
		Company: model.UserCompany{
			Name:        "Romaguera-Crona",
			CatchPhrase: "Multi-layered client-server neural-net",
			BS:          "harness real-time e-markets",
		},
	}).New()
	d.CreateUser(user)
	return user
}
