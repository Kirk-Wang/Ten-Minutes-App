package api

import (
	"github.com/gin-gonic/gin"
	"github.com/lotteryjs/ten-minutes-api/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strconv"
)

// The PostDatabase interface for encapsulating database access.
type PostDatabase interface {
	GetPosts(paging *model.Paging) []*model.Post
	GetPostByID(id primitive.ObjectID) *model.Post
}

// The PostAPI provides handlers for managing users.
type PostAPI struct {
	DB PostDatabase
}

// GetPosts returns all the users
// _end=5&_order=DESC&_sort=id&_start=0 adapt react-admin
func (a *PostAPI) GetPosts(ctx *gin.Context) {
	var (
		start int64
		end   int64
		sort  string
		order int
	)

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

	users := a.DB.GetPosts(
		&model.Paging{
			Skip:    &start,
			Limit:   &end,
			SortKey: sort,
			SortVal: order,
		})

	ctx.Header("X-Total-Count", strconv.Itoa(len(users)))
	ctx.JSON(200, users)
}
