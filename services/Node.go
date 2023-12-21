package services

import (
	"ztgo/requests"
	"ztgo/utils"

	"github.com/gin-gonic/gin"
)

func GetNodeType(ctx *gin.Context) {
	utils.ZTResponseOK(ctx, requests.GetNodeType())
}
