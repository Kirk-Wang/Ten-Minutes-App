package router

import (
	"time"

	"github.com/gin-gonic/gin"
)

// Create creates the gin engine with all routes.
func Create(db *database.GormDatabase, vInfo *model.VersionInfo, conf *config.Configuration) (*gin.Engine, func()) {

}
