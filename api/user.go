package api

import (
	"github.com/gin-gonic/gin"
	"github.com/lotteryjs/ten-minutes-api/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strconv"
)

// The UserDatabase interface for encapsulating database access.
type UserDatabase interface {
	GetUserByIDs(ids []primitive.ObjectID) []*model.User
	CreateUser(user *model.User) error
	GetUsers(paging *model.Paging) []*model.User
	CountUser() string
}

// The UserAPI provides handlers for managing users.
type UserAPI struct {
	DB UserDatabase
}

// GetUserByIDs returns the user by id
func (a *UserAPI) GetUserByIDs(ctx *gin.Context) {
	withIDs(ctx, "id", func(ids []primitive.ObjectID) {
		ctx.JSON(200, a.DB.GetUserByIDs(ids))
	})
}

// GetUsers returns all the users
// _end=5&_order=DESC&_sort=id&_start=0 adapt react-admin
func (a *UserAPI) GetUsers(ctx *gin.Context) {
	var (
		start int64
		end   int64
		sort  string
		order int
	)
	id := ctx.DefaultQuery("id", "")
	if id != "" {
		a.GetUserByIDs(ctx)
		return
	}
	start, _ = strconv.ParseInt(ctx.DefaultQuery("_start", "0"), 10, 64)
	end, _ = strconv.ParseInt(ctx.DefaultQuery("_end", "10"), 10, 64)
	sort = ctx.DefaultQuery("_sort", "_id")
	order = 1

	if sort == "id" {
		sort = "_id"
	}

	if ctx.DefaultQuery("_order", "DESC") == "DESC" {
		order = -1
	}

	limit := end - start
	users := a.DB.GetUsers(
		&model.Paging{
			Skip:    &start,
			Limit:   &limit,
			SortKey: sort,
			SortVal: order,
		})

	ctx.Header("X-Total-Count", a.DB.CountUser())
	ctx.JSON(200, users)
}
