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

type GetServiceStackPhpV81ApacheV24Response struct {
	success            output.ServiceStack
	err                error
	responseHeaders    http.Header
	responseStatusCode int
}

func (r GetServiceStackPhpV81ApacheV24Response) OutputInterface() (output interface{}, err error) {
	return r.success, r.err
}

func (r GetServiceStackPhpV81ApacheV24Response) Output() (output output.ServiceStack, err error) {
	return r.success, r.err
}

func (r GetServiceStackPhpV81ApacheV24Response) Err() error {
	return r.err
}
func (r GetServiceStackPhpV81ApacheV24Response) Headers() http.Header {
	return r.responseHeaders
}

func (r GetServiceStackPhpV81ApacheV24Response) StatusCode() int {
	return r.responseStatusCode
}

func (h Handler) GetServiceStackPhpV81ApacheV24(ctx context.Context, inputDtoPath path.ServiceStackId) (getServiceStackPhpV81ApacheV24Response GetServiceStackPhpV81ApacheV24Response, err error) {
	u := "/api/rest/public/service-stack/php_v8_1_apache_v2_4/" + inputDtoPath.Id.Native() + ""

	var response GetServiceStackPhpV81ApacheV24Response
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
