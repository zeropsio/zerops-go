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

type PostServiceStackUiAccessResponse struct {
	success            output.ServiceStackUi
	err                error
	responseHeaders    http.Header
	responseStatusCode int
}

func (r PostServiceStackUiAccessResponse) OutputInterface() (output interface{}, err error) {
	return r.success, r.err
}

func (r PostServiceStackUiAccessResponse) Output() (output output.ServiceStackUi, err error) {
	return r.success, r.err
}

func (r PostServiceStackUiAccessResponse) Err() error {
	return r.err
}
func (r PostServiceStackUiAccessResponse) Headers() http.Header {
	return r.responseHeaders
}

func (r PostServiceStackUiAccessResponse) StatusCode() int {
	return r.responseStatusCode
}

func (h Handler) PostServiceStackUiAccess(ctx context.Context, inputDtoPath path.ServiceStackId) (postServiceStackUiAccessResponse PostServiceStackUiAccessResponse, err error) {
	u := "/api/rest/public/service-stack/" + inputDtoPath.Id.Native() + "/ui-access"

	var response PostServiceStackUiAccessResponse
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
