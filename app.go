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

	db, err := database.New(conf.Database.Connection, conf.Database.Dbname)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	engine := router.Create(db, vInfo)
	runner.Run(engine, conf)
}
