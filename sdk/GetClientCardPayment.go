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

type GetClientCardPaymentResponse struct {
	success            output.CardPaymentList
	err                error
	responseHeaders    http.Header
	responseStatusCode int
}

func (r GetClientCardPaymentResponse) OutputInterface() (output interface{}, err error) {
	return r.success, r.err
}

func (r GetClientCardPaymentResponse) Output() (output output.CardPaymentList, err error) {
	return r.success, r.err
}

func (r GetClientCardPaymentResponse) Err() error {
	return r.err
}
func (r GetClientCardPaymentResponse) Headers() http.Header {
	return r.responseHeaders
}

func (r GetClientCardPaymentResponse) StatusCode() int {
	return r.responseStatusCode
}

func (h Handler) GetClientCardPayment(ctx context.Context, inputDtoPath path.ClientId, inputDtoQuery query.ListClientCardPayments) (getClientCardPaymentResponse GetClientCardPaymentResponse, err error) {
	u := "/api/rest/public/client/" + inputDtoPath.Id.Native() + "/card-payment"

	var queryParams []string
	if param, ok := inputDtoQuery.Limit.Get(); ok {
		queryParams = append(queryParams, "limit="+url.QueryEscape(strconv.Itoa(param.Native())))
	}
	if param, ok := inputDtoQuery.Offset.Get(); ok {
		queryParams = append(queryParams, "offset="+url.QueryEscape(strconv.Itoa(param.Native())))
	}
	if param, ok := inputDtoQuery.Statuses.Get(); ok {
		queryParams = append(queryParams, "statuses="+url.QueryEscape(strings.Join(param.Native(), ",")))
	}

	if len(queryParams) > 0 {
		u += "?" + strings.Join(queryParams, "&")
	}

	var response GetClientCardPaymentResponse
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
