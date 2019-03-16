package router

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/lotteryjs/ten-minutes-api/error"
	"github.com/lotteryjs/ten-minutes-api/model"
)

// Create creates the gin engine with all routes.
func Create(db *mongo.Database, vInfo *model.VersionInfo) (*gin.Engine, func()) {
	g := gin.New()

	g.Use(gin.Logger(), gin.Recovery(), error.Handler(), location.Default())
	g.NoRoute(error.NotFound())

}