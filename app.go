package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/lotteryjs/ten-minutes-api/config"
	"github.com/lotteryjs/ten-minutes-api/mode"
	"github.com/lotteryjs/ten-minutes-api/model"
	"github.com/lotteryjs/ten-minutes-api/router"
	"github.com/lotteryjs/ten-minutes-api/runner"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	// Version the version of TMAPI.
	Version = "unknown"
	// Commit the git commit hash of this version.
	Commit = "unknown"
	// BuildDate the date on which this binary was build.
	BuildDate = "unknown"
	// Mode the build mode
	Mode = mode.Dev
)

func main() {
	vInfo := &model.VersionInfo{Version: Version, Commit: Commit, BuildDate: BuildDate}
	mode.Set(Mode)

	fmt.Println("Starting TMAPI version", vInfo.Version+"@"+BuildDate)
	rand.Seed(time.Now().UnixNano())
	conf := config.Get()

	// connect to mongodb
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(conf.Database.Connection))
	if err != nil {
		panic(err)
	}
	defer client.Disconnect(ctx)

	// ping mongod
	ctx2, cancel2 := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel2()
	err = client.Ping(ctx2, readpref.Primary())
	if err != nil {
		panic(err)
	}

	db := client.Database(conf.Database.Dbname)
	engine := router.Create(db, vInfo)
	runner.Run(engine, conf)
}
