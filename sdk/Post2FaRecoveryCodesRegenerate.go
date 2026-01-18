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

type Post2FaRecoveryCodesRegenerateResponse struct {
	success            output.TwoFARecoveryCodes
	err                error
	responseHeaders    http.Header
	responseStatusCode int
}

func (r Post2FaRecoveryCodesRegenerateResponse) OutputInterface() (output interface{}, err error) {
	return r.success, r.err
}

func (r Post2FaRecoveryCodesRegenerateResponse) Output() (output output.TwoFARecoveryCodes, err error) {
	return r.success, r.err
}

func (r Post2FaRecoveryCodesRegenerateResponse) Err() error {
	return r.err
}
func (r Post2FaRecoveryCodesRegenerateResponse) Headers() http.Header {
	return r.responseHeaders
}

func (r Post2FaRecoveryCodesRegenerateResponse) StatusCode() int {
	return r.responseStatusCode
}

func (h Handler) Post2FaRecoveryCodesRegenerate(ctx context.Context) (post2FaRecoveryCodesRegenerateResponse Post2FaRecoveryCodesRegenerateResponse, err error) {
	u := "/api/rest/public/2fa/recovery-codes/regenerate"

	var response Post2FaRecoveryCodesRegenerateResponse
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
