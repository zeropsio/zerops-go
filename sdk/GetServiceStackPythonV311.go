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

type GetServiceStackPythonV311Response struct {
	success            output.ServiceStack
	err                error
	responseHeaders    http.Header
	responseStatusCode int
}

func (r GetServiceStackPythonV311Response) OutputInterface() (output interface{}, err error) {
	return r.success, r.err
}

func (r GetServiceStackPythonV311Response) Output() (output output.ServiceStack, err error) {
	return r.success, r.err
}

func (r GetServiceStackPythonV311Response) Err() error {
	return r.err
}
func (r GetServiceStackPythonV311Response) Headers() http.Header {
	return r.responseHeaders
}

func (r GetServiceStackPythonV311Response) StatusCode() int {
	return r.responseStatusCode
}

func (h Handler) GetServiceStackPythonV311(ctx context.Context, inputDtoPath path.ServiceStackId) (getServiceStackPythonV311Response GetServiceStackPythonV311Response, err error) {
	u := "/api/rest/public/service-stack/python_v3_11/" + inputDtoPath.Id.Native() + ""

	var response GetServiceStackPythonV311Response
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
