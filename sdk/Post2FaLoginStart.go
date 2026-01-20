// Generated ZEROPS sdk

package sdk

import (
	"encoding/json"
	"errors"
	"net/http"

	"context"

	"github.com/zeropsio/zerops-go/apiError"
	"github.com/zeropsio/zerops-go/dto/output"
	"github.com/zeropsio/zerops-go/sdkBase"
)

type Post2FaLoginStartResponse struct {
	success            output.WebAuthnOptions
	err                error
	responseHeaders    http.Header
	responseStatusCode int
}

func (r Post2FaLoginStartResponse) OutputInterface() (output interface{}, err error) {
	return r.success, r.err
}

func (r Post2FaLoginStartResponse) Output() (output output.WebAuthnOptions, err error) {
	return r.success, r.err
}

func (r Post2FaLoginStartResponse) Err() error {
	return r.err
}
func (r Post2FaLoginStartResponse) Headers() http.Header {
	return r.responseHeaders
}

func (r Post2FaLoginStartResponse) StatusCode() int {
	return r.responseStatusCode
}

func (h Handler) Post2FaLoginStart(ctx context.Context) (post2FaLoginStartResponse Post2FaLoginStartResponse, err error) {
	u := "/api/rest/public/2fa/login/start"

	var response Post2FaLoginStartResponse
	sdkResponse := sdkBase.Post(
		ctx,
		h.environment,
		u,
		struct{}{},
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
