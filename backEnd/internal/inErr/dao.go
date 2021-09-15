package inErr

import "errors"

// 错误码
var (
	ErrUserExist       = errors.New("用户已经存在")
	ErrUserNotExist    = errors.New("用户不存在")
	ErrInvalidPassword = errors.New("用户名或密码错误")
	ErrNoUser          = errors.New("数据库无用户")
	ErrVerify          = errors.New("验证失败")
	ErrInvalidUnit     = errors.New("非法单位")
	ErrNotFound        = errors.New("没有相关数据")
)
