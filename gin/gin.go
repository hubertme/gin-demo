package gin

import (
	"github.com/gin-gonic/gin"
)

const LOCAL_ADDRESS = "localhost:3030"
var router = gin.Default()

// Group routes
var TestAccessGroup *gin.RouterGroup

func SetupGinServer() error {
	initGinHandlerGroups()

	return router.Run(LOCAL_ADDRESS)
}
