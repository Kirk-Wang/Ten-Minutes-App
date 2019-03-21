package api

import (
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/gin-gonic/gin"
)

func withID(ctx *gin.Context, name string, f func(id primitive.ObjectID)) {
	value := ctx.DefaultQuery(name, "")
	if value == "" {
		value = ctx.Param(name)
	}
	if id, err := primitive.ObjectIDFromHex(value); err == nil {
		f(id)
	} else {
		ctx.AbortWithError(400, errors.New("invalid id"))
	}
}
