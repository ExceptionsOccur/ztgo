package utils

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 正常响应，网络响应码200,
// 返回客户端的数据{code: 2000, data: data}
func ZTResponseOK(ctx *gin.Context, data interface{}) {
	res := make(map[string]interface{})
	res["code"] = 2000
	res["data"] = data
	ctx.JSON(http.StatusOK, res)
}

// 错误响应，网络响应码200,
// 返回客户端的数据{code: 2100, data: data}
func ZTResponseError(ctx *gin.Context, data interface{}) {
	res := make(map[string]interface{})
	res["code"] = 2100
	res["data"] = data
	ctx.JSON(http.StatusOK, res)
}

// 数据错误响应，网络响应码200,
// 返回客户端的数据{code: 2001, data: "提交的数据不合法"}
func ZTResponseDataError(ctx *gin.Context) {
	res := make(map[string]interface{})
	res["code"] = 2001
	res["data"] = "提交的数据不合法"
	ctx.JSON(http.StatusOK, res)
}

// 空token响应，网络响应码200,
// 返回客户端的数据{code: 2003, data: "token为空"}
func ZTResponseEmptyAuth(ctx *gin.Context) {
	res := make(map[string]interface{})
	res["code"] = 2003
	res["data"] = "token为空"
	ctx.JSON(http.StatusOK, res)
}

// token无效响应，网络响应码200,
// 返回客户端的数据{code: 2004, data: "token无效"}
func ZTResponseAuthInvalid(ctx *gin.Context) {
	res := make(map[string]interface{})
	res["code"] = 2004
	res["data"] = "token无效"
	ctx.JSON(http.StatusOK, res)
}

// token验证通过响应，网络响应码200,
// 返回客户端的数据{code: 4000, data: "验证通过", token: token}
func ZTResponseAuthorizedOK(ctx *gin.Context, token string) {
	res := make(map[string]interface{})
	res["code"] = 4000
	res["data"] = "验证通过"
	res["token"] = token
	ctx.JSON(http.StatusOK, res)
}

// 未授权响应，网络响应码401,
// 返回客户端的数据{code: 4001, data: "访问未授权"}
func ZTResponseUnauthorized(ctx *gin.Context) {
	res := make(map[string]interface{})
	res["code"] = 4001
	res["data"] = "访问未授权"
	ctx.JSON(http.StatusUnauthorized, res)
}

// 构造新成员加入时websocket返回的数据,
// 返回客户端的数据{code: 6000, data: "检测到新成员加入", controller: controller_id, member: member_id}
func ZTResponseNewMemberJoins(cid string, mid string) []byte {
	res := make(map[string]interface{})
	res["code"] = 6000
	res["data"] = "检测到新成员加入"
	res["controller"] = cid
	res["member"] = mid
	bytes, _ := json.Marshal(res)
	return bytes
}

// 构造成员退出时websocket返回的数据,
// 返回客户端的数据{code: 6001, data: "检测到成员退出", controller: controller_id, member: member_id}
func ZTResponseMemberLeaves(cid string, mid string) []byte {
	res := make(map[string]interface{})
	res["code"] = 6001
	res["data"] = "检测到成员退出"
	res["controller"] = cid
	res["member"] = mid
	bytes, _ := json.Marshal(res)
	return bytes
}
