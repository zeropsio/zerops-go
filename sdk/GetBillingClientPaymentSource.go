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

type GetBillingClientPaymentSourceResponse struct {
	success            output.PaymentSources
	err                error
	responseHeaders    http.Header
	responseStatusCode int
}

func (r GetBillingClientPaymentSourceResponse) OutputInterface() (output interface{}, err error) {
	return r.success, r.err
}

func (r GetBillingClientPaymentSourceResponse) Output() (output output.PaymentSources, err error) {
	return r.success, r.err
}

func (r GetBillingClientPaymentSourceResponse) Err() error {
	return r.err
}
func (r GetBillingClientPaymentSourceResponse) Headers() http.Header {
	return r.responseHeaders
}

func (r GetBillingClientPaymentSourceResponse) StatusCode() int {
	return r.responseStatusCode
}

func (h Handler) GetBillingClientPaymentSource(ctx context.Context, inputDtoPath path.ClientId) (getBillingClientPaymentSourceResponse GetBillingClientPaymentSourceResponse, err error) {
	u := "/api/rest/public/billing/client/" + inputDtoPath.Id.Native() + "/payment-source"

	var response GetBillingClientPaymentSourceResponse
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
