package router

import (
	"github.com/gin-gonic/gin"
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

}
