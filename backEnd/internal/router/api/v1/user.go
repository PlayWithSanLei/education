package v1

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/impact-eintr/education/global"
	"github.com/impact-eintr/education/internal/inErr"
	"github.com/impact-eintr/education/internal/model"
	"github.com/impact-eintr/education/internal/service"
	"github.com/impact-eintr/education/pkg/resp"
	"go.uber.org/zap"
)

// GetUesrsHandler 获取用户数据接口
// @Summary 获取用户数据接口
// @Description 请求后可以获取用户数据
// @Tags 用户相关api
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Security ApiKeyAuth
// @Success 200 {object} resp.Resp{code=int,msg=string}
// @Router /api/v1/home/users [get]
func GetUsersHandler(c *gin.Context) {
	users, err := service.GetUsers()
	if err != nil {
		zap.L().Error("未知错误", zap.Error(err))
		resp.ResponseError(c, resp.CodeServerBusy)
		return
	}
	resp.ResponseSuccess(c, users)

}

// UpdateUesrsHandler 更新用户数据接口
// @Summary 更新用户数据接口
// @Description 请求后可以将post请求body中提供的用户新数据替换原来的用户数据
// @Tags 用户相关api
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param 用户信息 body model.UserReq true "用户信息"
// @Security ApiKeyAuth
// @Success 200 {object} resp.Resp{code=int,msg=string}
// @Router /api/v1/home/users [post]
func UpdateUsersHandler(c *gin.Context) {
	var err error
	reqUser := new(model.UserReq)
	if err := c.ShouldBindJSON(reqUser); err != nil {
		zap.L().Error("update user with invalid param", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			resp.ResponseError(c, resp.CodeInvalidParam)
			return
		}
		resp.ResponseErrorWithMsg(c, resp.CodeInvalidParam, errs.Translate(global.Trans))
		return
	}

	id, err := strconv.ParseInt(reqUser.UserID, 10, 64)
	if err != nil {
		zap.L().Error("解析错误", zap.Error(err))
		resp.ResponseError(c, resp.CodeInvalidPath)
		return
	}

	err = service.UpdateUser(reqUser, id)
	if err != nil {
		if err == inErr.ErrInvalidUnit {
			zap.L().Error("user 想要获取交科权限", zap.Error(err))
			resp.ResponseError(c, resp.CodeInvalidUnit)
			return
		}
		zap.L().Error("未知错误", zap.Error(err))
		resp.ResponseError(c, resp.CodeServerBusy)
		return
	}
	resp.ResponseSuccess(c, nil)

}

// DeleteUesrsHandler 删除用户数据接口
// @Summary 删除用户数据接口
// @Description 请求后可以删除用户数据
// @Tags 用户相关api
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Security ApiKeyAuth
// @Success 200 {object} resp.Resp{code=int,msg=string}
// @Router /api/v1/home/users/{id} [delete]
func DeleteUsersHandler(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		zap.L().Error("解析错误", zap.Error(err))
		resp.ResponseError(c, resp.CodeInvalidPath)
		return
	}

	err = service.DeleteUser(id)
	if err != nil {
		zap.L().Error("未知错误", zap.Error(err))
		resp.ResponseError(c, resp.CodeServerBusy)
		return
	}
	resp.ResponseSuccess(c, nil)

}

// BlockUesrsHandler 封禁用户接口
// @Summary 封禁用户接口
// @Description 请求后可以封禁用户
// @Tags 用户相关api
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Security ApiKeyAuth
// @Success 200 {object} resp.Resp{code=int,msg=string}
// @Router /api/v1/home/users/{id} [head]
func BlockUserHandler(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		zap.L().Error("解析错误", zap.Error(err))
		resp.ResponseError(c, resp.CodeInvalidPath)
		return
	}

	err = service.BlockUser(id)
	if err != nil {
		zap.L().Error("未知错误", zap.Error(err))
		resp.ResponseError(c, resp.CodeServerBusy)
		return
	}
	resp.ResponseSuccess(c, nil)

}
