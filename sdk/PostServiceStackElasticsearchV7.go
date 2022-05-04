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

type PostServiceStackElasticsearchV7Response struct {
	success            output.ServiceStackProcess
	err                error
	responseHeaders    http.Header
	responseStatusCode int
}

func (r PostServiceStackElasticsearchV7Response) OutputInterface() (output interface{}, err error) {
	return r.success, r.err
}

func (r PostServiceStackElasticsearchV7Response) Output() (output output.ServiceStackProcess, err error) {
	return r.success, r.err
}

func (r PostServiceStackElasticsearchV7Response) Err() error {
	return r.err
}
func (r PostServiceStackElasticsearchV7Response) Headers() http.Header {
	return r.responseHeaders
}

func (r PostServiceStackElasticsearchV7Response) StatusCode() int {
	return r.responseStatusCode
}

func (h Handler) PostServiceStackElasticsearchV7(ctx context.Context, inputDtoBody body.PostStandardServiceStack) (postServiceStackElasticsearchV7Response PostServiceStackElasticsearchV7Response, err error) {
	u := "/api/rest/public/service-stack/elasticsearch_v7"

	var response PostServiceStackElasticsearchV7Response
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
