package routes

import (
	"ztgo/middlewares"

	"github.com/gin-gonic/gin"
)

func Run() {
	// gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.Use(middlewares.Cors())
	v1 := router.Group("/zerotier")
	addUserRoutes(v1)
	addZerotierNodeRoutes(v1)
	addZerotierNetworkRoutes(v1)
	addZerotierControllerRoutes(v1)
	addSystemRoutes(v1)
	router.Run(":8000")
}
