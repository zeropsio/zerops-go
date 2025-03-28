// Generated ZEROPS sdk

package sdk

import (
	"encoding/json"
	"errors"
	"net/http"

	"context"
	"io"

	"github.com/zeropsio/zerops-go/apiError"
	"github.com/zeropsio/zerops-go/dto/input/path"
	"github.com/zeropsio/zerops-go/dto/output"
	"github.com/zeropsio/zerops-go/sdkBase"
)

type PutAppVersionUploadResponse struct {
	success            output.Success
	err                error
	responseHeaders    http.Header
	responseStatusCode int
}

func (r PutAppVersionUploadResponse) OutputInterface() (output interface{}, err error) {
	return r.success, r.err
}

func (r PutAppVersionUploadResponse) Output() (output output.Success, err error) {
	return r.success, r.err
}

func (r PutAppVersionUploadResponse) Err() error {
	return r.err
}
func (r PutAppVersionUploadResponse) Headers() http.Header {
	return r.responseHeaders
}

func (r PutAppVersionUploadResponse) StatusCode() int {
	return r.responseStatusCode
}

func (h Handler) PutAppVersionUpload(ctx context.Context, inputDtoPath path.AppVersionId, inputBody io.Reader) (putAppVersionUploadResponse PutAppVersionUploadResponse, err error) {
	u := "/api/rest/public/app-version/" + inputDtoPath.Id.Native() + "/upload"

	var response PutAppVersionUploadResponse
	sdkResponse := sdkBase.PutRaw(
		ctx,
		h.environment,
		u,
		inputBody,
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
