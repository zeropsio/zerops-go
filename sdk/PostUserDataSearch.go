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

type PostUserDataSearchResponse struct {
	success            output.EsUserDataResponse
	err                error
	responseHeaders    http.Header
	responseStatusCode int
}

func (r PostUserDataSearchResponse) OutputInterface() (output interface{}, err error) {
	return r.success, r.err
}

func (r PostUserDataSearchResponse) Output() (output output.EsUserDataResponse, err error) {
	return r.success, r.err
}

func (r PostUserDataSearchResponse) Err() error {
	return r.err
}
func (r PostUserDataSearchResponse) Headers() http.Header {
	return r.responseHeaders
}

func (r PostUserDataSearchResponse) StatusCode() int {
	return r.responseStatusCode
}

func (h Handler) PostUserDataSearch(ctx context.Context, inputDtoBody body.EsFilter) (postUserDataSearchResponse PostUserDataSearchResponse, err error) {
	u := "/api/rest/public/user-data/search"

	var response PostUserDataSearchResponse
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
