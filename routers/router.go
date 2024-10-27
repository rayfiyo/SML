package routers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/rayfiyo/simple-language-model/handles"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	// CORSの設定
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"GET", "POST"}
	config.AllowHeaders = []string{"Origin", "Content-Type"}
	r.Use(cors.New(config))

	r.POST("/message", handles.Message)

	return r
}
