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

type PostAppVersionResponse struct {
	success            output.PostAppVersion
	err                error
	responseHeaders    http.Header
	responseStatusCode int
}

func (r PostAppVersionResponse) OutputInterface() (output interface{}, err error) {
	return r.success, r.err
}

func (r PostAppVersionResponse) Output() (output output.PostAppVersion, err error) {
	return r.success, r.err
}

func (r PostAppVersionResponse) Err() error {
	return r.err
}
func (r PostAppVersionResponse) Headers() http.Header {
	return r.responseHeaders
}

func (r PostAppVersionResponse) StatusCode() int {
	return r.responseStatusCode
}

func (h Handler) PostAppVersion(ctx context.Context, inputDtoBody body.PostAppVersion) (postAppVersionResponse PostAppVersionResponse, err error) {
	u := "/api/rest/public/app-version"

	var response PostAppVersionResponse
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
