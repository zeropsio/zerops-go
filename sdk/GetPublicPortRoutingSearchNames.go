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

type GetPublicPortRoutingSearchNamesResponse struct {
	success            output.EsNames
	err                error
	responseHeaders    http.Header
	responseStatusCode int
}

func (r GetPublicPortRoutingSearchNamesResponse) OutputInterface() (output interface{}, err error) {
	return r.success, r.err
}

func (r GetPublicPortRoutingSearchNamesResponse) Output() (output output.EsNames, err error) {
	return r.success, r.err
}

func (r GetPublicPortRoutingSearchNamesResponse) Err() error {
	return r.err
}
func (r GetPublicPortRoutingSearchNamesResponse) Headers() http.Header {
	return r.responseHeaders
}

func (r GetPublicPortRoutingSearchNamesResponse) StatusCode() int {
	return r.responseStatusCode
}

func (h Handler) GetPublicPortRoutingSearchNames(ctx context.Context) (getPublicPortRoutingSearchNamesResponse GetPublicPortRoutingSearchNamesResponse, err error) {
	u := "/api/rest/public/public-port-routing/search/names"

	var response GetPublicPortRoutingSearchNamesResponse
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
