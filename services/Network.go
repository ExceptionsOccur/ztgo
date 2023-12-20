package services

import (
	"net/http"
	"ztgo/requests"
	"ztgo/utils"

	"github.com/gin-gonic/gin"
)

func GetAllNetworkType(ctx *gin.Context) {
	allNetworkType := requests.GetAllNetworkType()
	ctx.JSON(http.StatusOK, allNetworkType)
}
func JoinToNetwork(ctx *gin.Context) {
	postData := utils.GetPostData[string, string](ctx)

	if nwid, ok := postData["nwid"]; ok {
		ctx.JSON(http.StatusOK, requests.JoinToNetwork(nwid))
		return
	}
	ctx.JSON(http.StatusOK, map[string]string{
		"msg": "提交的数据不合法",
	})

}
func LeaveNetwork(ctx *gin.Context) {
	postData := utils.GetPostData[string, string](ctx)
	if nwid, ok := postData["nwid"]; ok {
		ctx.JSON(http.StatusOK, requests.LeaveNetwork(nwid))
		return
	}
	ctx.JSON(http.StatusOK, map[string]string{
		"msg": "提交的数据不合法",
	})
}
