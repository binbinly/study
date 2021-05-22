package errno

import (
	"fmt"
	"net/http"
)

// Error 返回错误码和消息的结构体
type Error struct {
	code    int      `json:"code"`
	msg     string   `json:"msg"`
	details []string `json:"details"`
}

func NewError(code int, msg string) *Error {
	return &Error{code: code, msg: msg}
}

func (e *Error) Error() string {
	return fmt.Sprintf("code：%d, msg:：%s", e.Code(), e.Msg())
}

func (e *Error) Code() int {
	return e.code
}

func (e *Error) Msg() string {
	return e.msg
}

func (e *Error) Details() []string {
	return e.details
}

func (e *Error) WithDetails(details ...string) *Error {
	newError := *e
	newError.details = []string{}
	for _, d := range details {
		newError.details = append(newError.details, d)
	}

	return &newError
}

func (e *Error) StatusCode() int {
	switch e.Code() {
	case Success.Code():
		return http.StatusOK
	case InternalServerError.Code():
		return http.StatusInternalServerError
	case ErrInvalidParam.Code():
		return http.StatusBadRequest
	case ErrToken.Code():
		fallthrough
	case ErrInvalidToken.Code():
		fallthrough
	case ErrTokenTimeout.Code():
		return http.StatusUnauthorized
	case TooManyRequests.Code():
		return http.StatusTooManyRequests
	}

	return http.StatusInternalServerError
}

// Err represents an error
type Err struct {
	Code    int
	Message string
	Err     error
}

func (err *Err) Error() string {
	return fmt.Sprintf("Err - code: %d, message: %s, error: %s", err.Code, err.Message, err.Err)
}

// DecodeErr 对错误进行解码，返回错误code和错误提示
func DecodeErr(err error) (int, string) {
	if err == nil {
		return Success.code, Success.msg
	}

	switch typed := err.(type) {
	case *Err:
		return typed.Code, typed.Message
	case *Error:
		return typed.code, typed.msg
	default:
	}

	return InternalServerError.Code(), err.Error()
}
