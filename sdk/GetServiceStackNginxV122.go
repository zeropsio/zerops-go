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

type GetServiceStackNginxV122Response struct {
	success            output.ServiceStackNginx
	err                error
	responseHeaders    http.Header
	responseStatusCode int
}

func (r GetServiceStackNginxV122Response) OutputInterface() (output interface{}, err error) {
	return r.success, r.err
}

func (r GetServiceStackNginxV122Response) Output() (output output.ServiceStackNginx, err error) {
	return r.success, r.err
}

func (r GetServiceStackNginxV122Response) Err() error {
	return r.err
}
func (r GetServiceStackNginxV122Response) Headers() http.Header {
	return r.responseHeaders
}

func (r GetServiceStackNginxV122Response) StatusCode() int {
	return r.responseStatusCode
}

func (h Handler) GetServiceStackNginxV122(ctx context.Context, inputDtoPath path.ServiceStackId) (getServiceStackNginxV122Response GetServiceStackNginxV122Response, err error) {
	u := "/api/rest/public/service-stack/nginx_v1_22/" + inputDtoPath.Id.Native() + ""

	var response GetServiceStackNginxV122Response
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
