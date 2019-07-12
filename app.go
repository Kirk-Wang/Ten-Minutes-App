package app

import (
	"github.com/lotteryjs/Ten-Minutes-App/infrastructure"
	"github.com/lotteryjs/Ten-Minutes-App/infrastructure/engines"
)

func init() {
	infrastructure.Register(&engines.GinEngine{})
}