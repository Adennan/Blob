package error

import (
	"fmt"
	"log"
	"net/http"
)

var (
	codes                     = make(map[int]string)
	Success                   = NewError(0, "成功")
	ServerError               = NewError(10000000, "服务内部错误")
	InvalidParams             = NewError(10000001, "参数错误")
	NotFound                  = NewError(10000002, "数据不存在")
	UnauthorizedAuthNotExist  = NewError(10000003, "鉴权失败：找不到对应的 Appkey 和 AppSerect")
	UnauthorizedTokenError    = NewError(10000004, "鉴权失败：Token 错误")
	UnauthorizedTokenTimeout  = NewError(10000005, "鉴权失败：Token 请求超时")
	UnauthorizedTokenGenerate = NewError(10000006, "鉴权失败：Token 生成失败")
)

// Error define new error type
type Error struct {
	Code    int      `json:"code"`
	Msg     string   `json:"msg"`
	Details []string `json:"details"`
}

func NewError(code int, msg string) *Error {
	if _, ok := codes[code]; ok {
		log.Fatalf("This code %d already exist, please change it", code)
		return nil
	}
	codes[code] = msg
	return &Error{
		Code: code,
		Msg:  msg,
	}
}

func (err *Error) Error() string {
	return fmt.Sprintf("Error code is %d, Error message is %s", err.Code, err.Msg)
}

func (err *Error) WithDetails(details ...string) *Error {
	err.Details = append(err.Details, details...)
	return err
}

// Status return status of error code
func (err *Error) Status() int {
	switch err.Code {
	case Success.Code:
		return http.StatusOK
	case ServerError.Code:
		return http.StatusInternalServerError
	case InvalidParams.Code:
		return http.StatusBadRequest
	case NotFound.Code:
		return http.StatusNotFound
	case UnauthorizedAuthNotExist.Code:
		fallthrough
	case UnauthorizedTokenError.Code:
		fallthrough
	case UnauthorizedTokenTimeout.Code:
		fallthrough
	case UnauthorizedTokenGenerate.Code:
		return http.StatusUnauthorized
	}
	return http.StatusInternalServerError
}
