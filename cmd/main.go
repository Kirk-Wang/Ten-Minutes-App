package main

import (
	"fmt"
	_ "github.com/lotteryjs/Ten-Minutes-App"
	"github.com/lotteryjs/Ten-Minutes-App/infrastructure"
)

const logo = `
___ ____ _  _    _  _ _ _  _ _  _ ___ ____ ____    ____ ___  ___  
|  |___ |\ | __ |\/| | |\ | |  |  |  |___ [__  __ |__| |__] |__] 
|  |___ | \|    |  | | | \| |__|  |  |___ ___]    |  | |    |    																				
`

func main() {

	fmt.Println(logo)

	app := infrastructure.New()
	app.Start()
}
