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

type GetProjectLogResponse struct {
	success            output.ProjectLog
	err                error
	responseHeaders    http.Header
	responseStatusCode int
}

func (r GetProjectLogResponse) OutputInterface() (output interface{}, err error) {
	return r.success, r.err
}

func (r GetProjectLogResponse) Output() (output output.ProjectLog, err error) {
	return r.success, r.err
}

func (r GetProjectLogResponse) Err() error {
	return r.err
}
func (r GetProjectLogResponse) Headers() http.Header {
	return r.responseHeaders
}

func (r GetProjectLogResponse) StatusCode() int {
	return r.responseStatusCode
}

func (h Handler) GetProjectLog(ctx context.Context, inputDtoPath path.ProjectId, inputDtoQuery query.GetProjectLog) (getProjectLogResponse GetProjectLogResponse, err error) {
	u := "/api/rest/public/project/" + inputDtoPath.Id.Native() + "/log"

	var queryParams []string
	if param, ok := inputDtoQuery.InstanceId.Get(); ok {
		queryParams = append(queryParams, "instanceId="+url.QueryEscape(param.Native()))
	}

	if len(queryParams) > 0 {
		u += "?" + strings.Join(queryParams, "&")
	}

	var response GetProjectLogResponse
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
