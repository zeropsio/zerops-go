// Generated ZEROPS sdk

package sdk

import (
	"encoding/json"
	"errors"
	"net/http"

	"context"

	"github.com/zeropsio/zerops-go/apiError"
	"github.com/zeropsio/zerops-go/dto/input/path"
	"github.com/zeropsio/zerops-go/dto/output"
	"github.com/zeropsio/zerops-go/sdkBase"
)

type PutServiceStackDisableSubdomainAccessResponse struct {
	success            output.Process
	err                error
	responseHeaders    http.Header
	responseStatusCode int
}

func (r PutServiceStackDisableSubdomainAccessResponse) OutputInterface() (output interface{}, err error) {
	return r.success, r.err
}

func (r PutServiceStackDisableSubdomainAccessResponse) Output() (output output.Process, err error) {
	return r.success, r.err
}

func (r PutServiceStackDisableSubdomainAccessResponse) Err() error {
	return r.err
}
func (r PutServiceStackDisableSubdomainAccessResponse) Headers() http.Header {
	return r.responseHeaders
}

func (r PutServiceStackDisableSubdomainAccessResponse) StatusCode() int {
	return r.responseStatusCode
}

func (h Handler) PutServiceStackDisableSubdomainAccess(ctx context.Context, inputDtoPath path.ServiceStackId) (putServiceStackDisableSubdomainAccessResponse PutServiceStackDisableSubdomainAccessResponse, err error) {
	u := "/api/rest/public/service-stack/" + inputDtoPath.Id.Native() + "/disable-subdomain-access"

	var response PutServiceStackDisableSubdomainAccessResponse
	sdkResponse := sdkBase.Put(
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
