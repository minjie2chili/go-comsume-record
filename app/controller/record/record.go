package record

import (
	"github.com/gin-gonic/gin"
)

func RecordRouter(router *gin.RouterGroup) {
	router.GET("/", Record)
}

func Record(c *gin.Context) {
	c.JSON(200, 32);
}
