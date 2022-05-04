package apiError

import (
	"errors"
	"fmt"
)

type Error struct {
	HttpStatusCode int         `json:"-"`
	ErrorCode      string      `json:"code"`
	Message        string      `json:"message"`
	Meta           interface{} `json:"meta,omitempty"`
}

func (e Error) Is(target error) bool {
	var apiError *Error
	if errors.As(target, &apiError) {
		return e.GetErrorCode() == apiError.GetErrorCode()
	}
	return false
}

func (e Error) Error() string {
	return fmt.Sprintf("[%d][%s] %s", e.HttpStatusCode, e.ErrorCode, e.Message)
}

func (e Error) GetHttpStatusCode() int {
	return e.HttpStatusCode
}

func (e Error) GetErrorCode() string {
	return e.ErrorCode
}

func (e Error) GetMessage() string {
	return e.Message
}

func (e Error) GetMeta() interface{} {
	return e.Meta
}
