package apiError

import (
	"errors"

	"github.com/zeropsio/zerops-go/errorCode"
)

func HasHTTPCode(err error, code int) bool {
	var apiErr Error
	if !errors.As(err, &apiErr) {
		return false
	}
	return apiErr.GetHttpStatusCode() == code
}

func HasErrorCode(err error, code errorCode.ErrorCode) bool {
	var apiErr Error
	if !errors.As(err, &apiErr) {
		return false
	}
	return errorCode.ErrorCode(apiErr.GetErrorCode()) == code
}
