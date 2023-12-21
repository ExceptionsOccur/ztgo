package routes

import (
	"time"
	"ztgo/middlewares"
	"ztgo/services"

	"github.com/gin-gonic/gin"
)

func addUserRoutes(rg *gin.RouterGroup) {
	// 访问限制，桶令牌，最大3个，每秒1个，即初始最大并大3次，用完后每秒最多1次
	routes := rg.Group("/user", middlewares.RateLimitMiddleware(time.Second, 3, 1))
	routes.POST("/valid", func(ctx *gin.Context) {
		services.EnsureSecretFileExist()
		services.ValidateTOTPCode(ctx)
	})
}
