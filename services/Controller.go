package services

import (
	"encoding/json"
	"ztgo/requests"
	"ztgo/utils"

	"github.com/gin-gonic/gin"
)

func GetAllControllerType(ctx *gin.Context) {
	utils.ZTResponseOK(ctx, requests.GetAllControllerType())
}

func GetMemberType(ctx *gin.Context) {
	postData := utils.GetPostData[string, string](ctx)
	if _, ok := postData["nwid"]; !ok {
		utils.ZTResponseDataError(ctx)
		return
	}
	if _, ok := postData["mid"]; !ok {
		utils.ZTResponseDataError(ctx)
		return
	}
	utils.ZTResponseOK(ctx, requests.GetMemberType(postData["nwid"], postData["mid"]))
}

func GetAllMembersTypeByNwid(ctx *gin.Context) {
	postData := utils.GetPostData[string, string](ctx)
	if _, ok := postData["nwid"]; !ok {
		utils.ZTResponseDataError(ctx)
		return
	}
	utils.ZTResponseOK(ctx, requests.GetAllMembersTypeByNwid(postData["nwid"]))
}

func UpdateController(ctx *gin.Context) {
	postData := utils.GetPostData[string, interface{}](ctx)
	if _, ok := postData["nwid"]; !ok {
		utils.ZTResponseDataError(ctx)
		return
	}
	dataJson, err := json.Marshal(postData)
	if err != nil {
		utils.ZTResponseDataError(ctx)
		return
	}
	utils.ZTResponseOK(ctx, requests.UpdateController(string(dataJson)))
}

func CreateController(ctx *gin.Context) {
	postData := utils.GetPostData[string, string](ctx)
	name := ""
	if netName, ok := postData["name"]; ok {
		name = netName
	}
	utils.ZTResponseOK(ctx, requests.CreateController(name))
}

func DeleteController(ctx *gin.Context) {
	postData := utils.GetPostData[string, string](ctx)
	if _, ok := postData["nwid"]; !ok {
		utils.ZTResponseDataError(ctx)
		return
	}
	utils.ZTResponseOK(ctx, requests.DeleteController(postData["nwid"]))
}

func CountMembers(ctx *gin.Context) {
	utils.ZTResponseOK(ctx, requests.CountMembers())
}

func UpdateMember(ctx *gin.Context) {
	postData := utils.GetPostData[string, interface{}](ctx)
	if _, ok := postData["nwid"]; !ok {
		utils.ZTResponseDataError(ctx)
		return
	}
	if _, ok := postData["id"]; !ok {
		utils.ZTResponseDataError(ctx)
		return
	}
	dataJson, err := json.Marshal(postData)
	if err != nil {
		utils.ZTResponseDataError(ctx)
		return
	}
	utils.ZTResponseOK(ctx, requests.UpdateMember(string(dataJson)))
}

func DeleteMember(ctx *gin.Context) {
	postData := utils.GetPostData[string, interface{}](ctx)
	if _, ok := postData["nwid"]; !ok {
		utils.ZTResponseDataError(ctx)
		return
	}
	if _, ok := postData["id"]; !ok {
		utils.ZTResponseDataError(ctx)
		return
	}
	_, err := json.Marshal(postData)
	if err != nil {
		utils.ZTResponseDataError(ctx)
		return
	}
	utils.ZTResponseOK(ctx, requests.DeleteMember(postData["nwid"].(string), postData["id"].(string)))
}
