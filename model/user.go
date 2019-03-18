package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// The User holds
type User struct {
	ID       primitive.ObjectID `bson:"_id"`
	Name     string             `bson:"name"`
	UserName string             `bson:"username"`
	Email    string             `bson:"email"`
	Address  UserAddress        `bson:"address"`
	Phone    string             `bson:"phone"`
	Website  string             `bson:"website"`
	Company  UserCompany        `bson:"company"`
	Created  time.Time          `bson:"created"`
	Updated  time.Time          `bson:"updated"`
}

// The UserAddress holds
type UserAddress struct {
	Street  string         `bson:"street"`
	Suite   string         `bson:"suite"`
	City    string         `bson:"city"`
	Zipcode string         `bson:"zipcode"`
	Geo     UserAddressGeo `bson:"geo"`
}

// The UserAddressGeo holds
type UserAddressGeo struct {
	Lat string `bson:"lat"`
	Lng string `bson:"lng"`
}

// The UserCompany holds
type UserCompany struct {
	Name        string `bson:"name"`
	CatchPhrase string `bson:"catchPhrase"`
	BS          string `bson:"bs"`
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
		Created:  u.Created,
		Updated:  u.Updated,
	}
}
