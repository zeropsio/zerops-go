// Generated ZEROPS sdk

package sdk

import (
	"encoding/json"
	"errors"
	"net/http"

	"context"

	"github.com/zeropsio/zerops-go/apiError"
	"github.com/zeropsio/zerops-go/dto/input/body"
	"github.com/zeropsio/zerops-go/dto/input/path"
	"github.com/zeropsio/zerops-go/dto/output"
	"github.com/zeropsio/zerops-go/sdkBase"
)

type PutPrometheusConfigResponse struct {
	success            output.ProcessNil
	err                error
	responseHeaders    http.Header
	responseStatusCode int
}

func (r PutPrometheusConfigResponse) OutputInterface() (output interface{}, err error) {
	return r.success, r.err
}

func (r PutPrometheusConfigResponse) Output() (output output.ProcessNil, err error) {
	return r.success, r.err
}

func (r PutPrometheusConfigResponse) Err() error {
	return r.err
}
func (r PutPrometheusConfigResponse) Headers() http.Header {
	return r.responseHeaders
}

func (r PutPrometheusConfigResponse) StatusCode() int {
	return r.responseStatusCode
}

func (h Handler) PutPrometheusConfig(ctx context.Context, inputDtoPath path.ProjectId, inputDtoBody body.CoreStackPrometheusConfig) (putPrometheusConfigResponse PutPrometheusConfigResponse, err error) {
	u := "/api/rest/public/project/" + inputDtoPath.Id.Native() + "/core-stack/prometheus-config"

	var response PutPrometheusConfigResponse
	sdkResponse := sdkBase.Put(
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
