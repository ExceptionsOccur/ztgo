package routes

import (
	"ztgo/middlewares"
	"ztgo/services"

	"github.com/gin-gonic/gin"
)

func addZerotierControllerRoutes(rg *gin.RouterGroup) {
	routes := rg.Group("/controller", middlewares.JWTAuth())
	routes.POST("/getAllControllerType", func(ctx *gin.Context) {
		services.GetAllControllerType(ctx)
	})
	routes.POST("/getMemberType", func(ctx *gin.Context) {
		services.GetMemberType(ctx)
	})
	routes.POST("getAllMembersTypeByNwid", func(ctx *gin.Context) {
		services.GetAllMembersTypeByNwid(ctx)
	})
	routes.POST("updateController", func(ctx *gin.Context) {
		services.UpdateController(ctx)
	})
	routes.POST("createController", func(ctx *gin.Context) {
		services.CreateController(ctx)
	})
	routes.POST("deleteController", func(ctx *gin.Context) {
		services.DeleteController(ctx)
	})
	routes.POST("updateMember", func(ctx *gin.Context) {
		services.UpdateMember(ctx)
	})
	routes.POST("deleteMember", func(ctx *gin.Context) {
		services.DeleteMember(ctx)
	})
	routes.POST("countMembers", func(ctx *gin.Context) {
		services.CountMembers(ctx)
	})
}
