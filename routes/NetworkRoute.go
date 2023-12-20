package routes

import (
	"ztgo/middlewares"
	"ztgo/services"

	"github.com/gin-gonic/gin"
)

func addZerotierNetworkRoutes(rg *gin.RouterGroup) {
	routes := rg.Group("/network", middlewares.JWTAuth())
	routes.POST("/getAllNetworkType", func(ctx *gin.Context) {
		services.GetAllNetworkType(ctx)
	})
	routes.POST("/joinToNetwork", func(ctx *gin.Context) {
		services.JoinToNetwork(ctx)
	})
	routes.POST("/leaveNetwork", func(ctx *gin.Context) {
		services.LeaveNetwork(ctx)
	})
}
