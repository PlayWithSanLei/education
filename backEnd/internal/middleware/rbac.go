package middleware

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/impact-eintr/education/global"
	"github.com/impact-eintr/education/internal/dao/cache"
	"github.com/impact-eintr/education/pkg/resp"
)

func RBACMiddleware(p string) gin.HandlerFunc {
	return func(c *gin.Context) {
		oldRole := c.GetString("userRole")
		bRole := cache.CacheCheck(fmt.Sprintf("%d", c.GetInt64("userID")))
		if len(bRole) == 0 {
			// TODO 如果缓存中没有当前用户数据(缓存失败 缓存失效) 尝试获取新的数据

		}
		curRole := string(bRole)
		log.Println(curRole, oldRole, c.GetInt64("userID"))
		// 先检测当前用户的角色是否已经发生改变
		if oldRole != curRole {
			resp.ResponseError(c, resp.CodeUserPermissionChanged)
			c.Abort()
			return
		}
		if global.Auth.RBAC.IsGranted(oldRole, global.Auth.Permissions[p], nil) {
			c.Next()
		} else {
			resp.ResponseError(c, resp.CodeUserPermissionDenied)
			c.Abort()
		}
	}
}
