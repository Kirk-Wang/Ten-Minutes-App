package api

import (
	"github.com/gin-gonic/gin"
	"github.com/lotteryjs/ten-minutes-api/model"
	"strconv"
)

// The UserDatabase interface for encapsulating database access.
type UserDatabase interface {
	CreateUser(user *model.User) error
	GetUsers(paging *model.Paging) []*model.User
}

// The UserAPI provides handlers for managing users.
type UserAPI struct {
	DB UserDatabase
}

// GetUsers returns all the users
// _end=5&_order=DESC&_sort=id&_start=0 adapt react-admin
func (a *UserAPI) GetUsers(ctx *gin.Context) {
	start, _ := strconv.ParseInt(ctx.DefaultQuery("_start", "0"), 10, 64)
	end, _ := strconv.ParseInt(ctx.DefaultQuery("_end", "10"), 10, 64)
	sort := ctx.DefaultQuery("_sort", "_id")

	var order int
	order = 1
	if ctx.DefaultQuery("_order", "DESC") == "DESC" {
		order = -1
	}

	var paging = &model.Paging{
		Skip:    &start,
		Limit:   &end,
		SortKey: sort,
		SortVal: order,
	}

	ctx.JSON(200, a.DB.GetUsers(paging))
}
