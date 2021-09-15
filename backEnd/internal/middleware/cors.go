package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	return func(context *gin.Context) {
		if len(context.Request.URL.Path) < 7 {
			return
		}
		v := context.Request.URL.Path[5:7]
		if v == "v1" {
			origin := context.Request.Header.Get("Origin")
			context.Header("Access-Control-Allow-Origin", origin)
			context.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
			context.Header("Access-Control-Allow-Methods", "POST, UPDATE, DELETE ,GET, OPTIONS")
			context.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access- Control-Allow-Headers, Content-Type")
			context.Header("Access-Control-Allow-Credentials", "true")
		} else if v == "v2" {
			context.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
			context.Header("Access-Control-Allow-Methods", "POST, UPDATE, DELETE ,GET, OPTIONS")
			context.Header("Access-Control-Allow-Credentials", "true")
		}

		method := context.Request.Method
		if method == "OPTIONS" {
			if v == "v2" {
				context.Header("Access-Control-Allow-Origin", "*")
			}
			context.AbortWithStatus(http.StatusNoContent)
		}
		context.Next()
	}
}
