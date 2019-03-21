package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/lotteryjs/ten-minutes-api/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strconv"
)

// The PostDatabase interface for encapsulating database access.
type PostDatabase interface {
	GetPosts(paging *model.Paging) []*model.Post
	GetPostByID(id primitive.ObjectID) *model.Post
	CreatePost(post *model.Post) *model.Post
	UpdatePost(post *model.Post) *model.Post
}

// The PostAPI provides handlers for managing posts.
type PostAPI struct {
	DB PostDatabase
}

// CreatePost creates a post.
func (a *PostAPI) CreatePost(ctx *gin.Context) {
	var post = model.Post{}
	if err := ctx.ShouldBind(&post); err == nil {
		if result := a.DB.CreatePost(&post); result != nil {
			ctx.JSON(200, result)
		} else {
			ctx.AbortWithError(500, errors.New("CreatePost error"))
		}
	} else {
		ctx.AbortWithError(500, errors.New("ShouldBind error"))
	}
}

// GetPosts returns all the posts
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

	posts := a.DB.GetPosts(
		&model.Paging{
			Skip:    &start,
			Limit:   &end,
			SortKey: sort,
			SortVal: order,
		})

	ctx.Header("X-Total-Count", strconv.Itoa(len(posts)))
	ctx.JSON(200, posts)
}

// GetPostByID returns the post by id
func (a *PostAPI) GetPostByID(ctx *gin.Context) {
	withID(ctx, "id", func(id primitive.ObjectID) {
		if post := a.DB.GetPostByID(id); post != nil {
			ctx.JSON(200, post)
		} else {
			ctx.AbortWithError(404, errors.New("post does not exist"))
		}
	})
}

// UpdatePostByID is
func (a *PostAPI) UpdatePostByID(ctx *gin.Context) {
	withID(ctx, "id", func(id primitive.ObjectID) {
		var post = model.Post{}
		if err := ctx.ShouldBind(&post); err == nil {
			if result := a.DB.UpdatePost(&post); result != nil {
				ctx.JSON(200, result)
			} else {
				ctx.AbortWithError(404, errors.New("post does not exist"))
			}
		} else {
			ctx.AbortWithError(404, errors.New("post does not exist"))
		}
	})
}
