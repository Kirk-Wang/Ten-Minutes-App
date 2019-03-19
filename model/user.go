package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// The User holds
type User struct {
	ID       primitive.ObjectID `bson:"_id" json:"id"`
	Name     string             `bson:"name" json:"name"`
	UserName string             `bson:"username" json:"username"`
	Email    string             `bson:"email" json:"email"`
	Address  UserAddress        `bson:"address" json:"address"`
	Phone    string             `bson:"phone" json:"phone"`
	Website  string             `bson:"website" json:"website"`
	Company  UserCompany        `bson:"company" json:"company"`
	Created  time.Time          `bson:"created" json:"created"`
	Updated  time.Time          `bson:"updated" json:"updated"`
}

// The UserAddress holds
type UserAddress struct {
	Street  string         `bson:"street" json:"street"`
	Suite   string         `bson:"suite" json:"suite"`
	City    string         `bson:"city" json:"city"`
	Zipcode string         `bson:"zipcode" json:"zipcode"`
	Geo     UserAddressGeo `bson:"geo" json:"geo"`
}

// The UserAddressGeo holds
type UserAddressGeo struct {
	Lat string `bson:"lat" json:"lat"`
	Lng string `bson:"lng" json:"lng"`
}

// The UserCompany holds
type UserCompany struct {
	Name        string `bson:"name" json:"name"`
	CatchPhrase string `bson:"catchPhrase" json:"catchPhrase"`
	BS          string `bson:"bs" json:"bs"`
}

// New is
func (u *User) New() *User {
	return &User{
		ID:       primitive.NewObjectID(),
		Name:     u.Name,
		UserName: u.UserName,
		Email:    u.Email,
		Address:  u.Address,
		Phone:    u.Phone,
		Website:  u.Website,
		Company:  u.Company,
		Created:  time.Now(),
		Updated:  time.Now(),
	}
}
