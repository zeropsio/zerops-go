// Generated ZEROPS sdk

package sdk

import (
	"encoding/json"
	"errors"
	"net/http"

	"context"

	"github.com/zeropsio/zerops-go/apiError"
	"github.com/zeropsio/zerops-go/dto/input/body"
	"github.com/zeropsio/zerops-go/dto/input/path"
	"github.com/zeropsio/zerops-go/dto/output"
	"github.com/zeropsio/zerops-go/sdkBase"
)

type PutAppVersionDeployResponse struct {
	success            output.Process
	err                error
	responseHeaders    http.Header
	responseStatusCode int
}

func (r PutAppVersionDeployResponse) OutputInterface() (output interface{}, err error) {
	return r.success, r.err
}

func (r PutAppVersionDeployResponse) Output() (output output.Process, err error) {
	return r.success, r.err
}

func (r PutAppVersionDeployResponse) Err() error {
	return r.err
}
func (r PutAppVersionDeployResponse) Headers() http.Header {
	return r.responseHeaders
}

func (r PutAppVersionDeployResponse) StatusCode() int {
	return r.responseStatusCode
}

func (h Handler) PutAppVersionDeploy(ctx context.Context, inputDtoPath path.AppVersionId, inputDtoBody body.PutAppVersionDeploy) (putAppVersionDeployResponse PutAppVersionDeployResponse, err error) {
	u := "/api/rest/public/app-version/" + inputDtoPath.Id.Native() + "/deploy"

	var response PutAppVersionDeployResponse
	sdkResponse := sdkBase.Put(
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
