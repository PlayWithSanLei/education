package middleware

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/impact-eintr/education/global"
	"github.com/impact-eintr/education/internal/dao/cache"
	"github.com/impact-eintr/education/internal/dao/db"
	"github.com/impact-eintr/education/internal/model"
	"github.com/impact-eintr/education/pkg/resp"
)

func RBACMiddleware(p string) gin.HandlerFunc {
	return func(c *gin.Context) {
		oldRole := c.GetString("userRole")
		bRole := cache.CacheCheck(fmt.Sprintf("%d", c.GetInt64("userID")))
		if len(bRole) == 0 {
			user := new(model.User)
			db.UserData(c.GetInt64("userID"), user)
			if user.UserID != 0 {
				cache.CacheUpdate(fmt.Sprintf("%d", user.UserID), user.Role)
			}
		}
		curRole := string(bRole)
		// 先检测当前用户的角色是否已经发生改变
		if oldRole != curRole {
			resp.ResponseError(c, resp.CodeUserPermissionChanged)
			c.Abort()
			return
		}
		if global.Auth.RBAC.IsGranted(oldRole, global.Auth.Permissions[p], nil) {
			c.Next()
		} else {
			log.Println(oldRole, global.Auth.Permissions)
			resp.ResponseError(c, resp.CodeUserPermissionDenied)
			c.Abort()
		}
	}
}
