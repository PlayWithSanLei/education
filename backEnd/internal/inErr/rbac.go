package inErr

import "errors"

var (
	ErrInvalidInher = errors.New("非法的继承关系")
	ErrRBACNotFound = errors.New("未找到用户自定义权限表")
)
