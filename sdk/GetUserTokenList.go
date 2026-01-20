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

type GetUserTokenListResponse struct {
	success            output.UserTokenList
	err                error
	responseHeaders    http.Header
	responseStatusCode int
}

func (r GetUserTokenListResponse) OutputInterface() (output interface{}, err error) {
	return r.success, r.err
}

func (r GetUserTokenListResponse) Output() (output output.UserTokenList, err error) {
	return r.success, r.err
}

func (r GetUserTokenListResponse) Err() error {
	return r.err
}
func (r GetUserTokenListResponse) Headers() http.Header {
	return r.responseHeaders
}

func (r GetUserTokenListResponse) StatusCode() int {
	return r.responseStatusCode
}

func (h Handler) GetUserTokenList(ctx context.Context) (getUserTokenListResponse GetUserTokenListResponse, err error) {
	u := "/api/rest/public/user-token/list"

	var response GetUserTokenListResponse
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
