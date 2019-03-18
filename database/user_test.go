package database

import (
	"github.com/lotteryjs/ten-minutes-api/model"
	"github.com/stretchr/testify/assert"
)

func (s *DatabaseSuite) TestUser() {
	s.db.DB.Collection("users").Drop(nil)

	kirk := (&model.User{
		Name:     "Leanne Graham",
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
	assert.Equal(s.T(), kirk, s.db.GetUserByName("Leanne Graham"))

}
