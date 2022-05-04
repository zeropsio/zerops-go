// Generated ZEROPS sdk

package sdk

import (
	"encoding/json"
	"errors"
	"net/http"

	"context"

	"github.com/zeropsio/zerops-go/apiError"
	"github.com/zeropsio/zerops-go/dto/input/body"
	"github.com/zeropsio/zerops-go/dto/output"
	"github.com/zeropsio/zerops-go/sdkBase"
)

type PostServiceStackPhpV81NginxV120Response struct {
	success            output.ServiceStackProcessNginx
	err                error
	responseHeaders    http.Header
	responseStatusCode int
}

func (r PostServiceStackPhpV81NginxV120Response) OutputInterface() (output interface{}, err error) {
	return r.success, r.err
}

func (r PostServiceStackPhpV81NginxV120Response) Output() (output output.ServiceStackProcessNginx, err error) {
	return r.success, r.err
}

func (r PostServiceStackPhpV81NginxV120Response) Err() error {
	return r.err
}
func (r PostServiceStackPhpV81NginxV120Response) Headers() http.Header {
	return r.responseHeaders
}

func (r PostServiceStackPhpV81NginxV120Response) StatusCode() int {
	return r.responseStatusCode
}

func (h Handler) PostServiceStackPhpV81NginxV120(ctx context.Context, inputDtoBody body.PostUserNginxServiceStack) (postServiceStackPhpV81NginxV120Response PostServiceStackPhpV81NginxV120Response, err error) {
	u := "/api/rest/public/service-stack/php_v8_1_nginx_v1_20"

	var response PostServiceStackPhpV81NginxV120Response
	sdkResponse := sdkBase.Post(
		ctx,
		h.environment,
		u,
		inputDtoBody,
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
