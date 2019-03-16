package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/lotteryjs/ten-minutes-api/model"
)

// The UserDatabase interface for encapsulating database access.
type UserDatabase interface {
	CreateUser(user *model.User) error
	GetUserByName(name string) *model.User
}

// The UserAPI provides handlers for managing users.
type UserAPI struct {
	DB               UserDatabase
	PasswordStrength int
}

// CreateUser creates a user
// swagger:operation POST /user user createUser
//
// Create a user.
//
// ---
// consumes: [application/json]
// produces: [application/json]
// security: [clientTokenHeader: [], clientTokenQuery: [], basicAuth: []]
// parameters:
// - name: body
//   in: body
//   description: the user to add
//   required: true
//   schema:
//     $ref: "#/definitions/UserWithPass"
// responses:
//   200:
//     description: Ok
//     schema:
//         $ref: "#/definitions/User"
//   400:
//     description: Bad Request
//     schema:
//         $ref: "#/definitions/Error"
//   401:
//     description: Unauthorized
//     schema:
//         $ref: "#/definitions/Error"
//   403:
//     description: Forbidden
//     schema:
//         $ref: "#/definitions/Error"
func (a *UserAPI) CreateUser(ctx *gin.Context) {
	user := model.UserExternalWithPass{}
	if err := ctx.Bind(&user); err == nil {
		internal := a.toInternalUser(&user, []byte{})
		if a.DB.GetUserByName(internal.Name) == nil {
			a.DB.CreateUser(internal)
			// if err := a.UserChangeNotifier.fireUserAdded(internal.ID); err != nil {
			// 	ctx.AbortWithError(500, err)
			// 	return
			// }
			ctx.JSON(200, toExternalUser(internal))
		} else {
			ctx.AbortWithError(400, errors.New("username already exists"))
		}
	}
}

func (a *UserAPI) toInternalUser(response *model.UserExternalWithPass, pw []byte) *model.User {
	user := &model.User{
		Name:  response.Name,
		Admin: response.Admin,
	}
	if response.Pass != "" {
		user.Pass = password.CreatePassword(response.Pass, a.PasswordStrength)
	} else {
		user.Pass = pw
	}
	return user
}

func toExternalUser(internal *model.User) *model.UserExternal {
	return &model.UserExternal{
		Name:  internal.Name,
		Admin: internal.Admin,
		ID:    internal.ID,
	}
}
