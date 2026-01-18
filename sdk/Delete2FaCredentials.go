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

type Delete2FaCredentialsResponse struct {
	success            output.Auth
	err                error
	responseHeaders    http.Header
	responseStatusCode int
}

func (r Delete2FaCredentialsResponse) OutputInterface() (output interface{}, err error) {
	return r.success, r.err
}

func (r Delete2FaCredentialsResponse) Output() (output output.Auth, err error) {
	return r.success, r.err
}

func (r Delete2FaCredentialsResponse) Err() error {
	return r.err
}
func (r Delete2FaCredentialsResponse) Headers() http.Header {
	return r.responseHeaders
}

func (r Delete2FaCredentialsResponse) StatusCode() int {
	return r.responseStatusCode
}

func (h Handler) Delete2FaCredentials(ctx context.Context, inputDtoPath path.WebAuthnCredentialId) (delete2FaCredentialsResponse Delete2FaCredentialsResponse, err error) {
	u := "/api/rest/public/2fa/credentials/" + inputDtoPath.Id.Native() + ""

	var response Delete2FaCredentialsResponse
	sdkResponse := sdkBase.Delete(
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
