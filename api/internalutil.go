package api

import (
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/gin-gonic/gin"
)

func withID(ctx *gin.Context, name string, f func(id primitive.ObjectID)) {
	if id, err := primitive.ObjectIDFromHex(name); err == nil {
		f(id)
	} else {
		ctx.AbortWithError(400, errors.New("invalid id"))
	}
}
