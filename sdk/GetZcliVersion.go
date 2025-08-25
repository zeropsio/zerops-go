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

type GetZcliVersionResponse struct {
	success            output.ZcliVersion
	err                error
	responseHeaders    http.Header
	responseStatusCode int
}

func (r GetZcliVersionResponse) OutputInterface() (output interface{}, err error) {
	return r.success, r.err
}

func (r GetZcliVersionResponse) Output() (output output.ZcliVersion, err error) {
	return r.success, r.err
}

func (r GetZcliVersionResponse) Err() error {
	return r.err
}
func (r GetZcliVersionResponse) Headers() http.Header {
	return r.responseHeaders
}

func (r GetZcliVersionResponse) StatusCode() int {
	return r.responseStatusCode
}

func (h Handler) GetZcliVersion(ctx context.Context) (getZcliVersionResponse GetZcliVersionResponse, err error) {
	u := "/api/rest/public/zcli/version"

	var response GetZcliVersionResponse
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
