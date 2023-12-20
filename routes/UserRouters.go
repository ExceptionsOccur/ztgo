package routes

import (
	"ztgo/services"

	"github.com/gin-gonic/gin"
)

func addUserRoutes(rg *gin.RouterGroup) {
	routes := rg.Group("/user")
	routes.POST("/valid", func(ctx *gin.Context) {
		services.EnsureSecretFileExist()
		services.ValidateTOTPCode(ctx)
	})
}
