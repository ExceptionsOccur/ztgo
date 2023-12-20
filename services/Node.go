package services

import (
	"net/http"
	"ztgo/requests"

	"github.com/gin-gonic/gin"
)

func GetNodeType(ctx *gin.Context) {
	nodeType := requests.GetNodeType()
	ctx.JSON(http.StatusOK, nodeType)
}
