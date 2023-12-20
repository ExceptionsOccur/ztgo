package routes

import (
	"ztgo/services"

	"github.com/gin-gonic/gin"
)

func addWebsocketRoutes(rg *gin.RouterGroup) {
	rg.GET("/ws", func(ctx *gin.Context) {
		services.HandleWebSocketConnection(ctx)
	})
}
