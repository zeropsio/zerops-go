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

type GetProjectEnvFileDownloadResponse struct {
	success            output.FileDownload
	err                error
	responseHeaders    http.Header
	responseStatusCode int
}

func (r GetProjectEnvFileDownloadResponse) OutputInterface() (output interface{}, err error) {
	return r.success, r.err
}

func (r GetProjectEnvFileDownloadResponse) Output() (output output.FileDownload, err error) {
	return r.success, r.err
}

func (r GetProjectEnvFileDownloadResponse) Err() error {
	return r.err
}
func (r GetProjectEnvFileDownloadResponse) Headers() http.Header {
	return r.responseHeaders
}

func (r GetProjectEnvFileDownloadResponse) StatusCode() int {
	return r.responseStatusCode
}

func (h Handler) GetProjectEnvFileDownload(ctx context.Context, inputDtoPath path.ProjectId, inputDtoQuery query.GetProjectEnvFile) (getProjectEnvFileDownloadResponse GetProjectEnvFileDownloadResponse, err error) {
	u := "/api/rest/public/project/" + inputDtoPath.Id.Native() + "/env-file-download"

	var queryParams []string
	{
		param := inputDtoQuery.Name.Native()
		queryParams = append(queryParams, "name="+url.QueryEscape(param))
	}
	{
		param := inputDtoQuery.OverrideEnvIsolation.Native()
		queryParams = append(queryParams, "overrideEnvIsolation="+url.QueryEscape(param))
	}
	{
		param := inputDtoQuery.UserOnly.Native()
		queryParams = append(queryParams, "userOnly="+url.QueryEscape(strconv.FormatBool(param)))
	}

	if len(queryParams) > 0 {
		u += "?" + strings.Join(queryParams, "&")
	}

	var response GetProjectEnvFileDownloadResponse
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
