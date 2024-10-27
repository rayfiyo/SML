package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rayfiyo/simple-language-model/data"
	"github.com/rayfiyo/simple-language-model/models"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/query", postQuery)
	r.POST("/reply", postReply)

	return r
}

func postQuery(c *gin.Context) {
	var newQuery models.Request

	if err := c.BindJSON(&newQuery); err != nil {
		c.IndentedJSON(http.StatusBadRequest, newQuery)
		return
	}

	if newQuery.Mode != "query" {
		c.IndentedJSON(http.StatusBadRequest, newQuery)
		return
	}

	data.Query = append(data.Query, newQuery.Content)
	c.IndentedJSON(http.StatusCreated, newQuery)
}

func postReply(c *gin.Context) {
	var newReply models.Request

	if err := c.BindJSON(&newReply); err != nil {
		c.IndentedJSON(http.StatusBadRequest, newReply)
		return
	}

	if newReply.Mode != "reply" {
		c.IndentedJSON(http.StatusBadRequest, newReply)
		return
	}

	data.Reply = append(data.Reply, newReply.Content)
	c.IndentedJSON(http.StatusCreated, newReply)
}
