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

type GetServiceStackPostgresqlV16Response struct {
	success            output.ServiceStackPostgreSql
	err                error
	responseHeaders    http.Header
	responseStatusCode int
}

func (r GetServiceStackPostgresqlV16Response) OutputInterface() (output interface{}, err error) {
	return r.success, r.err
}

func (r GetServiceStackPostgresqlV16Response) Output() (output output.ServiceStackPostgreSql, err error) {
	return r.success, r.err
}

func (r GetServiceStackPostgresqlV16Response) Err() error {
	return r.err
}
func (r GetServiceStackPostgresqlV16Response) Headers() http.Header {
	return r.responseHeaders
}

func (r GetServiceStackPostgresqlV16Response) StatusCode() int {
	return r.responseStatusCode
}

func (h Handler) GetServiceStackPostgresqlV16(ctx context.Context, inputDtoPath path.ServiceStackId) (getServiceStackPostgresqlV16Response GetServiceStackPostgresqlV16Response, err error) {
	u := "/api/rest/public/service-stack/postgresql_v16/" + inputDtoPath.Id.Native() + ""

	var response GetServiceStackPostgresqlV16Response
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
