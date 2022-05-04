// Generated ZEROPS sdk

package sdk

import (
	"encoding/json"
	"errors"
	"net/http"

	"net/url"
	"strconv"
	"strings"

	"context"

	"github.com/zeropsio/zerops-go/apiError"
	"github.com/zeropsio/zerops-go/dto/input/path"
	"github.com/zeropsio/zerops-go/dto/input/query"
	"github.com/zeropsio/zerops-go/dto/output"
	"github.com/zeropsio/zerops-go/sdkBase"
)

var _ strconv.NumError

type GetCardPaymentInvoiceResponse struct {
	success            output.Invoice
	err                error
	responseHeaders    http.Header
	responseStatusCode int
}

func (r GetCardPaymentInvoiceResponse) OutputInterface() (output interface{}, err error) {
	return r.success, r.err
}

func (r GetCardPaymentInvoiceResponse) Output() (output output.Invoice, err error) {
	return r.success, r.err
}

func (r GetCardPaymentInvoiceResponse) Err() error {
	return r.err
}
func (r GetCardPaymentInvoiceResponse) Headers() http.Header {
	return r.responseHeaders
}

func (r GetCardPaymentInvoiceResponse) StatusCode() int {
	return r.responseStatusCode
}

func (h Handler) GetCardPaymentInvoice(ctx context.Context, inputDtoPath path.InvoiceId, inputDtoQuery query.Invoice) (getCardPaymentInvoiceResponse GetCardPaymentInvoiceResponse, err error) {
	u := "/api/rest/public/card-payment/" + inputDtoPath.Id.Native() + "/invoice"

	var queryParams []string
	{
		param := inputDtoQuery.LanguageId.Native()
		queryParams = append(queryParams, "languageId="+url.QueryEscape(param))
	}

	if len(queryParams) > 0 {
		u += "?" + strings.Join(queryParams, "&")
	}

	var response GetCardPaymentInvoiceResponse
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
