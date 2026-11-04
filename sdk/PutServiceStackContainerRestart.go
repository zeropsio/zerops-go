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

type PutServiceStackContainerRestartResponse struct {
	success            output.Process
	err                error
	responseHeaders    http.Header
	responseStatusCode int
}

func (r PutServiceStackContainerRestartResponse) OutputInterface() (output interface{}, err error) {
	return r.success, r.err
}

func (r PutServiceStackContainerRestartResponse) Output() (output output.Process, err error) {
	return r.success, r.err
}

func (r PutServiceStackContainerRestartResponse) Err() error {
	return r.err
}
func (r PutServiceStackContainerRestartResponse) Headers() http.Header {
	return r.responseHeaders
}

func (r PutServiceStackContainerRestartResponse) StatusCode() int {
	return r.responseStatusCode
}

func (h Handler) PutServiceStackContainerRestart(ctx context.Context, inputDtoPath path.ServiceStackContainerId) (putServiceStackContainerRestartResponse PutServiceStackContainerRestartResponse, err error) {
	u := "/api/rest/public/service-stack/" + inputDtoPath.Id.Native() + "/container/" + inputDtoPath.ContainerId.Native() + "/restart"

	var response PutServiceStackContainerRestartResponse
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
