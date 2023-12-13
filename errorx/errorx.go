package errorx

import (
	"errors"
	"fmt"
	"net/http"
)

type Error struct {
	BizType  string `json:"biz_type"`
	Code     int    `json:"code"`
	Msg      string `json:"msg"`
	Metadata Metadata
	Err      error
}

type Metadata map[string]any

// New 创建自定义错误, 通常用于业务并没有err返回, 但是属于业务逻辑错误. 此时需要给前端一个友好提示
// eg: New("user-service-Login", 404001, "用户不存在")
func New(bizType string, code int, message string) *Error {
	return &Error{
		BizType: bizType,
		Code:    code,
		Msg:     message,
	}
}

// Internal 创建内部错误, 通常用于服务端内部错误, 如数据库, 缓存等
func Internal(err error, format string, args ...any) *Error {
	message := fmt.Sprintf(format, args...)
	return New(http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError, message).WithError(err)
}

func NotFound(format string, args ...any) *Error {
	message := fmt.Sprintf(format, args...)
	return New(http.StatusText(http.StatusNotFound), http.StatusNotFound, message)
}

func Unauthorized(format string, args ...any) *Error {
	message := fmt.Sprintf(format, args...)
	return New(http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized, message)
}

func BadRequest(format string, args ...any) *Error {
	message := fmt.Sprintf(format, args...)
	return New(http.StatusText(http.StatusBadRequest), http.StatusBadRequest, message)
}

func Exist(format string, args ...any) *Error {
	message := fmt.Sprintf(format, args...)
	return New(http.StatusText(http.StatusConflict), http.StatusConflict, message)
}

func From(err error) *Error {
	if err == nil {
		return nil
	}

	var customErr *Error
	if errors.As(err, &customErr) {
		return customErr
	}

	return Internal(err, CodeInternalErr.Msg())
}

func (e *Error) Error() string {
	if e.Err != nil {
		return e.Msg + ": " + e.Err.Error()
	}

	return e.Msg
}

func (e *Error) Unwrap() error {
	return e.Err
}

func (e *Error) WithMessage(format string, args ...any) *Error {
	e.Msg = fmt.Sprintf(format, args...)
	return e
}

func (e *Error) WithMetadata(metadata Metadata) *Error {
	e.Metadata = metadata
	return e
}

func (e *Error) WithError(err error) *Error {
	e.Err = err
	return e
}

func IsNotFound(err error) bool {
	return From(err).Code == http.StatusNotFound
}
