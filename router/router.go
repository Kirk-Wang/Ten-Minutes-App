package router

import (
	"github.com/gin-gonic/gin"

	"github.com/lotteryjs/ten-minutes-api/api"
	"github.com/lotteryjs/ten-minutes-api/database"
	"github.com/lotteryjs/ten-minutes-api/error"
	"github.com/lotteryjs/ten-minutes-api/model"
)

// Create creates the gin engine with all routes.
func Create(db *database.TenDatabase, vInfo *model.VersionInfo) *gin.Engine {
	g := gin.New()

	g.Use(gin.Logger(), gin.Recovery(), error.Handler())
	g.NoRoute(error.NotFound())

	userHandler := api.UserAPI{DB: db}

	authAdmin := g.Group("/user")
	{
		authAdmin.POST("", userHandler.CreateUser)
	}

	g.GET("version", func(ctx *gin.Context) {
		ctx.JSON(200, vInfo)
	})

	return g
}
