package router

import (
	"github.com/gin-gonic/gin"
	_ "github.com/impact-eintr/education/docs"
	"github.com/impact-eintr/education/internal/middleware"
	v1 "github.com/impact-eintr/education/internal/router/api/v1"
	"github.com/impact-eintr/education/pkg/logger"

	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func NewRouter() (r *gin.Engine, err error) {
	r = gin.New()

	r.Use(logger.GinLogger())
	r.Use(logger.GinRecovery(true))
	r.Use(middleware.Cors())

	// 注册swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// apiv1路由组
	apiv1 := r.Group("/api/v1")
	apiv1.POST("/signup", v1.SignUpHandler)
	apiv1.POST("/login", v1.LoginHandler)
	apiv1.POST("/password", v1.ForgetPasswordHandler)

	// 用户家目录路由
	homeGroup := apiv1.Group("/home")
	{
		homeGroup.Use(middleware.JWTAuthMiddleware())

		// 这里以角色代表组织
		homeGroup.POST("/org", middleware.RBACMiddleware("修改组织"), v1.UpdateRBAC)
		homeGroup.GET("/org", middleware.RBACMiddleware("获取组织"), v1.GetRBAC)
		homeGroup.GET("/org/:id", middleware.RBACMiddleware("查询组织"), v1.QueryRBAC)

		// 用户管理路由
		homeGroup.GET("/users", v1.GetUsersHandler)
		homeGroup.POST("/users", v1.UpdateUsersHandler)
		homeGroup.HEAD("/users/:id", v1.BlockUserHandler)
		homeGroup.POST("/password/:id", middleware.RBACMiddleware("修改密码"), v1.ResetPasswordHandler)
		homeGroup.DELETE("/users/:id", v1.DeleteUsersHandler)
	}

	// 栏目管理路由
	columnGroup := apiv1.Group("/column")
	{
		columnGroup.Use(middleware.JWTAuthMiddleware())
		columnGroup.POST("/article")
	}

	return r, nil

}
