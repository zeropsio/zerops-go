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

type PutProjectRevertChangesPublicPortRoutingResponse struct {
	success            output.Success
	err                error
	responseHeaders    http.Header
	responseStatusCode int
}

func (r PutProjectRevertChangesPublicPortRoutingResponse) OutputInterface() (output interface{}, err error) {
	return r.success, r.err
}

func (r PutProjectRevertChangesPublicPortRoutingResponse) Output() (output output.Success, err error) {
	return r.success, r.err
}

func (r PutProjectRevertChangesPublicPortRoutingResponse) Err() error {
	return r.err
}
func (r PutProjectRevertChangesPublicPortRoutingResponse) Headers() http.Header {
	return r.responseHeaders
}

func (r PutProjectRevertChangesPublicPortRoutingResponse) StatusCode() int {
	return r.responseStatusCode
}

func (h Handler) PutProjectRevertChangesPublicPortRouting(ctx context.Context, inputDtoPath path.ProjectId) (putProjectRevertChangesPublicPortRoutingResponse PutProjectRevertChangesPublicPortRoutingResponse, err error) {
	u := "/api/rest/public/project/" + inputDtoPath.Id.Native() + "/revert-changes-public-port-routing"

	var response PutProjectRevertChangesPublicPortRoutingResponse
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
