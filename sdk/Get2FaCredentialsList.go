// Generated ZEROPS sdk

package sdk

import (
	"encoding/json"
	"errors"
	"net/http"

	"context"

	"github.com/zeropsio/zerops-go/apiError"
	"github.com/zeropsio/zerops-go/dto/output"
	"github.com/zeropsio/zerops-go/sdkBase"
)

type Get2FaCredentialsListResponse struct {
	success            output.WebAuthnCredentialList
	err                error
	responseHeaders    http.Header
	responseStatusCode int
}

func (r Get2FaCredentialsListResponse) OutputInterface() (output interface{}, err error) {
	return r.success, r.err
}

func (r Get2FaCredentialsListResponse) Output() (output output.WebAuthnCredentialList, err error) {
	return r.success, r.err
}

func (r Get2FaCredentialsListResponse) Err() error {
	return r.err
}
func (r Get2FaCredentialsListResponse) Headers() http.Header {
	return r.responseHeaders
}

func (r Get2FaCredentialsListResponse) StatusCode() int {
	return r.responseStatusCode
}

func (h Handler) Get2FaCredentialsList(ctx context.Context) (get2FaCredentialsListResponse Get2FaCredentialsListResponse, err error) {
	u := "/api/rest/public/2fa/credentials/list"

	var response Get2FaCredentialsListResponse
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
