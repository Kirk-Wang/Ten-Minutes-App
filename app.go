package app

import (
	"github.com/lotteryjs/Ten-Minutes-App/infrastructure/engines"
)

func init() {
	engines.Register(&engines.GinEngine{})
}