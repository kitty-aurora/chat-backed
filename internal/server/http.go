package server

import (
	"chat-backed/internal/handler"
	"github.com/gin-gonic/gin"
)

func NewHTTPServer() *gin.Engine {
	r := gin.Default()

	// 注册路由
	r.GET("/ws", handler.ServeWs)

	return r
}
