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

type GetProjectBackupResponse struct {
	success            output.ProjectBackupList
	err                error
	responseHeaders    http.Header
	responseStatusCode int
}

func (r GetProjectBackupResponse) OutputInterface() (output interface{}, err error) {
	return r.success, r.err
}

func (r GetProjectBackupResponse) Output() (output output.ProjectBackupList, err error) {
	return r.success, r.err
}

func (r GetProjectBackupResponse) Err() error {
	return r.err
}
func (r GetProjectBackupResponse) Headers() http.Header {
	return r.responseHeaders
}

func (r GetProjectBackupResponse) StatusCode() int {
	return r.responseStatusCode
}

func (h Handler) GetProjectBackup(ctx context.Context, inputDtoPath path.ProjectId, inputDtoQuery query.ListProjectBackups) (getProjectBackupResponse GetProjectBackupResponse, err error) {
	u := "/api/rest/public/project/" + inputDtoPath.Id.Native() + "/backup"

	var queryParams []string
	if param, ok := inputDtoQuery.Limit.Get(); ok {
		queryParams = append(queryParams, "limit="+url.QueryEscape(strconv.Itoa(param.Native())))
	}
	if param, ok := inputDtoQuery.Offset.Get(); ok {
		queryParams = append(queryParams, "offset="+url.QueryEscape(strconv.Itoa(param.Native())))
	}
	if param, ok := inputDtoQuery.ServiceStackId.Get(); ok {
		queryParams = append(queryParams, "serviceStackId="+url.QueryEscape(param.Native()))
	}
	if param, ok := inputDtoQuery.Sort.Get(); ok {
		queryParams = append(queryParams, "sort="+url.QueryEscape(strings.Join(param.Native(), ",")))
	}

	if len(queryParams) > 0 {
		u += "?" + strings.Join(queryParams, "&")
	}

	var response GetProjectBackupResponse
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
