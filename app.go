package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/lotteryjs/ten-minutes-api/config"
	"github.com/lotteryjs/ten-minutes-api/mode"
	"github.com/lotteryjs/ten-minutes-api/model"
)

var (
	// Version the version of API.
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

	fmt.Println("Starting API version", vInfo.Version+"@"+BuildDate)
	rand.Seed(time.Now().UnixNano())
	conf := config.Get()

	db, err := database.New(conf.Database.Dialect, conf.Database.Connection)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	engine, closeable := router.Create(db, vInfo, conf)
	defer closeable()

	runner.Run(engine, conf)
}
