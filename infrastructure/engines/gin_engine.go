package engines

import (
	"github.com/gin-gonic/gin"
)

var ginEngine *gin.Engine

func Gin() *gin.Engine {
	return ginEngine
}

type GinEngine struct {
	Base
}

func (*GinEngine) Init(ctx EngineContext) {
	ginEngine = initGin()
}

func (*GinEngine) Start(ctx EngineContext) {
	Gin().Run(":8080")
}

func (*GinEngine) StartBlocking() bool {
	return true
}

func initGin() *gin.Engine {
	// Creates a router without any middleware by default
	g := gin.New()

	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	g.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	g.Use(gin.Recovery())

	// Per route middleware, you can add as many as you desire.
	g.GET("/benchmark", MyBenchLogger(), benchEndpoint)

	return g
}

