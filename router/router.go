package router

import (
	"github.com/gin-gonic/gin"
	"net/http"

	"github.com/lotteryjs/ten-minutes-api/api"
	"github.com/lotteryjs/ten-minutes-api/config"
	"github.com/lotteryjs/ten-minutes-api/database"
	"github.com/lotteryjs/ten-minutes-api/error"
	"github.com/lotteryjs/ten-minutes-api/model"
)

// Create creates the gin engine with all routes.
func Create(db *database.TenDatabase, vInfo *model.VersionInfo, conf *config.Configuration) *gin.Engine {
	g := gin.New()

	g.Use(gin.Logger(), gin.Recovery(), error.Handler())
	g.NoRoute(error.NotFound())

	g.Use(func(ctx *gin.Context) {
		ctx.Header("Content-Type", "application/json")
		for header, value := range conf.Server.ResponseHeaders {
			ctx.Header(header, value)
		}
		if ctx.Request.Method == "OPTIONS" {
			ctx.AbortWithStatus(http.StatusNoContent)
		}
	})

	userHandler := api.UserAPI{DB: db}
	postHandler := api.PostAPI{DB: db}

	g.GET("/users", userHandler.GetUsers)

	postG := g.Group("/posts")
	{
		postG.GET("", postHandler.GetPosts)
		postG.GET(":id", postHandler.GetPostByID)
	}

	g.GET("version", func(ctx *gin.Context) {
		ctx.JSON(200, vInfo)
	})

	return g
}
