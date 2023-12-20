package middlewares

import (
	"net/http"
	"strings"
	"ztgo/secure"

	"github.com/gin-gonic/gin"
)

func JWTAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("token")
		if authHeader == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code": 2003,
				"msg":  "请求头中auth为空",
			})
			ctx.Abort()
			return
		}
		// 按空格分割
		parts := strings.Split(authHeader, ".")
		if len(parts) != 3 {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code": 2004,
				"msg":  "请求头中auth格式有误",
			})
			ctx.Abort()
			return
		}
		claims, ok := secure.ParseToken(authHeader)
		if ok != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code": 2005,
				"msg":  "无效的token",
			})
			ctx.Abort()
			return
		}
		// ctx.JSON(http.StatusUnauthorized, gin.H{
		// 	"code": 2000,
		// 	"msg":  "验证完成",
		// })
		// 将当前请求的username信息保存到请求的上下文c上
		ctx.Set("username", claims.Username)
		ctx.Next() // 后续的处理函数可以用过c.Get("username")来获取当前请求的用户信息
	}
}
