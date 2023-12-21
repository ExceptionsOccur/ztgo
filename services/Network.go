package services

import (
	"ztgo/requests"
	"ztgo/utils"

	"github.com/gin-gonic/gin"
)

func GetAllNetworkType(ctx *gin.Context) {
	utils.ZTResponseOK(ctx, requests.GetAllNetworkType())
}
func JoinToNetwork(ctx *gin.Context) {
	postData := utils.GetPostData[string, string](ctx)

	if nwid, ok := postData["nwid"]; ok {
		utils.ZTResponseOK(ctx, requests.JoinToNetwork(nwid))
		return
	}
	utils.ZTResponseDataError(ctx)

}
func LeaveNetwork(ctx *gin.Context) {
	postData := utils.GetPostData[string, string](ctx)
	if _, ok := postData["nwid"]; !ok {
		utils.ZTResponseDataError(ctx)
		return
	}
	if _, ok := postData["id"]; !ok {
		utils.ZTResponseDataError(ctx)
		return
	}
	utils.ZTResponseOK(ctx, requests.LeaveNetwork(postData["nwid"]))
	requests.DeleteMember(postData["nwid"], postData["id"])
}
