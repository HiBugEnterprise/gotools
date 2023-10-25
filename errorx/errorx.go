package errorx

import (
	"fmt"
	"net/http"
)

type Error struct {
	Type     string `json:"type"` // 业务类型
	Code     int    `json:"code"`
	Msg      string `json:"msg"` // 友好信息
	Detail   string
	Metadata Metadata // 附加信息
	Err      error
}

type Metadata map[string]any

func New(t string, code int, message string) *Error {
	return &Error{
		Type: t,
		Code: code,
		Msg:  message,
	}
}

func WithCode(t string, code ResCode) *Error {
	return &Error{
		Type: t,
		Code: int(code),
		Msg:  code.Msg(),
	}
}

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
	if inErr, ok := err.(*Error); ok {
		return inErr
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
