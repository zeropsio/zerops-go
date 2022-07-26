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

type PostS3BucketResponse struct {
	success            output.Success
	err                error
	responseHeaders    http.Header
	responseStatusCode int
}

func (r PostS3BucketResponse) OutputInterface() (output interface{}, err error) {
	return r.success, r.err
}

func (r PostS3BucketResponse) Output() (output output.Success, err error) {
	return r.success, r.err
}

func (r PostS3BucketResponse) Err() error {
	return r.err
}
func (r PostS3BucketResponse) Headers() http.Header {
	return r.responseHeaders
}

func (r PostS3BucketResponse) StatusCode() int {
	return r.responseStatusCode
}

func (h Handler) PostS3Bucket(ctx context.Context, inputDtoPath path.ServiceStackIdNamed, inputDtoBody body.PostS3Bucket) (postS3BucketResponse PostS3BucketResponse, err error) {
	u := "/api/rest/public/s3/" + inputDtoPath.ServiceStackId.Native() + "/bucket"

	var response PostS3BucketResponse
	sdkResponse := sdkBase.Post(
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
