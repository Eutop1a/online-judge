package middlewares

import (
	"github.com/gin-gonic/gin"
	"online-judge/pkg/jwt"
	"online-judge/pkg/resp"
	"strings"
)

// JWTAuthMiddleware 基于JWT的认证中间件
func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		// 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
		// 这里假设Token放在Header的Authorization中，并使用Bearer开头
		// Authorization: Bearer xxxxxx.xxx.xxx / x-TOKEN xxx.xx.xx
		// 这里的具体实现方式要依据你的实际业务情况决定
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			resp.ResponseError(c, resp.CodeNeedLogin)
			c.Abort()
			return
		}
		// 按空格分割
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			resp.ResponseError(c, resp.CodeInvalidToken)
			c.Abort()
			return
		}
		// parts[1]是获取到的tokenString，我们使用之前定义好的解析JWT的函数来解析它
		mc, err := jwt.ParseToken(parts[1])
		if err != nil {
			resp.ResponseError(c, resp.CodeInvalidToken)
			c.Abort()
			return
		}
		// 将当前请求的useID信息和username保存到请求的上下文c上
		c.Set(resp.CtxUserIDKey, mc.UserID)
		c.Set(resp.CtxUserNameKey, mc.Username)
		c.Next() // 后续的处理请求的函数可以用过c.Get(CtxUserIDKey)来获取当前请求的用户信息
	}
}
