package api

import (
	// "errors"
	"github.com/gin-gonic/gin"
	// "github.com/lotteryjs/ten-minutes-api/auth/password"
	"github.com/lotteryjs/ten-minutes-api/model"
)

// The UserDatabase interface for encapsulating database access.
type UserDatabase interface {
	GetUserByName(name string) *model.User
}

// The UserAPI provides handlers for managing users.
type UserAPI struct {
	DB               UserDatabase
	PasswordStrength int
}
