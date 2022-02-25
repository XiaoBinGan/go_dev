package errors

import (
	"github.com/micro/go-micro/errors"
	"go_dev/day17/demo/src/share/config"
)

const (
	//error code 200
	errorCodeUserSuccess = 200
)
var (
	//error type
	ErrorUserSuccess = errors.New(
		 config.ServiceNameUser,"操作成功",errorCodeUserSuccess,
		 )
	ErrorUserFailed = errors.New(
		config.ServiceNameUser,"操作异常",errorCodeUserSuccess,
		)
	ErrorUserAlready = errors.New(
		config.ServiceNameUser,"该邮箱已经被注册过了～",errorCodeUserSuccess,
		)
	ErrorUserLoginFailed = errors.New(
		config.ServiceNameUser,"你没有买过该电影票，无法进行评分～～",errorCodeUserSuccess,
		)
)