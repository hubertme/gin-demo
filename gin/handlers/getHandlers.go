package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/hubertme/gin-demo/result"
	"net/http"
)

func PingHandler(c *gin.Context) {
	c.JSON(http.StatusAccepted, result.Success("This is lit!"))
}

func PongHandler(c *gin.Context) {
	dataRes := map[string]interface{} {
		"num": 69,
		"note": "This is a dummy note",
	}

	c.JSON(http.StatusAccepted, result.Success(dataRes))
}

func DevErrorHandler(c *gin.Context) {
	dataRes := map[string]interface{} {

	}

	c.JSON(http.StatusInternalServerError, result.DevError(dataRes))
}
