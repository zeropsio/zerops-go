// Generated ZEROPS sdk

package sdk

import (
	"encoding/json"
	"errors"
	"net/http"

	"context"

	"github.com/zeropsio/zerops-go/apiError"
	"github.com/zeropsio/zerops-go/dto/input/body"
	"github.com/zeropsio/zerops-go/dto/output"
	"github.com/zeropsio/zerops-go/sdkBase"
)

type PutUserPasswordResponse struct {
	success            output.Success
	err                error
	responseHeaders    http.Header
	responseStatusCode int
}

func (r PutUserPasswordResponse) OutputInterface() (output interface{}, err error) {
	return r.success, r.err
}

func (r PutUserPasswordResponse) Output() (output output.Success, err error) {
	return r.success, r.err
}

func (r PutUserPasswordResponse) Err() error {
	return r.err
}
func (r PutUserPasswordResponse) Headers() http.Header {
	return r.responseHeaders
}

func (r PutUserPasswordResponse) StatusCode() int {
	return r.responseStatusCode
}

func (h Handler) PutUserPassword(ctx context.Context, inputDtoBody body.UserPutPassword) (putUserPasswordResponse PutUserPasswordResponse, err error) {
	u := "/api/rest/public/user/password"

	var response PutUserPasswordResponse
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
