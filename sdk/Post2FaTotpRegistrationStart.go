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

type Post2FaTotpRegistrationStartResponse struct {
	success            output.TotpRegistration
	err                error
	responseHeaders    http.Header
	responseStatusCode int
}

func (r Post2FaTotpRegistrationStartResponse) OutputInterface() (output interface{}, err error) {
	return r.success, r.err
}

func (r Post2FaTotpRegistrationStartResponse) Output() (output output.TotpRegistration, err error) {
	return r.success, r.err
}

func (r Post2FaTotpRegistrationStartResponse) Err() error {
	return r.err
}
func (r Post2FaTotpRegistrationStartResponse) Headers() http.Header {
	return r.responseHeaders
}

func (r Post2FaTotpRegistrationStartResponse) StatusCode() int {
	return r.responseStatusCode
}

func (h Handler) Post2FaTotpRegistrationStart(ctx context.Context) (post2FaTotpRegistrationStartResponse Post2FaTotpRegistrationStartResponse, err error) {
	u := "/api/rest/public/2fa/totp/registration/start"

	var response Post2FaTotpRegistrationStartResponse
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
