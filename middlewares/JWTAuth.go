package middlewares

import (
	"strings"
	"ztgo/secure"
	"ztgo/utils"

	"github.com/gin-gonic/gin"
)

func JWTAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("token")
		if authHeader == "" {
			utils.ZTResponseEmptyAuth(ctx)
			ctx.Abort()
			return
		}
		// 按空格分割
		parts := strings.Split(authHeader, ".")
		if len(parts) != 3 {
			utils.ZTResponseAuthInvalid(ctx)
			ctx.Abort()
			return
		}
		claims, ok := secure.ParseToken(authHeader)
		if ok != nil {
			utils.ZTResponseAuthInvalid(ctx)
			ctx.Abort()
			return
		}
		// 将当前请求的username信息保存到请求的上下文c上
		ctx.Set("username", claims.Username)
		ctx.Next() // 后续的处理函数可以用过c.Get("username")来获取当前请求的用户信息
	}
}
