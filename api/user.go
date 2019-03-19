package api

import (
	"github.com/gin-gonic/gin"
	"github.com/lotteryjs/ten-minutes-api/model"
)

// The UserDatabase interface for encapsulating database access.
type UserDatabase interface {
	CreateUser(user *model.User) error
	GetUsers() []*model.User
}

// The UserAPI provides handlers for managing users.
type UserAPI struct {
	DB UserDatabase
}

// GetUsers returns all the users
// _end=5&_order=DESC&_sort=id&_start=0

func (a *UserAPI) GetUsers(ctx *gin.Context) {
	var (
		skip  int64 = 0
		limit int64 = 5
	)
	var sortKey string = "_id"
	var sortVal bool = true
	// skip := ctx.DefaultQuery("_start", "0")
	// limit := ctx.DefaultQuery("_end", "10")
	// sortKey := ctx.DefaultQuery("_sort", "_id")
	// sortVal := ctx.DefaultQuery("_order", "-1")
	// ctx.JSON(200, a.DB.GetUsers(skip, limit, sortKey, sortVal))
	ctx.JSON(200, a.DB.GetUsers(skip, limit, "_id", true))
}
