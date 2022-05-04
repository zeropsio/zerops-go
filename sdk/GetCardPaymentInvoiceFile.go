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

type GetCardPaymentInvoiceFileResponse struct {
	success            output.FileDownload
	err                error
	responseHeaders    http.Header
	responseStatusCode int
}

func (r GetCardPaymentInvoiceFileResponse) OutputInterface() (output interface{}, err error) {
	return r.success, r.err
}

func (r GetCardPaymentInvoiceFileResponse) Output() (output output.FileDownload, err error) {
	return r.success, r.err
}

func (r GetCardPaymentInvoiceFileResponse) Err() error {
	return r.err
}
func (r GetCardPaymentInvoiceFileResponse) Headers() http.Header {
	return r.responseHeaders
}

func (r GetCardPaymentInvoiceFileResponse) StatusCode() int {
	return r.responseStatusCode
}

func (h Handler) GetCardPaymentInvoiceFile(ctx context.Context, inputDtoPath path.DownloadToken, inputDtoQuery query.File) (getCardPaymentInvoiceFileResponse GetCardPaymentInvoiceFileResponse, err error) {
	u := "/api/rest/public/card-payment/invoice/file/" + inputDtoPath.Token.Native() + ""

	var queryParams []string
	if param, ok := inputDtoQuery.Preview.Get(); ok {
		queryParams = append(queryParams, "preview="+url.QueryEscape(strconv.FormatBool(param.Native())))
	}

	if len(queryParams) > 0 {
		u += "?" + strings.Join(queryParams, "&")
	}

	var response GetCardPaymentInvoiceFileResponse
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
