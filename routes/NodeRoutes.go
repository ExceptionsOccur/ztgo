package routes

import (
	"ztgo/middlewares"
	"ztgo/services"

	"github.com/gin-gonic/gin"
)

func addZerotierNodeRoutes(rg *gin.RouterGroup) {
	routes := rg.Group("/node", middlewares.JWTAuth())
	routes.POST("/status", func(ctx *gin.Context) {
		services.GetNodeType(ctx)
	})
}
