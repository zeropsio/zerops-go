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

type GetServiceStackMariadbV103Response struct {
	success            output.ServiceStackMariaDb
	err                error
	responseHeaders    http.Header
	responseStatusCode int
}

func (r GetServiceStackMariadbV103Response) OutputInterface() (output interface{}, err error) {
	return r.success, r.err
}

func (r GetServiceStackMariadbV103Response) Output() (output output.ServiceStackMariaDb, err error) {
	return r.success, r.err
}

func (r GetServiceStackMariadbV103Response) Err() error {
	return r.err
}
func (r GetServiceStackMariadbV103Response) Headers() http.Header {
	return r.responseHeaders
}

func (r GetServiceStackMariadbV103Response) StatusCode() int {
	return r.responseStatusCode
}

func (h Handler) GetServiceStackMariadbV103(ctx context.Context, inputDtoPath path.ServiceStackId) (getServiceStackMariadbV103Response GetServiceStackMariadbV103Response, err error) {
	u := "/api/rest/public/service-stack/mariadb_v10_3/" + inputDtoPath.Id.Native() + ""

	var response GetServiceStackMariadbV103Response
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
