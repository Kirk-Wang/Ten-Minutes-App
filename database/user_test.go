package database

import (
	"github.com/lotteryjs/ten-minutes-app/model"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *DatabaseSuite) TestCreateUser() {
	s.db.DB.Collection("users").Drop(nil)

	kirk := (&model.User{
		Name:     "Graham",
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
	err := s.db.CreateUser(kirk)
	assert.Nil(s.T(), err)
}

func (s *DatabaseSuite) TestGetUsers() {
	start := int64(0)
	limit := int64(10)
	sort := "_id"
	order := -1

	users := s.db.GetUsers(&model.Paging{
		Skip:      &start,
		Limit:     &limit,
		SortKey:   sort,
		SortVal:   order,
		Condition: nil,
	})

	assert.Len(s.T(), users, 1)
}

func (s *DatabaseSuite) TestGetUserByName() {
	user := s.db.GetUserByName("Graham")

	assert.Equal(s.T(), user.Name, "Graham")
}

func (s *DatabaseSuite) TestGetUserByIDs() {
	user := s.db.GetUserByName("Graham")
	objectIds := []primitive.ObjectID{user.ID}
	users := s.db.GetUserByIDs(objectIds)

	assert.Len(s.T(), users, 1)
}

func (s *DatabaseSuite) TestCountUser() {
	len := s.db.CountUser()
	assert.Equal(s.T(), len, "1")
}
