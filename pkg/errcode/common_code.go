package errcode

import (
	"fmt"
	"net/http"
)

var (
	Success                   = NewError(0, "成功")
	ServerError               = NewError(1000, "服务内部错误")
	InvalidParams             = NewError(1001, "入参错误")
	NotFound                  = NewError(1002, "找不到")
	UnauthorizedAuthNotExist  = NewError(1003, "鉴权失败，找不到对应的 AppKey 和 AppSecret")
	UnauthorizedTokenError    = NewError(1004, "鉴权失败，Token 错误")
	UnauthorizedTokenTimeout  = NewError(1005, "鉴权失败，Token 超时")
	UnauthorizedTokenGenerate = NewError(1006, "鉴权失败，Token 生成失败")
	TooManyRequests           = NewError(1007, "请求过多")
)

var (
	ErrorGetTagListFail = NewError(2001, "获取标签列表失败")
	ErrorCreateTagFail  = NewError(2002, "创建标签失败")
	ErrorUpdateTagFail  = NewError(2003, "更新标签失败")
	ErrorDeleteTagFail  = NewError(2004, "删除标签失败")
	ErrorCountTagFail   = NewError(2005, "统计标签失败")
)

type Error struct {
	Code    int      `json:"code"`
	Msg     string   `json:"msg"`
	Details []string `json:"details"`
}

var codes = map[int]string{}

func NewError(code int, msg string) *Error {
	if _, ok := codes[code]; ok {
		panic(fmt.Sprintf("错误码 %d 已经存在，请更换一个", code))
	}
	codes[code] = msg
	return &Error{Code: code, Msg: msg}
}

func (e *Error) Error() string {
	return fmt.Sprintf("错误码：%d, 错误信息:：%s", e.Code, e.Msg)
}

func (e *Error) Msgf(args []interface{}) string {
	return fmt.Sprintf(e.Msg, args...)
}

func (e *Error) WithDetails(details ...string) *Error {
	newError := *e
	newError.Details = []string{}
	for _, d := range details {
		newError.Details = append(newError.Details, d)
	}

	return &newError
}

func (e *Error) StatusCode() int {
	switch e.Code {
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
	case UnauthorizedTokenGenerate.Code:
		fallthrough
	case UnauthorizedTokenTimeout.Code:
		return http.StatusUnauthorized
	case TooManyRequests.Code:
		return http.StatusTooManyRequests
	}

	return http.StatusInternalServerError
}
