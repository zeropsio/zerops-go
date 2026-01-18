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

type Get2FaRecoveryCodesResponse struct {
	success            output.TwoFARecoveryCodes
	err                error
	responseHeaders    http.Header
	responseStatusCode int
}

func (r Get2FaRecoveryCodesResponse) OutputInterface() (output interface{}, err error) {
	return r.success, r.err
}

func (r Get2FaRecoveryCodesResponse) Output() (output output.TwoFARecoveryCodes, err error) {
	return r.success, r.err
}

func (r Get2FaRecoveryCodesResponse) Err() error {
	return r.err
}
func (r Get2FaRecoveryCodesResponse) Headers() http.Header {
	return r.responseHeaders
}

func (r Get2FaRecoveryCodesResponse) StatusCode() int {
	return r.responseStatusCode
}

func (h Handler) Get2FaRecoveryCodes(ctx context.Context) (get2FaRecoveryCodesResponse Get2FaRecoveryCodesResponse, err error) {
	u := "/api/rest/public/2fa/recovery-codes"

	var response Get2FaRecoveryCodesResponse
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
