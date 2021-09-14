package resp

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Resp struct {
	Code code        `json:"code"`
	Msg  interface{} `json:"msg"`
}

func ResponseError(c *gin.Context, code code) {
	rd := &Resp{
		Code: code,
		Msg:  code.Msg(),
	}
	c.JSON(http.StatusOK, rd)
}

func ResponseErrorWithMsg(c *gin.Context, code code, msg interface{}) {
	rd := &Resp{
		Code: code,
		Msg:  code.Msg(),
	}
	c.JSON(http.StatusOK, rd)
}

func ResponseSuccess(c *gin.Context, msg interface{}) {
	rd := &Resp{
		Code: CodeSuccess,
		Msg:  msg,
	}
	c.JSON(http.StatusOK, rd)
}
