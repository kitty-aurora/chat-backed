package router

import (
	"chat-backed/controller"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// 健康检查
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	// 用户接口
	r.GET("/user/:id", controller.GetUserByID)

	return r
}
