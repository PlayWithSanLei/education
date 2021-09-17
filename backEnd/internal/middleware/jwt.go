package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/impact-eintr/education/pkg/jwt"
	"github.com/impact-eintr/education/pkg/resp"
)

// JWT认证中间件 检查缓存中的权限信息
func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
		// 这里假设Token放在Header的Authorization中，并使用Bearer开头
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			resp.ResponseError(c, resp.CodeNullAuth)
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
		// 将当前请求的username信息保存到请求的上下文c上
		c.Set("userID", mc.UserID)
		c.Set("userRole", mc.UserRole)
		c.Set("userUnit", mc.UserUnit)
		c.Next() // 后续的处理函数可以用过c.Get("username")来获取当前请求的用户信息
	}
}
