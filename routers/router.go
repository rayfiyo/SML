package routers

import (
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/query", query)
	r.GET("/reply", reply)

	return r
}

func query(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "query ok",
	})
}

func reply(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "reply ok",
	})
}
