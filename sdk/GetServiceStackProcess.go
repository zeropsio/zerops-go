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

type GetServiceStackProcessResponse struct {
	success            output.ProcessList
	err                error
	responseHeaders    http.Header
	responseStatusCode int
}

func (r GetServiceStackProcessResponse) OutputInterface() (output interface{}, err error) {
	return r.success, r.err
}

func (r GetServiceStackProcessResponse) Output() (output output.ProcessList, err error) {
	return r.success, r.err
}

func (r GetServiceStackProcessResponse) Err() error {
	return r.err
}
func (r GetServiceStackProcessResponse) Headers() http.Header {
	return r.responseHeaders
}

func (r GetServiceStackProcessResponse) StatusCode() int {
	return r.responseStatusCode
}

func (h Handler) GetServiceStackProcess(ctx context.Context, inputDtoPath path.ServiceStackId, inputDtoQuery query.ListServiceStackProcesses) (getServiceStackProcessResponse GetServiceStackProcessResponse, err error) {
	u := "/api/rest/public/service-stack/" + inputDtoPath.Id.Native() + "/process"

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
	if param, ok := inputDtoQuery.ActionNameContains.Get(); ok {
		queryParams = append(queryParams, "actionNameContains="+url.QueryEscape(param.Native()))
	}

	if len(queryParams) > 0 {
		u += "?" + strings.Join(queryParams, "&")
	}

	var response GetServiceStackProcessResponse
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
