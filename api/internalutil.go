package api

import (
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/gin-gonic/gin"
)

func withID(ctx *gin.Context, name string, f func(id primitive.ObjectID)) {
	if id, err := primitive.ObjectIDFromHex(ctx.Param(name)); err == nil {
		f(id)
	} else {
		ctx.AbortWithError(400, errors.New("invalid id"))
	}
}

func withIDs(ctx *gin.Context, name string, f func(id []primitive.ObjectID)) {
	ids, b := ctx.GetQueryArray(name)
	var objectIds []primitive.ObjectID
	abort := errors.New("invalid id")
	if b {
		for index, id := range ids {
			if objID, err := primitive.ObjectIDFromHex(id); err == nil {
				objectIds[index] = objID
			} else {
				ctx.AbortWithError(400, abort)
			}
		}
		f(objectIds)
	} else {
		ctx.AbortWithError(400, abort)
	}
}
