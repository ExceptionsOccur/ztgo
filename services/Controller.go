package services

import (
	"encoding/json"
	"net/http"
	"ztgo/requests"
	"ztgo/utils"

	"github.com/gin-gonic/gin"
)

func GetAllControllerType(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, requests.GetAllControllerType())
}

func GetMemberType(ctx *gin.Context) {
	postData := utils.GetPostData[string, string](ctx)
	if _, ok := postData["nwid"]; !ok {
		ctx.JSON(http.StatusOK, map[string]string{
			"msg": "提交的数据不合法",
		})
		return
	}
	if _, ok := postData["mid"]; !ok {
		ctx.JSON(http.StatusOK, map[string]string{
			"msg": "提交的数据不合法",
		})
		return
	}
	ctx.JSON(http.StatusOK, requests.GetMemberType(postData["nwid"], postData["mid"]))
}

func GetAllMembersTypeByNwid(ctx *gin.Context) {
	postData := utils.GetPostData[string, string](ctx)
	if _, ok := postData["nwid"]; !ok {
		ctx.JSON(http.StatusOK, map[string]string{
			"msg": "提交的数据不合法",
		})
		return
	}
	ctx.JSON(http.StatusOK, requests.GetAllMembersTypeByNwid(postData["nwid"]))
}

func UpdateController(ctx *gin.Context) {
	postData := utils.GetPostData[string, interface{}](ctx)
	if _, ok := postData["nwid"]; !ok {
		ctx.JSON(http.StatusOK, map[string]string{
			"msg": "提交的数据不合法",
		})
		return
	}
	dataJson, err := json.Marshal(postData)
	if err != nil {
		ctx.JSON(http.StatusOK, map[string]string{
			"msg": "提交的数据不合法",
		})
		return
	}
	ctx.JSON(http.StatusOK, requests.UpdateController(string(dataJson)))
}

func CreateController(ctx *gin.Context) {
	postData := utils.GetPostData[string, string](ctx)
	name := ""
	if netName, ok := postData["name"]; ok {
		name = netName
	}
	ctx.JSON(http.StatusOK, requests.CreateController(name))
}

func DeleteController(ctx *gin.Context) {
	postData := utils.GetPostData[string, string](ctx)
	if _, ok := postData["nwid"]; !ok {
		ctx.JSON(http.StatusOK, map[string]string{
			"msg": "提交的数据不合法",
		})
		return
	}
	ctx.JSON(http.StatusOK, requests.DeleteController(postData["nwid"]))
}

func CountMembers(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, requests.CountMembers())
}

func UpdateMember(ctx *gin.Context) {
	postData := utils.GetPostData[string, interface{}](ctx)
	if _, ok := postData["nwid"]; !ok {
		ctx.JSON(http.StatusOK, map[string]string{
			"msg": "提交的数据不合法",
		})
		return
	}
	if _, ok := postData["id"]; !ok {
		ctx.JSON(http.StatusOK, map[string]string{
			"msg": "提交的数据不合法",
		})
		return
	}
	dataJson, err := json.Marshal(postData)
	if err != nil {
		ctx.JSON(http.StatusOK, map[string]string{
			"msg": "提交的数据不合法",
		})
		return
	}
	ctx.JSON(http.StatusOK, requests.UpdateMember(string(dataJson)))
}

func DeleteMember(ctx *gin.Context) {
	postData := utils.GetPostData[string, interface{}](ctx)
	if _, ok := postData["nwid"]; !ok {
		ctx.JSON(http.StatusOK, map[string]string{
			"msg": "提交的数据不合法",
		})
		return
	}
	if _, ok := postData["id"]; !ok {
		ctx.JSON(http.StatusOK, map[string]string{
			"msg": "提交的数据不合法",
		})
		return
	}
	_, err := json.Marshal(postData)
	if err != nil {
		ctx.JSON(http.StatusOK, map[string]string{
			"msg": "提交的数据不合法",
		})
		return
	}
	ctx.JSON(http.StatusOK, requests.DeleteMember(postData["nwid"].(string), postData["id"].(string)))
}
