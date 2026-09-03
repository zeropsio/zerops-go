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

type PostProjectBackupDownloadUrlResponse struct {
	success            output.Url
	err                error
	responseHeaders    http.Header
	responseStatusCode int
}

func (r PostProjectBackupDownloadUrlResponse) OutputInterface() (output interface{}, err error) {
	return r.success, r.err
}

func (r PostProjectBackupDownloadUrlResponse) Output() (output output.Url, err error) {
	return r.success, r.err
}

func (r PostProjectBackupDownloadUrlResponse) Err() error {
	return r.err
}
func (r PostProjectBackupDownloadUrlResponse) Headers() http.Header {
	return r.responseHeaders
}

func (r PostProjectBackupDownloadUrlResponse) StatusCode() int {
	return r.responseStatusCode
}

func (h Handler) PostProjectBackupDownloadUrl(ctx context.Context, inputDtoPath path.ProjectBackup) (postProjectBackupDownloadUrlResponse PostProjectBackupDownloadUrlResponse, err error) {
	u := "/api/rest/public/project/" + inputDtoPath.Id.Native() + "/backup/download-url/" + inputDtoPath.ServiceStackId.Native() + "/" + inputDtoPath.Date.Native() + ""

	var response PostProjectBackupDownloadUrlResponse
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
