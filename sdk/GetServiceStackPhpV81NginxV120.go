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

type GetServiceStackPhpV81NginxV120Response struct {
	success            output.ServiceStackPhpNginx
	err                error
	responseHeaders    http.Header
	responseStatusCode int
}

func (r GetServiceStackPhpV81NginxV120Response) OutputInterface() (output interface{}, err error) {
	return r.success, r.err
}

func (r GetServiceStackPhpV81NginxV120Response) Output() (output output.ServiceStackPhpNginx, err error) {
	return r.success, r.err
}

func (r GetServiceStackPhpV81NginxV120Response) Err() error {
	return r.err
}
func (r GetServiceStackPhpV81NginxV120Response) Headers() http.Header {
	return r.responseHeaders
}

func (r GetServiceStackPhpV81NginxV120Response) StatusCode() int {
	return r.responseStatusCode
}

func (h Handler) GetServiceStackPhpV81NginxV120(ctx context.Context, inputDtoPath path.ServiceStackId) (getServiceStackPhpV81NginxV120Response GetServiceStackPhpV81NginxV120Response, err error) {
	u := "/api/rest/public/service-stack/php_v8_1_nginx_v1_20/" + inputDtoPath.Id.Native() + ""

	var response GetServiceStackPhpV81NginxV120Response
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
