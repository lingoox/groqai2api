package initialize

import (
	"github.com/gin-gonic/gin"
	"groqai2api/middlewares"
	"groqai2api/router"
)

func InitRouter() *gin.Engine {
	Router := gin.Default()

	Router.Use(middlewares.Cors)

	Router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, world!",
		})
	})

	Router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	AuthGroup := Router.Group("/auth/")
	router.InitAuth(AuthGroup)
	v1Group := Router.Group("/v1/")
	router.InitChat(v1Group)

	PlatformGroup := Router.Group("/platform/")
	router.InitPlatform(PlatformGroup)

	return Router
}
