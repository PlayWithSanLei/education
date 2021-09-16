package v1

import (
	"errors"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/impact-eintr/education/global"
	"github.com/impact-eintr/education/internal/inErr"
	"github.com/impact-eintr/education/internal/model"
	"github.com/impact-eintr/education/internal/service"
	"github.com/impact-eintr/education/pkg/resp"
	"go.uber.org/zap"
)

func SignUpHandler(c *gin.Context) {
	var err error
	// 获取参数 参数校验
	p := new(model.ParamSignUp)

	if err = c.ShouldBindJSON(&p); err != nil {
		zap.L().Error("SignUp With invalid param", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			resp.ResponseError(c, resp.CodeInvalidParam)
		}
		log.Println(errs)
		resp.ResponseErrorWithMsg(c, resp.CodeInvalidParam, errs.Translate(global.Trans))
		return
	}
	// 业务处理
	if err := service.SignUp(p); err != nil {
		zap.L().Error("登陆参数非法", zap.Error(err))
		if errors.Is(err, inErr.ErrUserExist) {
			resp.ResponseError(c, resp.CodeUserExist)
		} else {
			resp.ResponseError(c, resp.CodeServerBusy)
		}
		return
	}
	// 返回响应
	resp.ResponseSuccess(c, nil)

}

func LoginHandler(c *gin.Context) {
	// 获得请求参数以及参数校验
	p := new(model.ParamLogin)

	if err := c.ShouldBindJSON(&p); err != nil {
		zap.L().Error("Login with invalid param", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			resp.ResponseError(c, resp.CodeInvalidParam)
			return
		}
		resp.ResponseErrorWithMsg(c, resp.CodeInvalidParam, errs.Translate(global.Trans))
		return
	}

	// 2. 业务处理
	userID, userName, userRole, userUnit, aToken, err := service.Login(p)
	if err != nil {
		zap.L().Error("登录失败", zap.Error(err))
		if errors.Is(err, inErr.ErrUserNotExist) {
			resp.ResponseError(c, resp.CodeUserNotExist)
		} else if errors.Is(err, inErr.ErrInvalidPassword) {
			resp.ResponseError(c, resp.CodeInvalidPassword)
		} else {
			resp.ResponseError(c, resp.CodeServerBusy)
		}
		return
	}

	data := map[string]string{
		"token": aToken,
		"role":  userRole,
		"unit":  userUnit,
		"id":    userID,
		"name":  userName,
	}

	// 3. 返回响应
	resp.ResponseSuccess(c, data)

}
