package main

import (
	"fmt"

	"github.com/spf13/pflag"

	_ "github.com/lotteryjs/Ten-Minutes-App"
	"github.com/lotteryjs/Ten-Minutes-App/infrastructure"
	"github.com/lotteryjs/Ten-Minutes-App/infrastructure/config"
)

var (
	cfg = pflag.StringP("config", "c", "", "apiserver config file path.")
)

const logo = `
___ ____ _  _    _  _ _ _  _ _  _ ___ ____ ____    ____ ___  ___  
|  |___ |\ | __ |\/| | |\ | |  |  |  |___ [__  __ |__| |__] |__] 
|  |___ | \|    |  | | | \| |__|  |  |___ ___]    |  | |    |    																				
`

func init() {
	fmt.Println(logo)
}

func main() {
	pflag.Parse()

	// init config
	if err := config.Init(*cfg); err != nil {
		panic(err)
	}

	app := infrastructure.New()
	app.Start()
}
