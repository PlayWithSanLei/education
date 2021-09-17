package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/impact-eintr/education/internal/service"
	"github.com/impact-eintr/education/pkg/resp"
	"go.uber.org/zap"
)

func GetUsersHandler(c *gin.Context) {
	users, err := service.GetUsers()
	if err != nil {
		zap.L().Error("未知错误", zap.Error(err))
		resp.ResponseError(c, resp.CodeServerBusy)
		return
	}
	resp.ResponseSuccess(c, users)

}
