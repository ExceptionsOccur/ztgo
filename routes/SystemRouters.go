package routes

import (
	"ztgo/services"

	"github.com/gin-gonic/gin"
)

func addSystemRoutes(rg *gin.RouterGroup) {
	routes := rg.Group("/sys")
	routes.POST("/getSystemInfo", func(ctx *gin.Context) {
		services.GetSystemInfo(ctx)
	})
}
