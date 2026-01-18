// Generated ZEROPS sdk

package sdk

import (
	"encoding/json"
	"errors"
	"net/http"

	"context"

	"github.com/zeropsio/zerops-go/apiError"
	"github.com/zeropsio/zerops-go/dto/input/body"
	"github.com/zeropsio/zerops-go/dto/output"
	"github.com/zeropsio/zerops-go/sdkBase"
)

type Post2FaTotpLoginResponse struct {
	success            output.TwoFAAuthLogin
	err                error
	responseHeaders    http.Header
	responseStatusCode int
}

func (r Post2FaTotpLoginResponse) OutputInterface() (output interface{}, err error) {
	return r.success, r.err
}

func (r Post2FaTotpLoginResponse) Output() (output output.TwoFAAuthLogin, err error) {
	return r.success, r.err
}

func (r Post2FaTotpLoginResponse) Err() error {
	return r.err
}
func (r Post2FaTotpLoginResponse) Headers() http.Header {
	return r.responseHeaders
}

func (r Post2FaTotpLoginResponse) StatusCode() int {
	return r.responseStatusCode
}

func (h Handler) Post2FaTotpLogin(ctx context.Context, inputDtoBody body.AuthTOTP) (post2FaTotpLoginResponse Post2FaTotpLoginResponse, err error) {
	u := "/api/rest/public/2fa/totp/login"

	var response Post2FaTotpLoginResponse
	sdkResponse := sdkBase.Post(
		ctx,
		h.environment,
		u,
		inputDtoBody,
	)
	if sdkResponse.Err != nil {
		return response, sdkResponse.Err
	}
	response.responseHeaders = sdkResponse.HttpResponse.Header
	response.responseStatusCode = sdkResponse.HttpResponse.StatusCode

	decoder := json.NewDecoder(sdkResponse.ResponseData)
	if sdkResponse.HttpResponse.StatusCode < http.StatusMultipleChoices {
		if err := decoder.Decode(&response.success); err != nil {
			return response, err
		}
	} else {
		responseString := sdkResponse.ResponseData.String()
		apiErrorResponse := struct {
			Error apiError.Error `json:"error"`
		}{}
		err := decoder.Decode(&apiErrorResponse)
		if err != nil {
			return response, errors.New(sdkResponse.HttpResponse.Status + ": " + responseString)
		}
		apiErrorResponse.Error.HttpStatusCode = sdkResponse.HttpResponse.StatusCode
		response.err = apiErrorResponse.Error
	}

	return response, nil
}
