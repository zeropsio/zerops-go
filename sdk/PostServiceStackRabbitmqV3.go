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

type PostServiceStackRabbitmqV3Response struct {
	success            output.ServiceStackProcess
	err                error
	responseHeaders    http.Header
	responseStatusCode int
}

func (r PostServiceStackRabbitmqV3Response) OutputInterface() (output interface{}, err error) {
	return r.success, r.err
}

func (r PostServiceStackRabbitmqV3Response) Output() (output output.ServiceStackProcess, err error) {
	return r.success, r.err
}

func (r PostServiceStackRabbitmqV3Response) Err() error {
	return r.err
}
func (r PostServiceStackRabbitmqV3Response) Headers() http.Header {
	return r.responseHeaders
}

func (r PostServiceStackRabbitmqV3Response) StatusCode() int {
	return r.responseStatusCode
}

func (h Handler) PostServiceStackRabbitmqV3(ctx context.Context, inputDtoBody body.PostStandardServiceStack) (postServiceStackRabbitmqV3Response PostServiceStackRabbitmqV3Response, err error) {
	u := "/api/rest/public/service-stack/rabbitmq_v3"

	var response PostServiceStackRabbitmqV3Response
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
