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

type GetAuthInfoResponse struct {
	success            output.Auth
	err                error
	responseHeaders    http.Header
	responseStatusCode int
}

func (r GetAuthInfoResponse) OutputInterface() (output interface{}, err error) {
	return r.success, r.err
}

func (r GetAuthInfoResponse) Output() (output output.Auth, err error) {
	return r.success, r.err
}

func (r GetAuthInfoResponse) Err() error {
	return r.err
}
func (r GetAuthInfoResponse) Headers() http.Header {
	return r.responseHeaders
}

func (r GetAuthInfoResponse) StatusCode() int {
	return r.responseStatusCode
}

func (h Handler) GetAuthInfo(ctx context.Context) (getAuthInfoResponse GetAuthInfoResponse, err error) {
	u := "/api/rest/public/auth/info"

	var response GetAuthInfoResponse
	sdkResponse := sdkBase.Get(
		ctx,
		h.environment,
		u,
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
