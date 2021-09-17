package v1

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/impact-eintr/education/global"
	"github.com/impact-eintr/education/internal/inErr"
	"github.com/impact-eintr/education/internal/service"
	"github.com/impact-eintr/education/pkg/resp"
	"go.uber.org/zap"
)

// UpdateRBAC 更新RBAC模型接口
// @Summary 更新RBAC模型接口
// @Description 请求后可以将post请求body中提供的RBAC新模型替换原来的RBAC模型
// @Tags 用户相关api
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param 用户信息 body model.UserReq true "用户信息"
// @Security ApiKeyAuth
// @Success 200 {object} resp.Resp{code=int,msg=string}
// @Router /api/v1/home/users [post]
func UpdateRBAC(c *gin.Context) {
	jsons := make([]map[string][]string, 2, 2) //注意该结构接受的内容
	if err := c.BindJSON(&jsons); err != nil {
		zap.L().Error("序列化失败", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			resp.ResponseError(c, resp.CodeInvalidParam)
		} else {
			resp.ResponseErrorWithMsg(c, resp.CodeInvalidParam, errs.Translate(global.Trans))
		}
		c.Abort()
		return
	} else if len(jsons) != 2 {
		resp.ResponseError(c, resp.CodeInvalidParam)
		c.Abort()
		return
	}

	if err := service.UpdateRBAC(jsons[0], jsons[1]); err != nil {
		if err == inErr.ErrInvalidInher {
			zap.L().Error("出现环继承", zap.Error(err))
			resp.ResponseError(c, resp.CodeInvalidInher)
			c.Abort()
			return
		} else {
			zap.L().Error("未知错误", zap.Error(err))
			errs, ok := err.(validator.ValidationErrors)
			if !ok {
				resp.ResponseError(c, resp.CodeServerBusy)
			} else {
				resp.ResponseErrorWithMsg(c, resp.CodeServerBusy, errs.Translate(global.Trans))
			}
			c.Abort()
			return
		}
	}
	resp.ResponseSuccess(c, nil)

}

// GetRBAC 获取RBAC数据接口
// @Summary 获取RBAC数据接口
// @Description 请求后可以获取RBAC数据
// @Tags 用户相关api
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Security ApiKeyAuth
// @Success 200 {object} resp.Resp{code=int,msg=string}
// @Router /api/v1/home/users [get]
func GetRBAC(c *gin.Context) {
	r, err := service.GetRBAC()
	if err != nil {
		if err == inErr.ErrRBACNotFound {
			zap.L().Error("未找到用户自定义权限文件", zap.Error(err))
			resp.ResponseError(c, resp.CodeServerBusy)
			return

		}
		zap.L().Error("未知错误", zap.Error(err))
		resp.ResponseError(c, resp.CodeServerBusy)
		return
	}
	resp.ResponseSuccess(c, r)
}

func QueryRBAC(c *gin.Context) {
	roles, err := service.QueryRBAC(c.Param("id"))
	if err != nil {
		log.Println(err)
		return
	}
	resp.ResponseSuccess(c, roles)
}
